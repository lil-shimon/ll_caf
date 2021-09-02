package model

import "gorm.io/gorm"

type Log struct {
  gorm.Model
  ItemId int `json:"itemId" gorm:"not null"`
  Item Item
}

// daily log
// Items 
// TAmount Total Amount
// TContain Total Contain
// CAmount Compared Amount
// CContain Compared Contain
// CARate Compared Amount Rate
// CCRate Compared Contain Rate
type DayLog struct {
  Items Items
  TAmount int
  TContain int
  CAmount int
  CContain int
  CARate int
  CCRate int
}

type Logs []*Log

