from django.shortcuts import render
from .forms import DemoForm
# Create your views here.
from django.http import HttpResponse
def home(request):
    return HttpResponse("hello world!!!")
def hello(request,name):
    content=f"<html> <body><h1> Welcome to my app {name}</h1></body></html>"
    return HttpResponse(content)

def form_view(request):
    form=DemoForm()
    if request.method=='POST':
        form=DemoForm(request.POST)
    if form.is_valid():
        form.save()    
    context={"form":form}
    return render(request,"home.html",context)
