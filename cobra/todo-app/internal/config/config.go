package config

import (
	"cli-todoapp/internal/fileIO"
)

type Config struct {
	Path string `json:"path"`
}

func LoadConfig() Config {
	conf := fileio.GetJson("./assets/config.json", Config{})
	return conf
}
