package models

type Todo struct {
	IdTodo    int `gorm:"primary_key;AUTO_INCREMENT"`
	Tugas     string
	Deskripsi string
	Deadline  string
	Status    bool `gorm:"default:false"`
}
