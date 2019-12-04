package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUser(c *gin.Context)
	Login(c *gin.Context)
	UpdateUser(c *gin.Context)
}
