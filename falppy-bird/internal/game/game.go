package game

import (
	"flappy-bird/internal/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player player.Player
}

func (g *Game) Init() error {
	g.player = player.PlayerInit()

	return nil
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()
	g.player.Update(dt)
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.DarkGray)
	g.player.Draw()
}
