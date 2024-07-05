package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"os"
)

const (
	configPath = "./config/config.json"
	keyPath    = "./config/ec_key.pem"
)

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

	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
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

	pem, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	c.PrivateKey, err = jwt.ParseECPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}

	c.PublicKey = &c.PrivateKey.PublicKey

	return c, nil
}
