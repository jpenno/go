package main

import (
	"flappy-bird/internal/game"
	"flappy-bird/internal/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 720
	SCREEN_HEIGHT = 1080
)

func main() {
	app_state := state.START

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "flappy bird clone")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := game.Game{}
	game.Init()

	for app_state != state.QUIT {
		switch app_state {
		case state.START:
			if rl.IsKeyPressed(rl.KeySpace) {
				app_state = state.PLAY
			}
			if rl.IsKeyPressed(rl.KeyEscape) {
				app_state = state.QUIT
			}
		case state.PLAY:
			app_state = game.Update()
		case state.PAUSE:
			if rl.IsKeyPressed(rl.KeySpace) {
				app_state = state.PLAY
			}
			if rl.IsKeyPressed(rl.KeyEscape) {
				app_state = state.START
			}
			if rl.IsKeyPressed(rl.KeyQ) {
				app_state = state.QUIT
			}
		case state.GAME_OVER:
			if rl.IsKeyPressed(rl.KeyR) {
				app_state = state.PLAY
				game.Init()
			}
			if rl.IsKeyPressed(rl.KeyEscape) {
				app_state = state.QUIT
			}
		}

		rl.BeginDrawing()

		switch app_state {
		case state.START:
			rl.ClearBackground(rl.DarkGray)
			text := "Press space to start"
			text_len := rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2, 32, rl.Green)
		case state.PLAY:
			game.Draw()
		case state.PAUSE:
			rl.ClearBackground(rl.DarkGray)

			text := "Pause"
			text_len := rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2, 32, rl.Blue)

			text = "Press space to resume"
			text_len = rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2+32, 32, rl.Blue)

			text = "Press q to quit"
			text_len = rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2+32*2, 32, rl.Blue)
		case state.GAME_OVER:
			rl.ClearBackground(rl.DarkGray)

			text := "Game over"
			text_len := rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2, 32, rl.Red)

			text = "Press r to restart"
			text_len = rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2+32, 32, rl.Red)

			text = "Press escape to quit"
			text_len = rl.MeasureText(text, 32) / 2
			rl.DrawText(text, SCREEN_WIDTH/2-text_len, SCREEN_HEIGHT/2+32*2, 32, rl.Red)
		}

		rl.EndDrawing()
	}
}
