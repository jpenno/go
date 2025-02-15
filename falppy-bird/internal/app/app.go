package app

import (
	"flappy-bird/internal/game"
	"flappy-bird/internal/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	state        state.APP_STATE
	game         game.Game
	screenWidth  int32
	screenHeight int32
}

func (a *App) Init() {
	a.screenWidth = 720
	a.screenHeight = 1080
	a.state = state.START

	rl.InitWindow(a.screenWidth, a.screenHeight, "flappy bird clone")

	rl.SetTargetFPS(60)

	a.game = game.Game{}
	a.game.Init()
}

func (a *App) Close() {
	rl.CloseWindow()
}

func (a *App) Run() {
	for a.state != state.QUIT {
		a.update()
		a.draw()
	}
}

func (a *App) update() {
	switch a.state {
	case state.START:
		a.updateStart()
	case state.PLAY:
		a.state = a.game.Update()
	case state.PAUSE:
		a.updatePaused()
	case state.GAME_OVER:
		a.updateGameover()
	}
}

func (a *App) draw() {
	rl.BeginDrawing()

	switch a.state {
	case state.START:
		a.drawStart()
	case state.PLAY:
		a.game.Draw()
	case state.PAUSE:
		a.drawPause()
	case state.GAME_OVER:
		a.drawGameover()
	}

	rl.EndDrawing()
}

func (a *App) updateGameover() {
	if rl.IsKeyPressed(rl.KeyR) {
		a.state = state.PLAY
		a.game.Init()
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		a.state = state.QUIT
	}
}

func (a *App) updatePaused() {
	if rl.IsKeyPressed(rl.KeySpace) {
		a.state = state.PLAY
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		a.state = state.START
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		a.state = state.QUIT
	}
}

func (a *App) updateStart() {
	if rl.IsKeyPressed(rl.KeySpace) {
		a.state = state.PLAY
		a.game.Init()
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		a.state = state.QUIT
	}
}

func (a *App) drawGameover() {
	rl.ClearBackground(rl.DarkGray)

	text := []string{
		"Game over",
		"Press r to restart",
		"Press escape to quit",
	}

	a.drawCenteredList(text, rl.Red)
}

func (a *App) drawPause() {
	rl.ClearBackground(rl.DarkGray)

	text := []string{
		"Pause",
		"Press space to resume",
		"Press escape to go back to start",
		"Press q to quit",
	}

	a.drawCenteredList(text, rl.Blue)
}

func (a *App) drawStart() {
	rl.ClearBackground(rl.DarkGray)

	text := []string{
		"Flappy bird game",
		"Press space to start",
		"Press escape to quit",
	}

	a.drawCenteredList(text, rl.Green)
}

func (a *App) drawCenteredList(list []string, color rl.Color) {
	var item_len int32 = int32(len(list))

	for i, item := range list {
		text_len := rl.MeasureText(item, 32) / 2

		x := a.screenWidth/2 - text_len
		y := (a.screenHeight / 2) + int32(32*i)
		y -= 32 * item_len / 2

		rl.DrawText(item, x, y, 32, color)
	}
}
