package main

import (
	"api/controller"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.ConnectDb()
}

func main() {

	r := gin.Default()
	r.POST("/addUser", controller.AddUser)
	r.GET("/login/:email/:password", controller.Login)
	r.GET("/adminUsers/:type", controller.AdminUsers)
	r.Run(":3000")
}
