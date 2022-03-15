package server

import (
	"api/server/routes"
	"api/support"

	"github.com/gin-gonic/gin"
)

func Run() {
	config := support.LoadConfig()
	g := create()
	g.Run(config.ServerPort)
}

func create() *gin.Engine {
	g := gin.Default()
	routes.Handle(g)

	return g
}

func CreateForFeatureTests() *gin.Engine {
	//g := gin.New()
	g := gin.Default()
	routes.Handle(g)

	return g
}
