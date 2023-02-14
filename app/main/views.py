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




def adminHome(request):
    conetext = {'ob':'admin/dashboard.html'}
    return render(request,'admin/adminHome.html',conetext)


def adminAddUser(request):
       if request.method == 'POST': 
        
         name = request.POST.get('name')
         email = request.POST.get('email') 
         password = request.POST.get('password')
         address = request.POST.get('address')
         type = request.POST.get('type') 
         contact = request.POST.get('contact')
         files = request.FILES
         data = {
            "name":name,
            "email":email,
            "password":password,
            "address":address,
            "type":type,
            "contact":contact,
          
            }
         url = URL+'addUser'     
         requests.post(url,json=data,files=files)
         return redirect('/adminHome')        
       else: 
        return render(request,'admin/add.html')
     


def studentAdmin(request):
    url = URL +'adminUsers/student'
    res = requests.get(url).json()
    context = {'ob':res}
    print('context',context)
    return render(request,'admin/users.html',context)

def teacherAdmin(request):
    url = URL +'adminUsers/teacher'
    res = requests.get(url).json()
    context = {'ob':res}
    return render(request,'admin/users.html',context)  
            
   


