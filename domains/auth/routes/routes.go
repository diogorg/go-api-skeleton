package routes

import (
	"api/domains/auth/controllers"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

var domain = "/auth"

func AuthRoutes(g *gin.Engine) {
	route := g.Group(domain)
	{
		route.POST("/login", controllers.Login)
		route.Use(middleware.UserAuth).GET("/me", controllers.Me)
	}
}
