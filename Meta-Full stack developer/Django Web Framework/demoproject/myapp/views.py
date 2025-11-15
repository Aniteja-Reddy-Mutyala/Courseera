from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
def home(request):
    return HttpResponse("hello world!!!")
def hello(request):
    content="<html> <body><h1> Welcome to my app</h1></body></html>"
    return HttpResponse(content)