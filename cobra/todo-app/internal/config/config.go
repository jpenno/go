package config

import (
	"cli-todoapp/internal/fileIO"
)

var (
	C Config
)

type Config struct {
	Path string `json:"path"`
}

func Init() {
	C = fileio.GetJson("./assets/config.json", Config{})
}
