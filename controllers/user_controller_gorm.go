package controllers

import (
	"net/http"

	"dscketujuh/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserControllerGorm struct {
	Db *gorm.DB
}

func (userController UserControllerGorm) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal menambahkan pengguna baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	userController.Db.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"pesan":         "Berhasil menambahkan pengguna baru",
		"pengguna_baru": user,
	})
}

func (userController UserControllerGorm) Login(c *gin.Context) {
	var body gin.H
	c.BindJSON(&body)

	var user models.User
	if userController.Db.Where("username = ?", body["username"]).First(&user).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data pengguna dengan username " + body["username"].(string),
		})
		return
	}

	if user.Password != body["password"] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"pesan": "Password yang dimasukkan salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan":         "Berhasil melakukan login",
		"pengguna_baru": user,
	})
}

func (userController UserControllerGorm) UpdateUser(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "Gagal update pengguna baru, pastikan untuk mengisi semua parameter yang dibutuhkan",
		})
		return
	}

	username := c.Param("username")

	var user models.User
	if userController.Db.Where("username = ?", username).First(&user).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "Tidak ditemukan data pengguna dengan username " + username,
		})
		return
	}

	userController.Db.Model(&user).Where("username = ?", username).Updates(updatedUser)
	c.JSON(http.StatusOK, gin.H{
		"pesan":         "Berhasil melakukan update user",
		"pengguna_baru": updatedUser,
	})
}
