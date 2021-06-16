package model

// Types must be the same as the table name
type Types struct {
	Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}
