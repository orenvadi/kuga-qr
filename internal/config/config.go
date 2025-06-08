package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(configPath string) *Config {
	var cfg Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %v", err))
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("failed to parse YAML: %v", err))
	}
	return &cfg
}
