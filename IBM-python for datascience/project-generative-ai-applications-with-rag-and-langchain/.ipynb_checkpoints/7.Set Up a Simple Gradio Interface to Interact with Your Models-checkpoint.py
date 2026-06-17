'''
import gradio as gr

def sentence_builder(quantity, tech_worker_type, countries, place, activity_list, morning):
    return f"""The {quantity} {tech_worker_type}s from {" and ".join(countries)} went to the {place} where they {" and ".join(activity_list)} until the {"morning" if morning else "night"}"""

demo = gr.Interface(
    fn=sentence_builder,
    inputs=[
        gr.Slider(3, 20, value=4, step=1, label="Count", info="Choose between 3 and 20"),
        gr.Dropdown(
            ["Data Scientist", "Software Developer", "Software Engineer"], 
			label="tech_worker_type", 
			info="Will add more tech worker types later!"
        ),
        gr.CheckboxGroup(["Canada", "Japan", "France"], label="Countries", info="Where are they from?"),
        gr.Radio(["office", "restaurant", "meeting room"], label="Location", info="Where did they go?"),
        gr.Dropdown(
            ["partied", "brainstormed", "coded", "fixed bugs"], 
			value=["brainstormed", "fixed bugs"], 
			multiselect=True, 
			label="Activities", 
			info="Which activities did they perform?"
        ),
        gr.Checkbox(label="Morning", info="Did they do it in the morning?"),
    ],
    outputs="text",
    examples=[
        [3, "Software Developer", ["Canada", "Japan"], "restaurant", ["coded", "fixed bugs"], True],
        [4, "Data Scientist", ["Japan"], "office", ["brainstormed", "partied"], False],
        [10, "Software Engineer", ["Canada", "France"], "meeting room", ["brainstormed"], False],
        [8, "Data Scientist", ["France"], "restaurant", ["coded"], True],
    ]
)

demo.launch(server_name="127.0.0.1", server_port= 7860)
'''
import gradio as gr 
from langchain_groq import ChatGroq
import os 
import getpass

def set_if_undefined(var:str): 
    if os.environ.get(var): 
        return
    os.environ[var] = getpass.getpass(var) 
set_if_undefined("GROQ_API_KEY")

groq_llm = ChatGroq(
    model = "llama-3.3-70b-versatile", 
    temperature = 0.5, 
    max_tokens = 512, 
    model_kwargs = {
        "top_p": 0.6
    }
)

#Function to generate a response 
def generate_response(prompt_txt): 
    response = groq_llm.invoke(prompt_txt)
    return response.content 

#Creating Gradio interface 
chat_app = gr.Interface(
    fn = generate_response,
    inputs = gr.Textbox(label = "Input", lines = 2, placeholder = "Enter your query"), 
    outputs = gr.Textbox(label = "Outputs"), 
    title = "Groq ai chatbot",  
    description = "This chatbot will try to answer your query"
)
chat_app.launch(server_name= "127.0.0.1", server_port = 7860)