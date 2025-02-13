package main

import (
	"flappy-bird/internal/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(1920, 1080, "flappy bird clone")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGray)

		player.PrintPlayer()

		rl.EndDrawing()
	}
}
