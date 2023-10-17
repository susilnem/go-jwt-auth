package main

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "go-jwt/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)


// @title Jwt Authentication API
// @description This is a sample jwt authentication api
// @BasePath /api
// @version 1
//
// @contact.name				For API Support
// @contact.email				susiltiwari750@gmail.com
//
// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
//
// @BasePath					/api
// @SecurityDefinitions.apikey	BearerAuth
// @Name						Authorization
// @In							header
// @Description				Add prefix of Bearer before  token Ex: "Bearer token"
// @Query.collection.format	multi
func main() {
	// load env variables
	loadEnv()

	// connect to database
	database.Connect()

	app := gin.Default()

	app.Use(gin.Logger())

	// set up routes
	router.Route(app)
	router.GetRoute(app)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	
	port := os.Getenv("SERVER_PORT")
	
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// run server on port 8080
	app.Run(port)
	fmt.Println("Server running on " + port + " port")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
