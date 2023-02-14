package controller

import (
	"api/initializers"
	"api/models"
	"strings"

	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
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

	var user models.User
	user.Image = "images/" + header.Filename
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

func AdminUsers(c *gin.Context) {
	ty := c.Param("type")
	var user models.User
	fmt.Println("TYPE", ty)
	initializers.DB.Where("type = ?", ty).Select(&user)
	c.JSON(200, &user)
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
