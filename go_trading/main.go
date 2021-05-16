package main

import (
	"fmt"
	"gomodtest/go_trading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
