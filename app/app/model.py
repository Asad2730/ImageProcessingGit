from django.db import models


class User(models.Model):
    id = models.AutoField(primary_key=True)
    name = models.CharField(max_length=24)
    email = models.EmailField(max_length=24)
    password = models.CharField(max_length=100)
    address = models.CharField(max_length=100)
    contact = models.CharField(max_length=20)
    type = models.CharField(max_length=20)
    image  = models.ImageField(upload_to='images')


class Course(models.Model):
     id = models.AutoField(primary_key=True)
     name = models.CharField(max_length=24)
     discipline = models.CharField(max_length=24)
     credit_hours = models.IntegerField(blank=True, null=True)
     qty_Points = models.FloatField(blank=True, null=True)
     course_code = models.CharField(max_length = 24)


class Enrollment(models.Model):
     id = models.AutoField(primary_key=True)
     uid = models.IntegerField(blank=True, null=True)
     cid = models.IntegerField(blank=True, null=True)
     semester = models.IntegerField(blank=True, null=True)
     grade = models.CharField(max_length=20)
     year = models.IntegerField(blank=True, null=True)
     section = models.CharField(max_length=20)

class Allocate(models.Model):    
     id = models.AutoField(primary_key=True)
     uid = models.IntegerField(blank=True, null=True)
     cid = models.IntegerField(blank=True, null=True)
     semester = models.IntegerField(blank=True, null=True)
     year = models.IntegerField(blank=True, null=True)
     section = models.CharField(max_length=20)


class Attendance(models.Model):
     id = models.AutoField(primary_key=True)
     uid = models.IntegerField(blank=True, null=True)
     cid = models.IntegerField(blank=True, null=True)
     aid = models.IntegerField(blank=True, null=True)
     status = models.CharField(max_length=24)
     date = models.CharField(max_length=24)




     

