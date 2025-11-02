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
	Jwt    Jwt    `yaml:"jwt"`
}
type Jwt struct {
	Secret   string        `yaml:"secret"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}

type Server struct {
	Port    string        `yaml:"port"`
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

func (c *Config) DbUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Db.Host, c.Db.Port, c.Db.Username, c.Db.Password, c.Db.DBName, c.Db.SSLMode)
}

func MustLoad(configPath string) Config {
	var cfg Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %v", err))
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("failed to parse YAML: %v", err))
	}
	return cfg
}
