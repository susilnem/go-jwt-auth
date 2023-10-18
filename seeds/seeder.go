package main

import (
	"fmt"
	models "go-jwt/models"
	"go-jwt/database"
)

func main() {
	db := database.Database

	// Create and seed roles
	roles := []models.Role{
		{
			Name:        "Admin",
			Description: "Administrator role",
		},
		{
			Name:        "User",
			Description: "User role",
		},
	}
	for _, role := range roles {
		result := db.Create(&role)
		if result.Error != nil {
			fmt.Printf("Error seeding roles: %v\n", result.Error)
		}
	}

	// Create and seed permissions
	permissions := []models.Permission{
		{
			Name:        "Create",
			Description: "Create permission",
		},
		{
			Name:        "Read",
			Description: "Read permission",
		},
		{
			Name:        "Update",
			Description: "Update permission",
		},
		{
			Name:        "Delete",
			Description: "Delete permission",
		},
	}
	for _, permission := range permissions {
		result := db.Create(&permission)
		if result.Error != nil {
			fmt.Printf("Error seeding permissions: %v\n", result.Error)
		}
	}

	// Create and seed a sample user
	user := models.User{
		Username: "adminuser",
		Password: "password123", // You should hash the password in a real application
		Roles:    roles,
		Permissions: []models.Permission{permissions[0], permissions[1]},
	}

	result := db.Create(&user)
	if result.Error != nil {
		fmt.Printf("Error seeding user: %v\n", result.Error)
	} else {
		fmt.Println("Seed data successfully created.")
	}
}
