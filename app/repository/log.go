package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
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
