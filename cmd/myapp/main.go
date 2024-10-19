package main

import (
	"fmt"
	"log/slog"

	"github.com/SensusX/myapp/internal/config"
	"github.com/SensusX/myapp/internal/log"
	"github.com/SensusX/myapp/internal/storage"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("ReadConfig Error: " + err.Error())
		return
	}

	log := log.SetupLogger(cfg.Env)
	//log = log.With(slog.String("env", cfg.Env))
	log.Info("myapp is running", slog.String("env", cfg.Env))
	//log.Debug("Debug...")

	storage, err := storage.New(cfg.StorageType, cfg.StorageDriver, *log)
	if err != nil {
		fmt.Println("storage.New Error: " + err.Error())
		return
	}
	storage.Open(*log)

}
