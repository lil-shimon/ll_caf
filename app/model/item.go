package model

import "gorm.io/gorm"

type Item struct {
  gorm.Model
  Name string `json:"name"`
  Amount int `json:"amount" gorm:"default:0"`
  Contain int `json:"contain" gorm:"default:0"`
  TypeId int `json:"typeId" gorm:"not null"`
}

type Items []*Item
