package main

import (
	"fmt"
	"frankie/universal/sdk/initialize"

	"github.com/spf13/viper"
)

func main() {
	engine := initialize.NewEngine("config")
	engine.Run(fmt.Sprintf(":%v", viper.GetInt("SERVER_PORT")))
}
