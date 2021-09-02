package model

import "gorm.io/gorm"

type Type struct {
  gorm.Model
  Name string `json:"name" gorm:"not null"`
  Items []Item
}

type Types []*Type
