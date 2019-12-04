package controllers

import (
	"github.com/gin-gonic/gin"
)

type TodoController interface {
	AddTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
	GetAllTodo(c *gin.Context)
}
