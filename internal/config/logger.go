package config

import "github.com/spf13/viper"

type LoggingFormatter string

var (
	LoggingJSONFormatter LoggingFormatter = "json"
	LoggingTextFormatter LoggingFormatter = "text"
)

type LoggingLevel string

var (
	LoggingInfoLevel  LoggingLevel = "info"
	LoggingDebugLevel LoggingLevel = "debug"
)

type LoggingConfig struct {
	Formatter LoggingFormatter `json:"formatter"`
	Level     LoggingLevel     `json:"level"`
}

func GetLoggingConfig() LoggingConfig {
	env := viper.GetString("app.env")

	switch env {
	case "production":
		return LoggingConfig{
			Formatter: LoggingJSONFormatter,
			Level:     LoggingInfoLevel,
		}

	case "staging":
		return LoggingConfig{
			Formatter: LoggingJSONFormatter,
			Level:     LoggingInfoLevel,
		}

	default:
		return LoggingConfig{
			Formatter: LoggingTextFormatter,
			Level:     LoggingDebugLevel,
		}
	}
}
