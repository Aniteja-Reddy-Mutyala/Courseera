from django.urls import path
from . import views

urlpatterns=[
    path('',views.home,name="home"),
    path('hello/<str:name>',views.hello,name="hello"),
    path("form/",views.form_view,name="forms")

]