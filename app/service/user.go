package service

import (
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetUser(id string) (model.User, error) {
	i, _ := repository.GetUser(id)
	return i, nil
}

func StoreUser(u model.User) (model.User, error) {
	return repository.StoreUser(u)
}
