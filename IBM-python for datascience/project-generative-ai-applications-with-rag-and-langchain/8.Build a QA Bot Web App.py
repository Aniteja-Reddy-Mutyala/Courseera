from langchain_community.document_loaders import PyPDFLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_aws import BedrockEmbeddings
from langchain_community.vectorstores import Chroma
from langchain_classic.chains import RetrievalQA
from langchain_groq import ChatGroq
import gradio as gr
import os 
import getpass

def set_if_undefined(var:str): 
    if os.environ.get(var): 
        return
    os.environ[var] = getpass.getpass(var)
set_if_undefined("GROQ_API_KEY")    

# ============================================================
# 1. Document Loader
# ============================================================
def document_loader(file):
    loader = PyPDFLoader(file)
    docs = loader.load()
    return docs

# ============================================================
# 2. Text Splitter
# ============================================================
def text_splitter(docs):
    splitter = RecursiveCharacterTextSplitter(
        chunk_size=1000,
        chunk_overlap=100,
        length_function=len
    )
    chunks = splitter.split_documents(docs)
    return chunks

# ============================================================
# 3. Embedding Model (AWS Bedrock)
# ============================================================
def watsonx_embedding():
    embedding_model = BedrockEmbeddings(
        model_id="amazon.titan-embed-text-v2:0",
        region_name="us-east-1"
    )
    return embedding_model

# ============================================================
# 4. Vector Database
# ============================================================
def vector_database(chunks):
    embedding_model = watsonx_embedding()
    vectordb = Chroma.from_documents(
        documents=chunks,
        embedding=embedding_model
    )
    return vectordb

# ============================================================
# 5. Retriever
# ============================================================
def retriever(file):
    docs = document_loader(file)
    chunks = text_splitter(docs)
    vectordb = vector_database(chunks)
    retriever = vectordb.as_retriever(
        search_type="similarity",
        search_kwargs={"k": 9}
    )
    return retriever

# ============================================================
# 6. LLM — Groq
# ============================================================
def get_llm():
    llm = ChatGroq(
        model="llama-3.3-70b-versatile",  # or "mixtral-8x7b-32768", "gemma2-9b-it"
        temperature=0.5,
        
        
    )
    return llm

# ============================================================
# 7. QA Chain
# ============================================================
def retriever_qa(file, query):
    llm = get_llm()
    retriever_obj = retriever(file)
    qa_chain = RetrievalQA.from_chain_type(
        llm=llm,
        chain_type="stuff",
        retriever=retriever_obj,
        return_source_documents=True
    )
    result = qa_chain.invoke({"query": query})
    return result["result"]

# ============================================================
# 8. Gradio Interface
# ============================================================
app = gr.Interface(
    fn=retriever_qa,
    inputs=[
        gr.File(label="Upload PDF", file_types=[".pdf"]),
        gr.Textbox(label="Ask a Question", placeholder="What this paper is talking about?")
    ],
    outputs=gr.Textbox(label="Answer"),
    title="QA Bot ",
    description="Upload a PDF and ask questions about its content"
)

if __name__ == "__main__":
    app.launch()