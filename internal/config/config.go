package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Db struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
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
