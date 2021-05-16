package main

import (
	"fmt"
	"gomodtest/go_trading/bitflyer"
	"gomodtest/go_trading/config"
	"gomodtest/go_trading/utils"
)

func main() {
	// fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
	utils.LogginSettings(config.Config.LogFile)
	// log.Println("test")
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
