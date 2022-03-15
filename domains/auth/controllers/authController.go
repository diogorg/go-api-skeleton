package controllers

import (
	"api/domains/auth/models"
	"api/domains/auth/services"
	"api/support"
	"api/support/response"
	"api/support/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var authLogin models.AuthLogin
	err := c.ShouldBindJSON(&authLogin)
	support.Panic(err)
	token, err := services.Login(authLogin)
	if err != nil {
		response.JsonUnauthorized(c)
	} else {
		response.JsonOk(c, token)
	}
}

func Me(c *gin.Context) {
	token := token.GetToken(c)
	user := services.Me(token)
	response.JsonOk(c, user)
}
