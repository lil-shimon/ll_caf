package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/util"
)

func GetLogs() (model.Logs, error) {
  l := model.Logs{}
  if err := database.DB.Find(&l).Error; err != nil {
    return l, err
  }
  return l, nil
}

func GetLog(id string) (model.Log, error) {
  l := model.Log{}
  if err := database.DB.Where("id = ?", id).First(&l).Error; err != nil {
    return l, err
  }

  return l, nil
}

func StoreLog(l model.Log) (model.Log, error) {
  if err := database.DB.Create(&l).Error; err != nil {
    return l, err
  }
  return l, nil
}

func UpdateLog(l model.Log) (model.Log, error) {
  if err := database.DB.Save(&l).Error; err != nil {
    return l, err
  }
  return l, nil
}

func GetLogsByItemId(itemId string) (model.Logs, error) {
  ls := model.Logs{}
  if err := database.DB.Where("item_id = ?", itemId).Find(&ls).Error; err != nil {
    return ls, err
  }
  return ls, nil
}

func GetDailyLog() (model.Logs, error) {
  ls := model.Logs{}
  td, tml := util.GetDaily()
  if err := database.DB.Where("created_at BETWEEN ? AND ?", td, tml).Find(&ls).Error; err != nil {
    return ls, err
  }

  return ls, nil
}
