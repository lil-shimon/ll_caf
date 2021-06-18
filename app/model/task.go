package model

type Task struct {
	Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT" param:"id"`
	CategoryId int `json:"category_id"`
	Name string `json:"name"`
	Content string `json:"content"`
}