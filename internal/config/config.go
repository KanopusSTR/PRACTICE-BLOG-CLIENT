package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func GetURL() (string, error) {
	if url := os.Getenv("HTTP_URL"); url == "" {
		return "", errors.New("url not set")
	} else {
		return url, nil
	}
}
