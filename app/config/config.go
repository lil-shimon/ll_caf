package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		// finish program
		os.Exit(1)
	}

	Config = ConfigList{
		DbHost:     cfg.Section("db").Key("host").String(),
		DbPort:     cfg.Section("db").Key("port").MustInt(),
		DbUser:     cfg.Section("db").Key("user").String(),
		DbPassword: cfg.Section("db").Key("password").String(),
		DbName:     cfg.Section("db").Key("name").String(),
	}
}
