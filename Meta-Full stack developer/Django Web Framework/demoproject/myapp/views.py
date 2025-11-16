from django.shortcuts import render
from .forms import DemoForm
from .models import Menu
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
def menu_item(request):
    menuItem={"mains":[{"name":"Greek salad","price":"15"},{"name":"Falafel","price":"20"},{"name":"gyro","price":"25"}]}
    return render(request,"menu.html",menuItem)

def menu_id(request):
    newMenu=Menu.objects.all()
    newMenu_dict={'menu':newMenu}
    return render(request,'menu_cards.html',newMenu_dict)
