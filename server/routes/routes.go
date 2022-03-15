package routes

import (
	auth "api/domains/auth/routes"
	user "api/domains/users/routes"

	"github.com/gin-gonic/gin"
)

func Handle(g *gin.Engine) {
	auth.AuthRoutes(g)
	user.UserRoutes(g)
}
