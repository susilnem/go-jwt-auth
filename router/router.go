package router

import (
	"go-jwt/controller"

	"github.com/gin-gonic/gin"
)

func Route(app *gin.Engine) {
	auth := app.Group("/auth")
	{
		auth.POST("/register", controller.Register)
	}

}
