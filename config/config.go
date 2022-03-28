package config

import (
	"os"

	"github.com/spf13/cast"
)

var PerPageSize = 100

// Config ...
type Config struct {
	Environment string // develop, staging, production

	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
