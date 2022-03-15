package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonStatus(c *gin.Context, obj interface{}, statusCode int) {
	c.JSON(statusCode, WrapData(obj))
}

func JsonOk(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, WrapData(obj))
}

func JsonAccepted(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusAccepted, WrapData(obj))
}

func JsonCreated(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusCreated, WrapData(obj))
}

func JsonBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func JsonNotFound(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": msg,
	})
}

func JsonUnprocessableEntity(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"error": err.Error(),
	})
}

func JsonUnauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

func WrapData(obj interface{}) gin.H {
	return gin.H{
		"data": obj,
	}
}
