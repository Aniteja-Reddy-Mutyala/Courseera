import requests
from PIL import Image
from transformers import AutoProcessor,BlipForConditionalGeneration
# Loading the pretrained processor and model
processor =AutoProcessor.from_pretrained("Salesforce/blip-image-captioning-base")
model = BlipForConditionalGeneration.from_pretrained("Salesforce/blip-image-captioning-base")

#Loading an image
IMG_PATH = "25.jpg"

#Converting it into RGB format
image = Image.open(IMG_PATH).convert("RGB")

text = "The image of"
inputs = processor(images = image, text =text ,return_tensors = "pt")

#Generating a caption for an image
output = model.generate( **inputs,max_length = 50)

#Decoding the generated tokens to text
caption = processor.decode(output[0],skip_special_tokens = True)
print(caption)