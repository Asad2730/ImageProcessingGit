from django.shortcuts import render,redirect
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


def signUp(request):
    if request.method == 'POST': 
         name = request.POST.get('name')
         email = request.POST.get('email') 
         password = request.POST.get('password')
         address = request.POST.get('address')
         type = request.POST.get('type') 
         contact = request.POST.get('contact')
         data = {
            "name":name,
            "email":email,
            "password":password,
            "address":address,
            "type":type,
            "contact":contact,
            "image":"./img.jpg"
            }
         url = URL+'signup'
         requests.post(url,json=data)
         return redirect('login')        
    else: 
        return render(request,'signup.html')        
            
   


