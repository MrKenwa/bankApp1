package config

import (
	"encoding/json"
	"os"
)

const configPath = "./config/config.json"

type Config struct {
	Server struct {
		Host string
	}
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
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
