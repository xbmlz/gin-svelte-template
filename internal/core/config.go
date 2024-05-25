package core

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Path           string `mapstructure:"path" yaml:"path"`
	Name           string `mapstructure:"name" yaml:"name"`
	Level          string `mapstructure:"level" yaml:"level"`
	MaxAge         int    `mapstructure:"max_age" yaml:"max_age"`
	RotationTime   int    `mapstructure:"rotation_time" yaml:"rotation_time"`
	CallerFullPath bool   `mapstructure:"caller_full_path" yaml:"caller_full_path"`
}

type HTTPConfig struct {
	Host string `mapstructure:"host" yaml:"host" validate:"ipv4"`
	Port int    `mapstructure:"port" yaml:"port" validate:"gte=1,lte=65535"`
}

type DBConfig struct {
	Driver string `mapstructure:"driver" yaml:"driver"`
	DSN    string `mapstructure:"dsn" yaml:"dsn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"pool_size" yaml:"pool_size"`
}

type AuthConfig struct {
	Enable          bool     `mapstructure:"enable" yaml:"enable"`
	TokenSecretKey  string   `mapstructure:"token_secret_key" yaml:"token_secret_key"`
	TokenExpireTime int64    `mapstructure:"token_expire_time" yaml:"token_expire_time"`
	IgnorePaths     []string `mapstructure:"ignore_paths" yaml:"ignore_paths"`
}

type Config struct {
	Log   *LogConfig   `mapstructure:"log" yaml:"log"`
	HTTP  *HTTPConfig  `mapstructure:"http" yaml:"http"`
	DB    *DBConfig    `mapstructure:"database" yaml:"database"`
	Redis *RedisConfig `mapstructure:"redis" yaml:"redis"`
	Auth  *AuthConfig  `mapstructure:"auth" yaml:"auth"`
}

var configPath = "config/config.yaml"

var defaultConfig = Config{
	Log: &LogConfig{
		Path:         "logs",
		Name:         "server",
		Level:        "debug",
		MaxAge:       24 * 7,
		RotationTime: 24,
	},
	HTTP: &HTTPConfig{
		Host: "0.0.0.0",
		Port: 8765,
	},
	DB: &DBConfig{
		Driver: "sqlite",
		DSN:    "./db.sqlite",
	},
	Redis: &RedisConfig{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
		PoolSize: 10,
	},
	Auth: &AuthConfig{
		Enable:          true,
		TokenSecretKey:  "secret",
		TokenExpireTime: 3600 * 24 * 7,
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

func (a *RedisConfig) Addr() string {
	if err := validator.New().Struct(a); err != nil {
		return defaultConfig.Redis.Addr()
	}

	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}
