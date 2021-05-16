package main

import (
	"fmt"
	"gomodtest/go_trading/config"
	"gomodtest/go_trading/utils"
	"log"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
	utils.LogginSettings(config.Config.LogFile)
	log.Println("test")
}
