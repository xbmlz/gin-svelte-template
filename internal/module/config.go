package module

import (
	"fmt"

	"github.com/spf13/viper"
)

type LogConfig struct {
	Path  string `mapstructure:"path"`
	Name  string `mapstructure:"name"`
	Level string `mapstructure:"level"`
}

type Config struct {
	Log *LogConfig `mapstructure:"log"`
}

var configPath = "config.yaml"

var defaultConfig = Config{
	Log: &LogConfig{
		Path:  "logs",
		Name:  "app.log",
		Level: "debug",
	},
}

func LoadConfig() Config {
	config := defaultConfig

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return config
}
