package main

import (
	controller "api/controller/admin"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.ConnectDb()
}

func main() {

	r := gin.Default()
	r.POST("/addUser", controller.AddUser)
	r.POST("/uploadImage", controller.UploadImage)
	r.GET("/login/:email/:password", controller.Login)
	r.GET("/adminUsers/:type", controller.AdminUsers)
	r.GET("/singleUser/:id", controller.SingleUser)
	r.PUT("/updateUser/:id", controller.UpdateUser)
	r.GET("/images/:filename", controller.GetImage)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)
	r.GET("/getCount", controller.GetCount)
	r.POST("/addCourse",controller.AddCourse)
	r.DELETE("/deleteCourse/:id", controller.DeleteCourse)
	r.GET("/getCourses",controller.GetCourses)
	r.Run(":3000")
}
