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
