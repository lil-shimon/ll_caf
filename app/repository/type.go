package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

func GetTypes() (model.Types, error) {
  t := model.Types{}

  if err := database.DB.Find(&t).Error; err != nil {
    return t, err
  }
  return t, nil
}

func ShowType(id string) (model.Type, error) {
  t := model.Type{}
  if err := database.DB.Where("id = ?", id).First(&t).Error; err != nil {
    return t, err
  }
  return t, nil
}

func StoreType(t model.Type) (model.Type, error) {
  if err := database.DB.Create(&t).Error; err != nil {
    return t, err
  }
  return t, nil
}

func UpdateType(t model.Type) (model.Type, error) {
  if err := database.DB.Save(&t).Error; err != nil {
    return t, err
  }
  return t, nil
}
