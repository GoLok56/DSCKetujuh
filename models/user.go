package models

type User struct {
	Username string `gorm:"primary_key"`
	Password string
	Nama     string
}
