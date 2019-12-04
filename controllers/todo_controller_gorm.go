package controllers

import (
	"net/http"

	"dscketujuh/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TodoControllerGorm struct {
	Db *gorm.DB
}

func (todoController TodoControllerGorm) AddTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal menambahkan todo baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	todoController.Db.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"pesan":     "Berhasil menambahkan todo baru",
		"todo_baru": todo,
	})
}

func (todoController TodoControllerGorm) UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal update todo, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	idTodo := c.Param("idTodo")

	var todo models.Todo
	if todoController.Db.Where("id_todo = ?", idTodo).First(&todo).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan todo dengan id " + idTodo,
		})
		return
	}

	todoController.Db.Model(&todo).Where("id_todo = ?", idTodo).Updates(updatedTodo)
	c.JSON(http.StatusOK, gin.H{
		"pesan":     "Berhasil melakukan update user",
		"todo_baru": todo,
	})
}

func (todoController TodoControllerGorm) DeleteTodo(c *gin.Context) {
	idTodo := c.Param("idTodo")

	var todo models.Todo
	if todoController.Db.Where("id_todo = ?", idTodo).First(&todo).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data todo dengan id " + idTodo,
		})
		return
	}

	todoController.Db.Model(&todo).Where("id_todo = ?", idTodo).Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"pesan":         "Berhasil melakukan update todo",
		"pengguna_baru": todo,
	})
}

func (todoController TodoControllerGorm) GetAllTodo(c *gin.Context) {
	var todos []models.Todo
	todoController.Db.Find(&todos)
	c.JSON(http.StatusCreated, gin.H{
		"pesan": "Berhasil mendapatkan semua todo",
		"todo":  todos,
	})
}
