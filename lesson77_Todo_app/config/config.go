package config

import (
	"lesson77_Todo_app/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

//iniファイルをconfig.goで読みこむ

type ConfigList struct {
	Port      string //ポート番号
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

//ConfigListを外部のパッケージからでも呼び出せるようにグローバルに変数宣言
var Config ConfigList

//main関数より前に実行したい
//init関数はmain関数より先に呼び出される
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

//iniファイルを読み込み、ConfitList構造体に値を設定する関数の作成
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(), //app/viewsの静的ファイルcss、jsを設定

	}
}
