package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini") // config読み込み
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1) // エラーコード1で抜ける
	}
	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
	}
}
