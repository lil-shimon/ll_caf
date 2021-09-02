package service

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetLogs () (model.Logs, error) {
  ls, _ := repository.GetLogs()
  for _, l := range ls {
    l.Item, _ = repository.ShowItem(strconv.Itoa(l.ItemId))
  }
  return ls, nil
}

func GetLog (id string)  (model.Log, error) {
  l, _ := repository.GetLog(id)
  l.Item, _ = repository.ShowItem(strconv.Itoa(l.ItemId))
  return l, nil
}

func StoreLog (l model.Log) (model.Log, error) {
  l, _ = repository.StoreLog(l)
  return l, nil
}

func UpdateLog (c echo.Context) (model.Log, error) {
  l, _ := repository.GetLog(c.Param("id"))
  c.Bind(&l)
  l, _ = repository.UpdateLog(l)
  return l, nil
}

func GetLogsByItemId (itemId string) (model.Logs, error) {
  ls, _ := repository.GetLogsByItemId(itemId)
  for _, l := range ls {
    l.Item, _ = repository.ShowItem(itemId)
  }
  return ls, nil
}
