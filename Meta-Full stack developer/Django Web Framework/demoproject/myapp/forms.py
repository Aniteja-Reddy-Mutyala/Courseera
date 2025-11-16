from django import forms
from .models import Menu
class DemoForm(forms.ModelForm):
    class Meta:
        model=Menu
        fields='__all__'
    