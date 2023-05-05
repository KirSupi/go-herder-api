package main

import (
	"errors"
	"go-herder-api/internal/api"
	"go-herder-api/internal/config"
	"log"
	"os"
	"runtime"
)

func main() {
	var c config.Config
	var err = errors.New("config was not been loaded")
	if os.Getenv("LOAD_CONFIG_FROM_ENV") != "" {
		c, err = config.LoadFromEnv()
	} else if configPath := os.Getenv("LOAD_CONFIG_FROM_JSON"); configPath != "" {
		c, err = config.LoadFromJson(configPath)
	}
	c.Api.Host = "0.0.0.0"
	c.Api.Port = 80
	if err != nil {
		log.Fatalln(err.Error())
	}
	if c.Herder.MaxWorkersCount == 0 {
		c.Herder.MaxWorkersCount = runtime.NumCPU()
	}
	a := api.New(c)
	if err := a.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
