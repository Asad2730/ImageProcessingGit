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
	r.POST("/signup", controller.SignUp)
	r.GET("/login/:email/:password", controller.Login)
	r.Run(":3000")
}
