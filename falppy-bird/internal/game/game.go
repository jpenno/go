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

const (
	NUMBER_OF_PIPES = 20
)

type Game struct {
	player Player
	pipes  [NUMBER_OF_PIPES]pipe
	state  GAME_STATE

	pipeSpawnTime  float32
	pipeSpawnTimer float32
}

func (g *Game) Init() error {
	g.player = playerInit()

	g.initPipes()

	// g.spawnPipes()

	g.state = PLAYING

	g.pipeSpawnTime = 2.0
	g.pipeSpawnTimer = 0.0

	return nil
}

func (g *Game) spawnPipes() {
	newPipes := spawnPipePair(rl.Rectangle{
		X:      0,
		Y:      float32(rand.IntN(700-100) + 100),
		Width:  0,
		Height: 400,
	})

	for j := range newPipes {
		for i := range g.pipes {
			if !g.pipes[i].Active {
				g.pipes[i] = newPipes[j]
				break
			}
		}
	}
}

func (g *Game) initPipes() {
	for i := 0; i < NUMBER_OF_PIPES; i++ {
		g.pipes[i] = pipe{
			rect:     rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
			velocity: rl.Vector2{X: 0, Y: 0},
			Active:   false,
		}
	}
}

func (g *Game) Update() GAME_STATE {
	dt := rl.GetFrameTime()

	if g.player.update(dt) == PLAYER_STATE(DEAD) {
		g.state = GAME_OVER
	}

	for i := range g.pipes {
		g.pipes[i].update(dt)
	}

	g.pipeSpawnTimer -= dt
	if g.pipeSpawnTimer <= 0 {
		g.spawnPipes()
		g.pipeSpawnTimer = g.pipeSpawnTime
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
