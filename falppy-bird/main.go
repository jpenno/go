package main

import (
	"flappy-bird/internal/app"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

// const (
// 	SCREEN_WIDTH  = 720
// 	SCREEN_HEIGHT = 1080
// )

func main() {
	app := app.App{}
	app.Init()
	app.Run()
	app.Close()
}
