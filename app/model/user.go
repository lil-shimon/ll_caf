package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
	Bio      string `json:bio`
	contact  string `json:contact`
	Website  string `json:website`
	Image    string `json:image`
}
