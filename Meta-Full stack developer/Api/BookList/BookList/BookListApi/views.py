from django.shortcuts import render

# Create your views here.

from .models import Book
from .serializers import BookSerializer
from rest_framework import generics
from rest_framework.permissions import IsAuthenticated
from rest_framework.decorators import permission_classes,api_view
from rest_framework.response import Response


class BookView(generics.ListCreateAPIView):
    queryset = Book.objects.all()
    serializer_class = BookSerializer

class SingleBookView(generics.RetrieveUpdateAPIView):
    queryset = Book.objects.all()
    serializer_class = BookSerializer
@api_view()
@permission_classes([IsAuthenticated])
def secret(request):
    return   Response({"message":"Secret"})  