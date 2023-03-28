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
    url = URL+'getCount'
    response = requests.get(url).json()
    context = {'ob':'admin/dashboard.html','count':response}
    return render(request,'admin/adminHome.html',context)



def adminAddUser(request):
       if request.method == 'POST':       
         name = request.POST.get('name')
         email = request.POST.get('email') 
         password = request.POST.get('password')
         address = request.POST.get('address')
         type = request.POST.get('type') 
         contact = request.POST.get('contact')
         request.FILES['image'].name = email
         files =  request.FILES
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
         url = URL+'deleteUser/'+str(id)     
         requests.delete(url)
         return redirect('/adminHome')


def addCourse(request):
       if request.method == 'POST':
           data = {
               'name':request.POST.get('name'),
               'discipline':request.POST.get('discipline'),
               'qualityPoints':float(request.POST.get('qualityPoints')),
               'code':request.POST.get('code'),
           }
           url = URL+'addCourse/'
           print('DATA',data)
           requests.post(url,json=data)
           return redirect('/adminHome')  
       else:    
        return render(request,'admin/addCourse.html')
       


def courses(request):
    url = URL +'getCourses'
    response = requests.get(url).json()
    context = {'ob':response}
    print('context',context)
    return render(request,'admin/courses.html',context)




def deleteCourse(request,id):     
         url = URL+'deleteCourse/'+str(id)     
         requests.delete(url)
         return redirect('/adminHome')




def enrollment(request):
     url1 = URL +'getCourses'
     response1 = requests.get(url1).json()
     url2 = URL +'adminUsers/student'
     response2 = requests.get(url2).json()
     context = {'courses':response1,'students':response2}
     return render(request,'admin/enrollment.html',context)


      
        

