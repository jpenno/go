package game

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GAME_STATE int

const (
	PLAYING GAME_STATE = iota
	PAUSE
	GAME_OVER
	QUIT
)

type Game struct {
	player Player
	pipes  []pipe
	state  GAME_STATE
}

func (g *Game) Init() error {
	g.player = playerInit()

	g.pipes = spawnPipePair(rl.Rectangle{
		X:      0,
		Y:      float32(rand.IntN(800-100) + 100),
		Width:  0,
		Height: 400,
	})

	g.state = PLAYING

	return nil
}

func (g *Game) Update() GAME_STATE {
	dt := rl.GetFrameTime()

	if g.player.update(dt) == PLAYER_STATE(DEAD) {
		g.state = GAME_OVER
	}

	for i := range g.pipes {
		g.pipes[i].update(dt)
	}

	g.collision()

	if rl.IsKeyPressed(rl.KeyEscape) {
		return PAUSE
	}

	return g.state
}

func (g *Game) Draw() {
	rl.ClearBackground(rl.DarkGray)
	g.player.draw()

	for i := range g.pipes {
		g.pipes[i].draw()
	}
}

func (g *Game) collision() {
	for i := range g.pipes {
		if rl.CheckCollisionRecs(
			g.pipes[i].rect,
			g.player.rect,
		) {
			g.state = GAME_OVER
		}
	}
}
