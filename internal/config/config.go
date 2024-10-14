package config

import (
	"fmt"
	"os"
)

type Config struct {
	BaseURL          string
	TokenID          string
	Token            string
	AuthCollectionID string
	Username         string
	Password         string
	ProjectID        string
	CollectionID     string
	Environment      string
}

func Load() (*Config, error) {
	cfg := &Config{
		BaseURL:          os.Getenv("BASE_URL"),
		TokenID:          os.Getenv("TOKEN_ID"),
		Token:            os.Getenv("TOKEN"),
		AuthCollectionID: os.Getenv("AUTH_COLLECTION_ID"),
		Username:         os.Getenv("USERNAME"),
		Password:         os.Getenv("PASSWORD"),
		ProjectID:        os.Getenv("PROJECT_ID"),
		CollectionID:     os.Getenv("COLLECTION_ID"),
		Environment:      os.Getenv("ENVIRONMENT"),
	}

	if cfg.BaseURL == "" || cfg.TokenID == "" || cfg.Token == "" || cfg.AuthCollectionID == "" ||
		cfg.Username == "" || cfg.Password == "" || cfg.ProjectID == "" || cfg.CollectionID == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
