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
	} else {
		c, err = config.LoadFromJson()
	}
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