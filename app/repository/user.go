package repository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

func GetUser(id string) (model.User, error) {
	u := model.User{}
	if err := database.DB.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func StoreUser(u model.User) (model.User, error) {

	if err := database.DB.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
