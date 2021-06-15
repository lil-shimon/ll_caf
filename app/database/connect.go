package database

import (
	"fmt"
	"os"

	//"strconv"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/lil-shimon/workManager/config"
)

var DB *gorm.DB

func Connect() {
	_, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("error occurred")
		os.Exit(1)
	}

	//host := config.Config.DbHost
	//port := config.Config.DbPort
	//user := config.Config.DbName
	//password := config.Config.DbPassword
	//databaseName := config.Config.DbName

	//dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	/**
	/* 直書きから上記変数に戻す
	*/
	dsn := "root:root@tcp(db:3306)/workManager?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
