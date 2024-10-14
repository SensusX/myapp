package main

import (
	"fmt"
	"log/slog"

	"github.com/SensusX/myapp/internal/config"
	"github.com/SensusX/myapp/internal/log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	log := log.SetupLogger(cfg.Env)
	//log = log.With(slog.String("env", cfg.Env))
	log.Info("myapp is running", slog.String("env", cfg.Env))
	//log.Debug("Debug...")

}
