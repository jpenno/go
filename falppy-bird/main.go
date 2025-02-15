package main

import (
	"flappy-bird/internal/app"
)

func main() {
	app := app.App{}
	app.Init()
	app.Run()
	app.Close()
}
