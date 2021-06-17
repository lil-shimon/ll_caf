package regository

import (
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
)

// GetRepoCategories
// return [type] Categories, error
///**
func GetRepoCategories() (model.Categories, error) {
	cats := model.Categories{}

	if err := database.DB.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

// ShowRepoCategory
// return [type] Category, error
///**
func ShowRepoCategory(id string) (model.Category, error) {
	cat := model.Category{}

	if err := database.DB.Where("id = ?", id).First(&cat).Error; err != nil {
		return cat, err
	}
	return cat, nil
}

func UpdateRepoCategory(id string) (model.Category, error) {
	cat, _ := ShowRepoCategory(id)
	req := model.Category{}
	cat = req

	if err := database.DB.Save(&cat).Error; err != nil {
		return cat, err
	}
	return cat, nil
}
