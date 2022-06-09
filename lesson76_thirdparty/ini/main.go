package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

//ini

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(8080), //値が存在しない場合は8080が入る

		DbName: cfg.Section("db").Key("name").MustString("example.sql"),

		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %T %v\n", Config.Port, Config.Port)
	fmt.Printf("DbName = %T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("SQLDriver = %T %v\n", Config.SQLDriver, Config.SQLDriver)
}
