package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//setConfig used for setup application configuration
func setConfig(config string) {
	viper.SetConfigName(config)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Configuration file not found: %s", err))
	}

	viper.AutomaticEnv()

	if viper.GetString("APP") == "PRD" {
		gin.SetMode(gin.ReleaseMode)
	}
}
