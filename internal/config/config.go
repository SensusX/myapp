package config

import (
	"errors"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Storage    `yaml:"storage"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

type Storage struct {
	StorageName   string `yaml:"storage_name" env-default:"PostgreSQL"`
	StorageType   string `yaml:"storage_type" env-default:"SQL"`
	StorageDriver string `yaml:"storage_driver" env-default:"pgx"`
	StoragePath   string `yaml:"storage_path" env-default:""`
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
