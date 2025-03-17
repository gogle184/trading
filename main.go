package main

import (
	"fmt"
	"log"
	"trading/config"
	"trading/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	log.Println("test")
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
