package main

import (
	"api/initializers"
	"api/models"
)

func init() {

	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Allocate{},
		&models.Attendance{},
		&models.Enrollment{},
		&models.Course{},
	)
}
