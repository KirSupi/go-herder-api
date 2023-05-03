package config

import (
	"encoding/json"
	"github.com/kirsupi/go-herder"
	"io"
	"os"
	"strconv"
)

type Config struct {
	Api    ApiConfig     `json:"api"`
	Herder herder.Config `json:"herder"`
}
type ApiConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func LoadFromJson(path string) (c Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonBytes, &c)
	if err != nil {
		return
	}
	return
}
func LoadFromEnv() (c Config, err error) {
	c.Api.Host = os.Getenv("API_HOST")
	c.Api.Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		return
	}
	if val := os.Getenv("HERDER_MAX_WORKERS_COUNT"); val != "" {
		c.Herder.MaxWorkersCount, err = strconv.Atoi(val)
		if err != nil {
			return
		}
	}
	if val := os.Getenv("HERDER_DEFAULT_MAX_STDOUT_LEN"); val != "" {
		c.Herder.DefaultMaxStdoutLen, err = strconv.Atoi(val)
		if err != nil {
			return
		}
	}
	if val := os.Getenv("HERDER_DEFAULT_MAX_STDERR_LEN"); val != "" {
		c.Herder.DefaultMaxStderrLen, err = strconv.Atoi(os.Getenv("HERDER_DEFAULT_MAX_STDERR_LEN"))
		if err != nil {
			return
		}
	}
	return
}
