# ================================================================================
# GROQ CHATBOT WITH GRADIO
# ================================================================================

# ================================
# 1. IMPORT NECESSARY PACKAGES
# ================================

from langchain_groq import ChatGroq
import gradio as gr

# ================================
# 2. MODEL AND PARAMETERS SETUP
# ================================

# Model selection
model_id = 'llama-3.1-70b-versatile'  # Groq's Llama model
# Alternative models:
# model_id = 'llama-3.1-8b-instant'     # Faster, smaller
# model_id = 'llama-3.3-70b-versatile'  # Newer Llama 3.3
# model_id = 'gemma2-9b-it'             # Google's Gemma

# Set parameters
groq_api_key = "your-groq-api-key-here"  # Replace with your actual API key

# Create Groq LLM
groq_llm = ChatGroq(
    model=model_id,
    groq_api_key=groq_api_key,
    temperature=0.5,      # Creativity/randomness (0.0-1.0)
    max_tokens=256,       # Max tokens to generate
)

print("✅ Groq LLM initialized successfully!")

# ================================
# 3. FUNCTION TO GENERATE RESPONSE
# ================================

def generate_response(prompt_txt):
    """Generate a response from the Groq model"""
    try:
        generated_response = groq_llm.invoke(prompt_txt)
        return generated_response.content  # Extract text from response
    except Exception as e:
        return f"Error: {str(e)}"

# ================================
# 4. CREATE GRADIO INTERFACE
# ================================

chat_application = gr.Interface(
    fn=generate_response,
    allow_flagging="never",
    inputs=gr.Textbox(
        label="Input",
        lines=2,
        placeholder="Type your question here..."
    ),
    outputs=gr.Textbox(label="Output"),
    title="Groq Chatbot with Llama",
    description="Ask any question and the chatbot will try to answer using Groq's Llama model."
)

# ================================
# 5. LAUNCH THE APP
# ================================

print("🚀 Launching Gradio app...")
print("📍 Access at: http://127.0.0.1:7860")

chat_application.launch(
    server_name="127.0.0.1",
    server_port=7860,
    share=False  # Set to True to create a public link
)