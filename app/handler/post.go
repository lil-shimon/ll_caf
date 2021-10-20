package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/service"
)

/// ユーザーごとの投稿を取得
func GetPostsByUserId(c echo.Context) error {
	p, _ := service.GetPostsByUserId(c.Param("id"))
	return c.JSON(http.StatusOK, p)
}
