{
 "cells": [
  {
   "cell_type": "code",
   "execution_count": 2,
   "id": "ae60a6b5-951b-4b16-8aa4-8e41a68263e4",
   "metadata": {},
   "outputs": [
    {
     "name": "stdout",
     "output_type": "stream",
     "text": [
      "\u001b[33mWARNING: Ignoring invalid distribution ~orch (/Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages)\u001b[0m\u001b[33m\n",
      "\u001b[0mRequirement already satisfied: jinja2 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (3.1.6)\n",
      "Requirement already satisfied: starlette in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (0.50.0)\n",
      "Requirement already satisfied: uvicorn in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (0.40.0)\n",
      "Requirement already satisfied: MarkupSafe>=2.0 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (from jinja2) (3.0.3)\n",
      "Requirement already satisfied: anyio<5,>=3.6.2 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (from starlette) (4.12.1)\n",
      "Requirement already satisfied: typing-extensions>=4.10.0 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (from starlette) (4.15.0)\n",
      "Requirement already satisfied: idna>=2.8 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (from anyio<5,>=3.6.2->starlette) (3.11)\n",
      "Requirement already satisfied: click>=7.0 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (from uvicorn) (8.4.1)\n",
      "Requirement already satisfied: h11>=0.8 in /Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages (from uvicorn) (0.16.0)\n",
      "\u001b[33mWARNING: Ignoring invalid distribution ~orch (/Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages)\u001b[0m\u001b[33m\n",
      "\u001b[0m\u001b[33mWARNING: Ignoring invalid distribution ~orch (/Users/anirudh/.local/share/virtualenvs/Desktop-sFnGVMJ4/lib/python3.11/site-packages)\u001b[0m\u001b[33m\n",
      "\u001b[0m"
     ]
    }
   ],
   "source": [
    "!pip install jinja2 starlette uvicorn"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 1,
   "id": "56867eac-73ed-4ce8-8fbc-3b5d4d388272",
   "metadata": {},
   "outputs": [],
   "source": [
    "import gradio as gr \n",
    "from huggingface_hub import HfFolder\n"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": null,
   "id": "d6aa1a81-55c1-4cae-b487-a0de36617bf4",
   "metadata": {},
   "outputs": [],
   "source": []
  },
  {
   "cell_type": "code",
   "execution_count": null,
   "id": "226fefc7-9eae-44cc-a4ee-ce8e1ca4df7b",
   "metadata": {},
   "outputs": [],
   "source": [
    "def add_two_numbers(a,b): \n",
    "    return a + \" \"+ b\n",
    "\n",
    "demo = gr.interface(\n",
    "    fn = add_two_numbers, \n",
    "    inputs = [gr., gr.Number()], \n",
    "    outputs = gr.Number() \n",
    "    \n",
    ")\n",
    "# Launch the interface\n",
    "demo.launch(server_name=\"127.0.0.1\", server_port= 7860)"
   ]
  },
  {
   "cell_type": "code",
   "execution_count": 1,
   "id": "edefcff9-8aa5-4da8-ac70-45daade45bf9",
   "metadata": {},
   "outputs": [
    {
     "name": "stdin",
     "output_type": "stream",
     "text": [
      "GROQ_API_KEY ········\n"
     ]
    },
    {
     "name": "stdout",
     "output_type": "stream",
     "text": [
      "* Running on local URL:  http://127.0.0.1:7860\n",
      "* To create a public link, set `share=True` in `launch()`.\n"
     ]
    },
    {
     "data": {
      "text/html": [
       "<div><iframe src=\"http://127.0.0.1:7860/\" width=\"100%\" height=\"500\" allow=\"autoplay; camera; microphone; clipboard-read; clipboard-write;\" frameborder=\"0\" allowfullscreen></iframe></div>"
      ],
      "text/plain": [
       "<IPython.core.display.HTML object>"
      ]
     },
     "metadata": {},
     "output_type": "display_data"
    }
   ],
   "source": [
    "%run \"8.Build a QA Bot Web App.py\""
   ]
  },
  {
   "cell_type": "code",
   "execution_count": null,
   "id": "b5b697ab-3e6e-4a5d-9abd-18d970beda29",
   "metadata": {},
   "outputs": [],
   "source": []
  }
 ],
 "metadata": {
  "kernelspec": {
   "display_name": "Python 3 (ipykernel)",
   "language": "python",
   "name": "python3"
  },
  "language_info": {
   "codemirror_mode": {
    "name": "ipython",
    "version": 3
   },
   "file_extension": ".py",
   "mimetype": "text/x-python",
   "name": "python",
   "nbconvert_exporter": "python",
   "pygments_lexer": "ipython3",
   "version": "3.11.4"
  }
 },
 "nbformat": 4,
 "nbformat_minor": 5
}
