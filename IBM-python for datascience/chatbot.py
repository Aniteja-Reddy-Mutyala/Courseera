from transformers import AutoTokenizer, AutoModelForSeq2SeqLM

model_name = "facebook/blenderbot-400M-distill"

model = AutoModelForSeq2SeqLM.from_pretrained(model_name)
tokenizer = AutoTokenizer.from_pretrained(model_name)

conversation_history = []

print("Chatbot: Hello! I'm BlenderBot. Type 'quit' to exit.")

while True:
    # Get user input
    input_text = input("You: ")
    
    # Exit condition
    if input_text.lower() in ['quit', 'exit', 'bye']:
        print("Chatbot: Goodbye!")
        break
    
    # Build conversation context
    history_string = "\n".join(conversation_history)
    
    # Encode input with history
    inputs = tokenizer.encode_plus(
        history_string, 
        input_text, 
        return_tensors="pt"  
    
    # Generate response
    outputs = model.generate(**inputs, max_length=100)
    
    # Decode response
    response = tokenizer.decode(outputs[0], skip_special_tokens=True)
    print(f"Chatbot: {response}")
    
    # Update conversation history
    conversation_history.append(input_text)
    conversation_history.append(response)