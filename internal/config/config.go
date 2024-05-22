package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Path           string `mapstructure:"Path" yaml:"path"`
	Name           string `mapstructure:"Name" yaml:"name"`
	Level          string `mapstructure:"Level" yaml:"level"`
	MaxAge         int    `mapstructure:"MaxAge" yaml:"max_age"`
	RotationTime   int    `mapstructure:"RotationTime" yaml:"rotation_time"`
	CallerFullPath bool   `mapstructure:"CallerFullPath" yaml:"caller_full_path"`
}

type HTTPConfig struct {
	Host string `mapstructure:"Host" yaml:"host" validate:"ipv4"`
	Port int    `mapstructure:"Port" yaml:"port" validate:"gte=1,lte=65535"`
}

type Config struct {
	Log  *LogConfig  `mapstructure:"Log"`
	HTTP *HTTPConfig `mapstructure:"HTTP"`
}

var configPath = "config.yaml"

var defaultConfig = Config{
	Log: &LogConfig{
		Path:  "logs",
		Name:  "app.log",
		Level: "debug",
	},
	HTTP: &HTTPConfig{
		Host: "0.0.0.0",
		Port: 8765,
	},
}

func NewConfig() Config {
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

func (a *HTTPConfig) ListenAddr() string {
	if err := validator.New().Struct(a); err != nil {
		return defaultConfig.HTTP.ListenAddr()
	}

	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}
