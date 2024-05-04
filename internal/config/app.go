package config

import (
	"github.com/spf13/viper"
	"log"

	"github.com/quocanh1897/sample-gin-server/internal/utils"
)

type AppConfig struct {
	ServiceName string
	Log         LoggingConfig `json:"log"`
	Env         string        `json:"env"`
	Http        HTTPConfig    `json:"http"`
}

func GetAppConfig() AppConfig {
	validate := utils.GetValidatorInstance()
	config := AppConfig{
		ServiceName: "sample_go_server",
		Env:         viper.GetString("app.env"),
		Http:        GetHTTPConfig(),
		Log:         GetLoggingConfig(),
	}
	if err := validate.Struct(&config); err != nil {
		log.Fatalf(err.Error())
	}
	return config
}

type HTTPConfig struct {
	Port uint64 `json:"port"`
}

func GetHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Port: viper.GetUint64("http.port"),
	}
}
