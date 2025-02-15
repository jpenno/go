package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player Player
	pipe   Pipe
	state  GAME_STATE
}
type GAME_STATE int

const (
	PLAYING GAME_STATE = iota
	PAUSE
	GAME_OVER
	QUIT
)

func (g *Game) Init() error {
	g.player = playerInit()
	g.pipe = pipeInit(rl.Rectangle{X: 600.0, Y: 0.0, Width: 100, Height: 500})
	g.state = PLAYING

	return nil
}

func (g *Game) Update() GAME_STATE {
	dt := rl.GetFrameTime()
	g.player.update(dt)
	g.pipe.update(dt)

	g.collision()

	if rl.IsKeyPressed(rl.KeyEscape) {
		return PAUSE
	}

	return g.state
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.DarkGray)
	g.player.draw()
	g.pipe.draw()
}

func (g *Game) collision() {
	if rl.CheckCollisionRecs(g.pipe.rect, g.player.rect) {
		g.state = GAME_OVER
	}
}
