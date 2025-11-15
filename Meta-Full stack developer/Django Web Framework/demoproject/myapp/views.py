from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
def home(request):
    return HttpResponse("hello world!!!")
def hello(request,name):
    content=f"<html> <body><h1> Welcome to my app {name}</h1></body></html>"
    return HttpResponse(content)