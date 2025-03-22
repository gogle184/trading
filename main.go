package main

import (
	"fmt"
	"trading/config"
	"trading/utils"
	"trading/bitflyer"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
