package config

import (
	"errors"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-default:"./storage/storage.db"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"3s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func ReadConfig() (*Config, error) {
	configPath := os.Getenv("SERVER_CONFIG_PATH")
	if len(configPath) == 0 {
		return nil, errors.New("config path is empty")
	}

	if fileInfo, err := os.Stat(configPath); os.IsNotExist(err) || fileInfo.IsDir() {
		return nil, errors.New("config file is not exist")
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
