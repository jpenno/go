package config

import (
	fileio "cli-todoapp/internal/fileIO"
)

var (
	C Config
)

type Config struct {
	Path string `json:"path"`
	Todo string `json:"todo"`
}

func Init() {
	C = fileio.GetJson("./assets/config.json", Config{})
}
