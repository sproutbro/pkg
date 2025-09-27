package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Sqlite Sqlite
		Redis  Redis
		PG     PG

		Only Only
	}

	Only struct {
		Kakao  Kakao
		Google Google
		Redis  Redis
		PORT   string `env:"ONLY_PORT"`
	}

	// DB
	Sqlite struct {
		SQLITE1 string `env:"SQLITE1"`
		SQLITE2 string `env:"SQLITE2"`
		SQLITE3 string `env:"SQLITE3"`
	}

	Redis struct {
		ADDR     string `env:"REDIS_ADDR"`
		PASS     string `env:"REDIS_PASS"`
		DB       int    `env:"REDIS_DB"`
		PROTOCOL int    `env:"REDIS_PROTOCOL"`
	}

	PG struct {
		HOST string `env:"PG_HOST"`
		NAME string `env:"PG_NAME"`
		USER string `env:"PG_USER"`
		PASS string `env:"PG_PASS"`
		PORT string `env:"PG_PORT"`
	}

	// Auth
	Kakao struct {
		ClientID     string   `json:"clientID" env:"KAKAO_CLIENT"`
		ClientSecret string   `json:"clientSecret" env:"KAKAO_SECRET"`
		RedirectURL  string   `json:"redirectURL" env:"KAKAO_REDIRECT"`
		Scopes       []string `json:"scopes" env:"KAKAO_SCOPES"`
	}

	Google struct {
		Client   string   `json:"clientID" env:"GOOGLE_CLIENT"`
		Secret   string   `json:"clientSecret" env:"GOOGLE_SECRET"`
		Redirect string   `json:"redirectURL" env:"GOOGLE_REDIRECT"`
		Scopes   []string `json:"scopes" env:"GOOGLE_SCOPES"`
	}
)

func Load(envpath string) (*Config, error) {
	err := godotenv.Load(envpath)
	if err != nil {
		return nil, fmt.Errorf("config.go 61: %w", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config.go 66: %w", err)
	}

	return cfg, nil
}
