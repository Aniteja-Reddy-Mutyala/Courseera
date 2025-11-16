from django import forms

class DemoForm(forms.Form):
    name=forms.CharField(required=False)
    