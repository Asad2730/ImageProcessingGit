package models

import ("gorm.io/gorm")

type User struct {	
	gorm.Model
   Id int `json:"id" gorm:"primaryKey;autoIncrement:true"`
   Name string `json:"name"`
   Email  string `json:"email"`
   Password string `json:"password"`
   Address string `json:"address"`
   Contact string `json:"contact"`
   Type string `json:"type"`
   Image string `json:"image"`
}


type Course struct {	
	gorm.Model
	Id int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
	Discipline string `json:"discipline"`
	QualityPoints  float32 `json:"qualityPoints"`
	Code string `json:"code"`
 }


 type Enrollment struct {	
	gorm.Model
	Id int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Uid int `json:"uid"`
	Cid int `json:"cid"`
	Semester string `json:"semester"`
	Grade string `json:"grade"`
	Year  int `json:"year"`
	Section string `json:"section"`
 }


 type Allocate struct {	
	gorm.Model
	Id int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Uid int `json:"uid"`
	Cid int `json:"cid"`
	Semester string `json:"semester"`
	Year  int `json:"year"`
	Section string `json:"section"`
 }


 type Attendance struct {
	gorm.Model
	Id int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Eid int `json:"eid"`
	Aid int `json:"aid"`
	Status string `json:"status"`
	Date  int `json:"date"`
 }