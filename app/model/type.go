package model

type TaskType struct {
	Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}
