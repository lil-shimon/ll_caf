package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

func GetItems() (model.Items, error) {
  i := model.Items{}

  if err := database.DB.Find(&i).Error; err != nil {
    return i, err
  }
  return i, nil
}

func StoreItem(i model.Item) (model.Item, error) {

  if err := database.DB.Create(&i).Error; err != nil {
    return i, err
  }

  return i, nil
}

func ShowItem(id string) (model.Item, error) {
  i := model.Item{}
  if err := database.DB.Where("id = ?", id).First(&i).Error; err != nil {
    return i, err
  }
  return i, nil
}

func UpdateItem(i model.Item) (model.Item, error) {
  if err := database.DB.Save(&i).Error; err != nil {
    return i, err
  }
  return i, nil
}
