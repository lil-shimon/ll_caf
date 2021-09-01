package service

import (
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetItems() (model.Items,  error) {
  i, _ := repository.GetItems()
  return i, nil
}
