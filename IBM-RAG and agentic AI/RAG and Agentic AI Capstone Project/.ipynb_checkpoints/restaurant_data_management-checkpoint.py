from langchain_aws import ChatBedrock
from pydantic import BaseModel, Field, ValidationError
from typing import List, Optional
import json
import unittest
import os
import io
import shutil
from unittest.mock import patch

FILEPATH = 'structured_restaurant_data.json'
BACKUP_PATH = 'structured_restaurant_data.json.bak'

# ============================================================
# Pydantic Schema
# ============================================================
class Restaurant(BaseModel):
    name: str
    location: str
    type: str
    food_style: str
    rating: Optional[float] = None
    price_range: Optional[int] = None
    signatures: List[str] = Field(default_factory=list)
    vibe: Optional[str] = None
    environment: str
    shortcomings: List[str] = Field(default_factory=list)

# ============================================================
# Example for one-shot prompting
# ============================================================
EXAMPLE_RESTAURANT_PARAGRAPH = 'Down in **Santa Monica**, **Mar de Cortez** serves as a **sun-drenched**, **casual taqueria** specializing in **Baja-style seafood**. With a **4.2/5** rating, it captures the salt-air energy of the coast through its signature beer-battered snapper tacos and zesty octopus ceviche, making it a premier spot for open-air dining near the pier. Price range: $'

EXAMPLE_OUTPUT = """
{{
    "name": "Mar de Cortez",
    "location": "Santa Monica",
    "type": "casual taqueria",
    "food_style": "Baja-style seafood",
    "rating": 4.2,
    "price_range": 1,
    "signatures": [
        "beer-battered snapper tacos",
        "zesty octopus ceviche"
    ],
    "vibe": "salt-air energy",
    "environment": "a premier sun-drenched spot for open-air dining near the pier.",
    "shortcomings": []
}}
"""

# ============================================================
# Helper Functions
# ============================================================
def load_data(file_path):
    with open(file_path, 'r') as f:
        return json.load(f)

def save_data(file_path, data):
    # Create backup before saving
    if os.path.exists(file_path):
        shutil.copy(file_path, file_path + '.bak')
    with open(file_path, 'w') as f:
        json.dump(data, f, indent=4)

def show_restaurant_card(res, index):
    print(f"\n--- Record {index} ---")
    for key, value in res.items():
        print(f"{key}: {value}")

# ============================================================
# Prompt Generation
# ============================================================
def restaurant_data_structure_prompt_generation(restaurant_paragraph):
    base_system_msg = """
    You are a data extraction assistant that converts restaurant descriptions into structured JSON format.
    Follow these rules strictly:
    - Extract only information explicitly mentioned in the description.
    - For price_range, convert dollar signs to an integer (e.g., $ = 1, $$ = 2, $$$ = 3, $$$$ = 4).
    - For shortcomings, extract any negatives mentioned, otherwise return an empty list [].
    - Return ONLY valid JSON with no extra text, explanation, or markdown backticks.
    """

    base_user_prompt = f"""
    Task:
    Convert the restaurant description below into a structured JSON object with the following fields:
    name, location, type, food_style, rating, price_range, signatures, vibe, environment, shortcomings.

    Restaurant description:
    {restaurant_paragraph}

    Example:
    Input Restaurant Description: {EXAMPLE_RESTAURANT_PARAGRAPH}
    Output:
    {EXAMPLE_OUTPUT}

    Now extract the JSON for the restaurant description above. Return only the JSON object.
    """

    return base_system_msg, base_user_prompt

# ============================================================
# LLM Model
# ============================================================
def llm_model(system_msg, prompt_txt, params=None):
    if params is None:
        params = {"max_tokens": 1000}

    llm = ChatBedrock(
        model_id="amazon.nova-micro-v1:0",
        region_name="us-east-1",
        model_kwargs=params
    )

    messages = [
        {"role": "system", "content": system_msg},
        {"role": "user",   "content": prompt_txt}
    ]

    response = llm.invoke(messages)
    return response.content

# ============================================================
# JSON Auto Repair Prompts
# ============================================================
def JSON_auto_repair_prompts(candidate_json_output, error_message):
    auto_repair_system_msg = """
    You are a JSON repair expert. Your sole responsibility is to fix invalid or malformed JSON outputs.
    Follow these rules strictly:
    - Return ONLY the corrected valid JSON object, no explanations or markdown backticks.
    - Do not change or remove any correct data, only fix what is broken.
    - Ensure the output is parseable by Python's json.loads() function.
    - If a field is missing, use a sensible default (e.g., [] for lists, null for unknown values).
    - For price_range, ensure it is an integer (e.g., $ = 1, $$ = 2, $$$ = 3, $$$$ = 4).
    """

    auto_repair_prompt = f"""
    The following JSON output is invalid or does not conform to the required schema.

    Incorrect JSON output:
    {candidate_json_output}

    Error message from schema validation:
    {error_message}

    Please fix the JSON output based on the error message above.
    Return ONLY the corrected valid JSON object.
    """

    return auto_repair_system_msg, auto_repair_prompt

# ============================================================
# New Data Entry Process
# ============================================================
def new_data_entry_process(paragraph, itemId):

    ### Step 1: Generate initial output
    system_msg, user_prompt = restaurant_data_structure_prompt_generation(paragraph)
    response = llm_model(system_msg, user_prompt)

    ### Step 2: Validation and auto correction loop
    for _ in range(3):
        try:
            parsed = json.loads(response)
            Restaurant.model_validate(parsed)
            break
        except (json.JSONDecodeError, ValidationError) as e:
            error_message = str(e)
            repair_system_msg, repair_prompt = JSON_auto_repair_prompts(response, error_message)
            response = llm_model(repair_system_msg, repair_prompt)

    ### Step 3: Add itemId
    parsed['itemId'] = itemId

    return parsed

# ============================================================
# Restaurant Management UI
# ============================================================
def manage_restaurants(file_path, backup_path):
    while True:
        data = load_data(file_path)
        print(f"\n🏨 RESTAURANT DATABASE | Records: {len(data)}")
        print("1. Browse All (Names)")
        print("2. View Detailed Record")
        print("3. Add New Restaurant")
        print("4. Edit Restaurant Info")
        print("5. Delete Restaurant")
        print("6. Exit")

        choice = input("\nAction: ")

        if choice == '1':
            print("\n--- Current Listings ---")
            for i, record in enumerate(data):
                name = record.get('name', 'N/A')
                print(f"{i}. {name}")

        elif choice == '2':
            try:
                index = int(input("Enter record index: "))
                if 0 <= index < len(data):
                    show_restaurant_card(data[index], index)
                else:
                    print("Invalid index.")
            except ValueError:
                print("Invalid index.")

        elif choice in ['3', '4', '5']:
            print("\n❗ SECURITY WARNING: You are entering write-mode.")
            print("Changes will be saved to the database immediately.")
            confirm = input("Are you sure? (type 'yes' to proceed): ").lower()
            if confirm != 'yes':
                print("Operation cancelled.")
                continue

            if choice == '3':  # ADD NEW DATA
                itemId = 1000000 + len(data) + 1
                paragraph = input("Enter new restaurant description: ")
                new_restaurant = new_data_entry_process(paragraph, itemId)
                data.append(new_restaurant)
                save_data(file_path, data)
                print("✅ Restaurant added.")

            elif choice == '4':  # EDIT DATA
                try:
                    index = int(input("Enter record index to edit: "))
                    if 0 <= index < len(data):
                        record = data[index]
                        print(f"Editing: {record.get('name', 'N/A')} (Press Enter to skip a field)")
                        for key in record.keys():
                            current_value = record[key]
                            new_value = input(f"{key} [{current_value}]: ")
                            if new_value.strip() != '':
                                try:
                                    if isinstance(current_value, int):
                                        record[key] = int(new_value)
                                    elif isinstance(current_value, float):
                                        record[key] = float(new_value)
                                    elif isinstance(current_value, list):
                                        record[key] = json.loads(new_value)
                                    else:
                                        record[key] = new_value
                                except (ValueError, json.JSONDecodeError):
                                    record[key] = new_value
                        data[index] = record
                        save_data(file_path, data)
                        print("✅ Record updated.")
                    else:
                        print("Invalid index.")
                except ValueError:
                    print("Invalid index.")

            elif choice == '5':  # DELETE DATA
                try:
                    index = int(input("Enter record index to delete: "))
                    if 0 <= index < len(data):
                        deleted = data.pop(index)
                        save_data(file_path, data)
                        print(f"✅ '{deleted.get('name', 'N/A')}' deleted.")
                    else:
                        print("Invalid index.")
                except ValueError:
                    print("Invalid index.")

        elif choice == '6':
            break

        else:
            print("Invalid input.")

# ============================================================
# Unit Tests
# ============================================================
class TestRestaurantDatabase(unittest.TestCase):

    def setUp(self):
        """Create a temporary clean database for testing."""
        self.test_file = 'structured_restaurant_data_unit_test.json'
        self.test_file_backup = 'structured_restaurant_data_unit_test.json.bak'
        self.initial_data = [{"name": "Test Cafe", "location": "Test City"}]
        with open(self.test_file, 'w') as f:
            json.dump(self.initial_data, f)

    def tearDown(self):
        """Clean up the test file after tests."""
        if os.path.exists(self.test_file):
            os.remove(self.test_file)
        if os.path.exists(self.test_file_backup):
            os.remove(self.test_file_backup)

    @patch('builtins.input')
    @patch('sys.stdout', new_callable=io.StringIO)
    def test_add_and_delete_restaurant_success(self, mock_stdout, mock_input):
        """
        Test Scenario: Add a new restaurant.
        Inputs: '3' (Add), 'yes' (Confirm), 'New Burger Joint', '6' (Exit)
        """
        mock_restaurant = 'The Copper Sprout is a high-concept, Modern Appalachian farm-to-table destination that blends an industrial-chic aesthetic with rustic forest charm, featuring reclaimed wood and amber lighting to create a sophisticated yet cozy vibe. Priced in the $$ category, the menu celebrates seasonal foraging and local heritage, headlined by signature dishes like Cast-Iron Smoked Trout with pickled fiddlehead ferns and hand-foraged Wild Mushroom Risotto with aged goat cheese. The experience is designed to be intimate and earthy, making it a premier spot for those seeking high-quality, smokehouse-influenced cuisine in a refined, atmospheric setting.'
        mock_input.side_effect = ['3', 'yes', mock_restaurant, '6']

        try:
            manage_restaurants(self.test_file, self.test_file_backup)
        except SystemExit:
            pass

        with open(self.test_file, 'r') as f:
            data = json.load(f)

        print(data)
        self.assertEqual(len(data), 2)
        self.assertIn("✅ Restaurant added.", mock_stdout.getvalue())

        mock_input.side_effect = ['5', 'yes', '1', '6']

        try:
            manage_restaurants(self.test_file, self.test_file_backup)
        except SystemExit:
            pass

        with open(self.test_file, 'r') as f:
            data = json.load(f)

        print(data)
        self.assertEqual(len(data), 1)

    @patch('builtins.input')
    @patch('sys.stdout', new_callable=io.StringIO)
    def test_delete_security_cancel(self, mock_stdout, mock_input):
        """
        Test Scenario: Try to delete but say 'no' to security warning.
        Inputs: '5' (Delete), 'no' (Cancel), '6' (Exit)
        """
        mock_input.side_effect = ['5', 'no', '6']

        manage_restaurants(self.test_file, self.test_file_backup)

        with open(self.test_file, 'r') as f:
            data = json.load(f)

        self.assertEqual(len(data), 1)
        self.assertIn("Operation cancelled.", mock_stdout.getvalue())

# ============================================================
# Entry Point
# ============================================================
if __name__ == "__main__":
    unittest.main()
    #manage_restaurants(FILEPATH, BACKUP_PATH)  # Uncomment to run UI