package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type LogConfig struct {
	Path           string `json:"path" mapstructure:"path" yaml:"path"`
	Name           string `json:"name" mapstructure:"name" yaml:"name"`
	Level          string `json:"level" mapstructure:"level" yaml:"level"`
	MaxAge         int    `json:"max_age" mapstructure:"max_age" yaml:"max_age"`
	RotationTime   int    `json:"rotation_time" mapstructure:"rotation_time" yaml:"rotation_time"`
	CallerFullPath bool   `json:"caller_full_path" mapstructure:"caller_full_path" yaml:"caller_full_path"`
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
