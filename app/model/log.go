package model

import "gorm.io/gorm"

type Log struct {
  gorm.Model
  ItemId int `json:"itemId" gorm:"not null"`
  Item Item
}

type Logs []*Log

// daily log
// Items 
// TAmount Total Amount
// TContain Total Contain
// CAmount Compared Amount
// CContain Compared Contain
// CARate Compared Amount Rate
// CCRate Compared Contain Rate
type DayLog struct {
  Logs Logs
  Items Items
  TAmount int
  TContain int
  // CAmount int
  // CContain int
  // CARate int
  // CCRate int
}

