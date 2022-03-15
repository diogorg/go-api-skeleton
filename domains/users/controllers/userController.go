package controllers

import (
	"api/domains/users/adapters"
	"api/domains/users/repositories/providers"
	"api/domains/users/resources"
	"api/support/cast"
	"api/support/response"

	"github.com/gin-gonic/gin"
)

var repository = providers.GetUserRespository()

func Index(c *gin.Context) {
	users := repository.GetAll()
	response.JsonOk(c, resources.GetUsersResource(users))
}

func Show(c *gin.Context) {
	id := c.Param("id")
	user := repository.FindById(cast.StringToInt(id))
	checkUserNotFound(c, int(user.ID))
	if c.IsAborted() == false {
		response.JsonOk(c, resources.GetUserResource(user))
	}
}

func Store(c *gin.Context) {
	user, err := adapters.Store(c)
	checkUnprocessableEntity(c, err)
	if c.IsAborted() == false {
		repository.Store(&user)
		response.JsonCreated(c, user.ID)
	}
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	repository.Delete(cast.StringToInt(id))
	response.JsonOk(c, nil)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	user := repository.FindById(cast.StringToInt(id))
	checkUserNotFound(c, int(user.ID))
	updated, err := adapters.Update(c)
	checkUnprocessableEntity(c, err)
	if c.IsAborted() == false {
		repository.Update(int(user.ID), &updated)
		response.JsonOk(c, &updated)
	}
}

func Find(c *gin.Context) {
	email := c.Param("email")
	user := repository.FindByEmail(email)
	checkUserNotFound(c, int(user.ID))
	if c.IsAborted() == false {
		response.JsonOk(c, resources.GetUserResource(user))
	}
}

func checkUserNotFound(c *gin.Context, id int) {
	if id == 0 {
		response.JsonNotFound(c, "User Not Found")
	}
}

func checkUnprocessableEntity(c *gin.Context, err error) {
	if err != nil {
		response.JsonUnprocessableEntity(c, err)
	}
}
