package router

import (
	controller "go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func Route(app *gin.Engine) {
	auth := app.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

}
