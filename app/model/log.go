package model

import "gorm.io/gorm"

type Log struct {
  gorm.Model
  ItemId int `json:"itemId" gorm:"not null"`
}

type Logs []*Log