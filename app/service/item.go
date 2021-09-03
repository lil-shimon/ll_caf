package service

import (
	"strconv"

	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetItems() (model.Items,  error) {
  is, _ := repository.GetItems()
  for _, i := range is {
    i.Type, _ = repository.ShowType(strconv.Itoa(i.TypeId))
  }
  return is, nil
}

func StoreItem(t model.Item) (model.Item, error) {
  i, _ := repository.StoreItem(t)
  return i, nil
}

func ShowItem(id string) (model.Item, error) {
  i, _ := repository.ShowItem(id)
  i.Type, _ = repository.ShowType(strconv.Itoa(i.TypeId))
  return i, nil
}

func UpdateItem(i model.Item) (model.Item, error) {
  i, _ = repository.UpdateItem(i)
  return i, nil
}

func GetItemsByTypeId(tid string) (model.Items, error) {
  is, err := repository.GetItemsTypeId(tid)
  if err != nil {
    return is, err
  }
  return is, nil
}

