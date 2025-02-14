package game

import (
	"flappy-bird/internal/player"
	"flappy-bird/internal/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player player.Player
}

func (g *Game) Init() error {
	g.player = player.PlayerInit()

	return nil
}

func (g *Game) Update() state.APP_STATE {
	dt := rl.GetFrameTime()
	g.player.Update(dt)

	if rl.IsKeyPressed(rl.KeyEscape) {
		return state.PAUSE
	}

	if g.player.State == player.DEAD {
		return state.GAME_OVER
	}

	return state.PLAY
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.DarkGray)
	g.player.Draw()
}
