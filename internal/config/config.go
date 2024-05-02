package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

func Load() {
	viper.AddConfigPath("$APP_HOME")
	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		slog.Warn(fmt.Sprintf("fail to read config from file: %s", err.Error()))
		slog.Info("Continue to load config from env")
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	GetAppConfig()
}
