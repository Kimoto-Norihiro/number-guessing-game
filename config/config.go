package config

import(
	"log"
	"fmt"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	Static string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig(){  //config.iniからデータを取得
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("config")
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		Static: cfg.Section("web").Key("static").String(),
	}
}