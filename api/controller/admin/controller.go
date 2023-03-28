package controller

import (
	"api/initializers"
	"api/models"
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
)

func GetCount(c *gin.Context) {
	var students int64
	var teachers int64

	initializers.DB.Model(models.User{}).Where("type = ?", "student").Count(&students)
	initializers.DB.Model(models.User{}).Where("type = ?", "teacher").Count(&teachers)

	c.JSON(200, gin.H{
		"students": students,
		"teachers": teachers,
	})
}

func AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	initializers.DB.Create(&user)
	c.JSON(200, &user)
}

func GetImage(c *gin.Context) {
	filename := c.Param("filename")
	c.File("images/" + filename)
}

func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")

	if err != nil {
		fmt.Println("Error uploading file:", err.Error())
		return
	}

	defer file.Close()

	out, err := os.Create("images/" + header.Filename)
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println("Error copying file:", err.Error())
		return
	}

	c.JSON(200, "Image uploaded")

}

func Login(c *gin.Context) {
	email := strings.ToUpper(c.Param("email"))
	password := strings.ToUpper(c.Param("password"))
	var user models.User
	initializers.DB.Where("UPPER(email) = ? and UPPER(password) = ?", email, password).First(&user)
	c.JSON(200, &user)
}

func AdminUsers(c *gin.Context) {
	ty := c.Param("type")
	var users []models.User
	initializers.DB.Find(&users, "type = ?", ty)
	c.JSON(200, &users)
}

func SingleUser(c *gin.Context) {
	var user models.User
	initializers.DB.Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	initializers.DB.Where("id = ?", c.Param("id")).Updates(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	initializers.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, "User Deleted")
}

func AddCourse(c *gin.Context) {
	var course models.Course
	c.BindJSON(&course)
	initializers.DB.Create(&course)
	c.JSON(200, &course)
}


func DeleteCourse(c *gin.Context){
	var course models.Course
	initializers.DB.Where("id = ?",c.Param("id")).Delete(&course)
	c.JSON(200, "Course Deleted")
}

func GetCourses(c *gin.Context){
	var courses []models.Course
	initializers.DB.Find(&courses)
	c.JSON(200,&courses)
}


func AddEnrollment(c *gin.Context){
	var enrolment models.Enrollment
	c.BindJSON(&enrolment)
	initializers.DB.Create(&enrolment)
	c.JSON(200, &enrolment)
}

func GetEnrollment(c *gin.Context){
	var enrolments []models.Enrollment
	initializers.DB.Find(&enrolments)
	c.JSON(200, &enrolments)
}

func DeleteEnrollment(c *gin.Context){
	var enrolment models.Enrollment
	initializers.DB.Where("id = ?",c.Param("id")).Delete(&enrolment)
	c.JSON(200, "Enrollment Deleted")
}


func Allocate(c *gin.Context){
	var allocate models.Allocate
	c.BindJSON(&allocate)
	initializers.DB.Create(&allocate)
	c.JSON(200, &allocate)
}
