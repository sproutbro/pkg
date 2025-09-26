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

		Kakao  Kakao
		Google Google
	}

	// DB
	Sqlite struct {
		SQLITE1 string `env:"SQLITE1"`
		SQLITE2 string `env:"SQLITE2"`
		SQLITE3 string `env:"SQLITE3"`
		SQLITE4 string `env:"SQLITE4"`
		SQLITE5 string `env:"SQLITE5"`
	}

	Redis struct {
		ADDR string `env:"REDIS_ADDR"`
		PASS string `env:"REDIS_PASS"`
	}

	PG struct {
		DNS string `env:"PG_DNS"`
	}

	// Auth
	Kakao struct {
		Client   string `env:"KAKAO_CLIENT"`
		Secret   string `env:"KAKAO_SECRET"`
		Redirect string `env:"KAKAO_REDIRECT"`
		Endpoint struct {
			AUTH  string `env:"KAKAO_AUTH"`
			TOKEN string `env:"KAKAO_TOKEN"`
			INFO  string `env:"KAKAO_INFO"`
		}
	}

	Google struct {
		Client   string `env:"GOOGLE_CLIENT"`
		Secret   string `env:"GOOGLE_SECRET"`
		Redirect string `env:"GOOGLE_REDIRECT"`
		Endpoint struct {
			AUTH  string `env:"GOOGLE_AUTH"`
			TOKEN string `env:"GOOGLE_TOKEN"`
			INFO  string `env:"GOOGLE_INFO"`
		}
	}
)

func Load(envpath string) (*Config, error) {
	err := godotenv.Load(envpath)
	if err != nil {
		return nil, fmt.Errorf("config.go 65: %w", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config.go 71: %w", err)
	}

	return cfg, nil
}
