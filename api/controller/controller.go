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
