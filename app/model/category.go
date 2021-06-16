package model

// Category struct must be the same as the table name
type Category struct {
	Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT" param:"id"`
	Name string `json:"name"`
}

type Categories []*Category
