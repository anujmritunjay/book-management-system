package main

import (
	"github.com/anujmritunjay/book-management-system/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Auto Migrate the schema
	config.DB.AutoMigrate(&User{})

	// Perform CRUD operations...
	config.DB.Create(&User{Name: "John Doe", Email: "john.doe@example.com"})

	// Example of reading data
	var user User
	config.DB.First(&user, 1)                      // Find user with integer primary key
	config.DB.First(&user, "name = ?", "John Doe") // Find user with name
}
