package config

import (
	"bankApp1/internal/models"
	"encoding/json"
	"os"
)

const configPath = "./config/config.json"

type Config struct {
	Server struct {
		Host   string
		Domain string
	}
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}
	SessionSettings struct {
		SessionTTLSeconds models.TTL
	}
}

func LoadConfig() (*Config, error) {
	var c *Config
	jsonFile, err := os.Open(configPath)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(jsonFile).Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
