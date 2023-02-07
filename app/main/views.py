from django.shortcuts import render,redirect,get_object_or_404,get_list_or_404
import requests

# Create your views here.

URL = 'http://localhost:3000/'

def login(request):
    if request.method == 'POST':
        email = request.POST.get('email') 
        password = request.POST.get('password') 
        url = URL+'login/'+email+'/'+password
        res = requests.get(url).json()
        print('RESPONSE',res)
        return render(request, 'login.html')
    else: 
        return render(request, 'login.html')
        
            
   


