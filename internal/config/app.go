package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"

	"github.com/quocanh1897/sample-gin-server/internal/utils"
)

type AppConfig struct {
	Service ServiceConfig `json:"service"`

	ServiceName string
	Tls         TlsConfig     `json:"tls" validate:"required"`
	Log         LoggingConfig `json:"log"`
	Env         string        `json:"env"`
	Http        HTTPConfig    `json:"http"`
}

type TlsConfig struct {
	Key  string `json:"key"`
	Cert string `json:"cert"`
}

type ZooKeeperConfig struct {
	Env   string   `json:"env" validate:"required"`
	Hosts []string `json:"hosts" validate:"required"`
}

func GetAppConfig() AppConfig {
	validate := utils.GetValidatorInstance()
	config := AppConfig{
		Service: GetServiceConfig(),

		ServiceName: "sample_go_server",
		Env:         viper.GetString("app.env"),
		Tls: TlsConfig{
			Key:  viper.GetString("tls.key"),
			Cert: viper.GetString("tls.cert"),
		},
		Http: GetHTTPConfig(),
		Log:  GetLoggingConfig(),
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

type ServiceConfig struct {
	DisabledFeatures []string
}

func GetServiceConfig() ServiceConfig {
	viper.SetDefault("app.disabled_features", "")

	disabledFeatures := []string{}
	for _, v := range strings.Split(viper.GetString("app.disabled_features"), ",") {
		if disabledFeature := strings.TrimSpace(v); disabledFeature != "" {
			disabledFeatures = append(disabledFeatures, disabledFeature)
		}
	}

	return ServiceConfig{
		DisabledFeatures: disabledFeatures,
	}
}
