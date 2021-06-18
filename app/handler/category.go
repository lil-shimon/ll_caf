package handler

import (
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/database"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/regository"
	"net/http"
)

func CreateCategory(c echo.Context) error {
	cat := model.Category{}
	if err := c.Bind(&cat); err != nil {
		return err
	}
	database.DB.Create(&cat)
	return c.JSON(http.StatusCreated, cat)
}

func GetCategories(c echo.Context) error {
	categories, _ := regository.GetRepoCategories()
	return c.JSON(http.StatusOK, categories)
}

func ShowCategory(c echo.Context) error {
	id := c.Param("id")
	category, _ := regository.ShowRepoCategory(string(id))
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	cat := model.Category{}
	if err := c.Bind(&cat); err != nil {
		return err
	}
	database.DB.Save(&cat)
	return c.JSON(http.StatusOK, cat)
}

func DeleteCategory(c echo.Context) error {
	cat := model.Category{}
	if err := c.Bind(&cat); err != nil {
		return err
	}
	database.DB.Delete(&cat)
	return c.JSON(http.StatusNoContent, cat)
}
