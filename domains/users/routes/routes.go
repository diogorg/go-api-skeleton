package routes

import (
	"api/domains/users/controllers"

	"github.com/gin-gonic/gin"
)

var domain = "/users"

func UserRoutes(g *gin.Engine) {
	route := g.Group(domain)
	{
		route.GET("", controllers.Index)
		route.GET("/:id", controllers.Show)
		route.GET("/find/email/:email", controllers.Find)
		route.POST("", controllers.Store)
		route.DELETE("/:id", controllers.Delete)
		route.PATCH("/:id", controllers.Update)
	}
}
