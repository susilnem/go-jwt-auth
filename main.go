package main

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	docs "go-jwt/docs"
)

func main() {
	// load env variables
	loadEnv()

	// connect to database
	database.Connect()

	app := gin.Default()

	// set up routes
	docs.SwaggerInfo.BasePath = "/api"
	router.Route(app)
	router.GetRoute(app)
	
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
