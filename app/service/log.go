package service

import (
	"log"
	"reflect"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/model"
	"github.com/lil-shimon/workManager/repository"
)

func GetLogs() (model.Logs, error) {
	ls, _ := repository.GetLogs()
	for _, l := range ls {
		l.Item, _ = repository.ShowItem(strconv.Itoa(l.ItemId))
	}
	return ls, nil
}

func GetLog(id string) (model.Log, error) {
	l, _ := repository.GetLog(id)
	l.Item, _ = repository.ShowItem(strconv.Itoa(l.ItemId))
	return l, nil
}

func StoreLog(l model.Log) (model.Log, error) {
	l, _ = repository.StoreLog(l)
	return l, nil
}

func UpdateLog(c echo.Context) (model.Log, error) {
	l, _ := repository.GetLog(c.Param("id"))
	c.Bind(&l)
	l, _ = repository.UpdateLog(l)
	return l, nil
}

func GetLogsByItemId(itemId string) (model.Logs, error) {
	ls, _ := repository.GetLogsByItemId(itemId)
	for _, l := range ls {
		l.Item, _ = repository.ShowItem(itemId)
	}
	return ls, nil
}

func GetDailyLog() (model.DayLog, error) {
	dls := model.DayLog{}
	dls.Logs, _ = repository.GetDailyLog()
	var itemIds []string
	for _, l := range dls.Logs {
		itemIds = append(itemIds, strconv.Itoa(l.ItemId))
		l.Item, _ = repository.ShowItem(strconv.Itoa(l.ItemId))
	}
	dls.Items, _ = repository.GetItemByIds(itemIds)
	for _, dl := range dls.Items {
		dls.TAmount += dl.Amount
		dls.TContain += dl.Contain
	}
	return dls, nil
}

func GetLogByType(tid string) (model.Logs, error) {
	ls := model.Logs{}

	is, err := repository.GetItemsTypeId(tid)
	if err != nil {
		return ls, err
	}

	for _, i := range is {
		log.Println(reflect.TypeOf(i.ID))
	}
	return ls, nil
}
