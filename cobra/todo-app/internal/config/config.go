package config

import (
	"cli-todoapp/internal/fileIO"
)

var (
	CONFIG Config
)

type Config struct {
	Path string `json:"path"`
}

func Init() {
	CONFIG = fileio.GetJson("./assets/config.json", Config{})
}
