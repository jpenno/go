package main

import (
	"flappy-bird/internal/game"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(720, 1080, "flappy bird clone")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := game.Game{}
	err := game.Init()
	if err != nil {
		fmt.Printf("error: %q", err)
	}

	for !rl.WindowShouldClose() {
		game.Update()

		rl.BeginDrawing()
		game.Draw()
		rl.EndDrawing()
	}
}
