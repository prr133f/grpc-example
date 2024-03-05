package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type PGConfig struct {
	DSN string `yaml:"dsn"`
}

type AppConfig struct {
	LogLevel string `yaml:"log_level"`
	Port     int    `yaml:"port"`
}

type Config struct {
	PGConfig  `yaml:"postgres"`
	AppConfig `yaml:"app"`
}

var (
	instance *Config
	once     sync.Once
)

func ParseConfig() *Config {
	once.Do(func() {
		configFile, err := os.ReadFile("/app/config/config.yaml")
		if err != nil {
			panic(err)
		}

		var appConfig Config
		if err := yaml.Unmarshal(configFile, &appConfig); err != nil {
			panic(err)
		}

		instance = &appConfig
	})
	return instance
}
