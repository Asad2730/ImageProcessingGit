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
         path = "images/"+request.FILES['image'].name
         jsn = {
            "name":name,
            "email":email,
            "password":password,
            "address":address,
            "type":type,
            "contact":contact,
             "image":path,
            }
         data = {'file': files, 'json_data': jsn}   
        
         url = URL+'uploadImage'     
         response = requests.post(url,files=data['file'])
         if response.ok:
           url = URL+'addUser'     
           requests.post(url,json=data['json_data'])
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
            
   
def updateUser(request,id):
       if request.method == 'POST':     
         name = request.POST.get('name')
         email = request.POST.get('email') 
         password = request.POST.get('password')
         address = request.POST.get('address')
         contact = request.POST.get('contact')
         files = request.FILES
         data = {
            "name":name,
            "email":email,
            "password":password,
            "address":address,
            "contact":contact,
          
            }
         url = URL+'updateUser/'+id     
         requests.post(url,json=data,files=files)
         return redirect('/adminHome')        
       else:
        url = URL+'singleUser/'+id
        data = requests.get(url).json() 
        context = {'ob':data}
        return render(request,'admin/update.html',context)


def deleteUser(request,id):
       if request.method == 'POST':     
         url = URL+'deleteUser/'+id     
         requests.delete(url)
         return redirect('/adminHome')        
        

