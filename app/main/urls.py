from . import views
from django.urls import path

urlpatterns = [
    path('login', views.login,name='login'),
    path('adminHome/',views.adminHome),
    path('adminAddUser/',views.adminAddUser),
    path('adminStudent/',views.studentAdmin),
    path('adminTeacher/',views.teacherAdmin),
    path('singleUser/<int:id>/',views.updateUser),
    path('deleteUser/<int:id>/',views.deleteUser),
    path('addCourse/',views.addCourse ),
    path('adminCourses/',views.courses),
    path('deleteCourse/<int:id>/',views.deleteCourse),
    path('enrollment/',views.enrollment),
]