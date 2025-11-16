from django.urls import path
from . import views

urlpatterns=[
    path('',views.home,name="home"),
    path('hello/<str:name>',views.hello,name="hello"),
    path("form/",views.form_view,name="forms"),
    path("menu/",views.menu_item,name="menu_item"),
    path("menu/cards/",views.menu_id,name="menu_id")


]