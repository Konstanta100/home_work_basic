package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() (*Cfg, error) {
	if err := load(); err != nil {
		return nil, fmt.Errorf("can't load config: %w", err)
	}

	cfg := &Cfg{}
	opts := env.Options{
		Prefix:                "APP_",
		UseFieldNameByDefault: true,
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, fmt.Errorf("can't parse config: %w", err)
	}

	return cfg, nil
}

func load() error {
	cfgEnv := os.Getenv("ENV_FILE")
	if cfgEnv == "" {
		cfgEnv = ".env"
	}

	err := godotenv.Load(cfgEnv)
	if err != nil {
		log.Printf("Предупреждение: не удалось загрузить конфиг из %s: %v", cfgEnv, err)
		return err
	}

	return nil
}
