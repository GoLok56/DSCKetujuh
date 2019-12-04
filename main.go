package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"dscketujuh/controllers"
	"dscketujuh/models"
	"fmt"
)

var Db *gorm.DB

var userController controllers.UserController
var todoController controllers.TodoController

func main() {
	var err error
	Db, err = gorm.Open("mysql", "golok:112156@/golangdsc?charset=utf8&parseTime=True&loc=Local")
	defer Db.Close()

	if err != nil {
		fmt.Errorf("terjadi kesalahan saat membuka koneksi ke server mysql")
		return
	}
	fmt.Println("Berhasil melakukan koneksi ke server MySQL...")

	Db.AutoMigrate(&models.User{}, &models.Todo{})

	router := gin.Default()

	userController = controllers.UserControllerGorm{Db}
	todoController = controllers.TodoControllerGorm{Db}

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userController.AddUser)
		userGroup.POST("/login", userController.Login)
		userGroup.PUT("/:username", userController.UpdateUser)
	}

	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/", todoController.GetAllTodo)
		todoGroup.POST("/", todoController.AddTodo)
		todoGroup.PUT("/:idTodo", todoController.UpdateTodo)
		todoGroup.DELETE("/:idTodo", todoController.DeleteTodo)
	}

	router.Run()
}
