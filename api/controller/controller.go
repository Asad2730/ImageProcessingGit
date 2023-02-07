package controller

import (
	"api/initializers"
	"api/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	initializers.DB.Create(&user)
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	email := strings.ToUpper(c.Param("email"))
	password := strings.ToUpper(c.Param("password"))
	var user models.User
	initializers.DB.Where("UPPER(email) = ? and UPPER(password) = ?", email, password).First(&user)
	c.JSON(200, &user)
}
