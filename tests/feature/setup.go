package tests

import (
	"api/server"

	"github.com/gin-gonic/gin"
)

func SetupFeatureTests() *gin.Engine {
	routes := server.CreateForFeatureTests()

	return routes
}
