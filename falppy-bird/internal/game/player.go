package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PLAYER_STATE int

const (
	ALIVE PLAYER_STATE = iota
	DEAD
)

type Player struct {
	rect     rl.Rectangle
	velocity rl.Vector2
	garvity  float32
	State    PLAYER_STATE
}

func playerInit() Player {
	var garvity float32 = 300
	return Player{
		garvity: garvity,
		rect: rl.Rectangle{
			X:      100,
			Y:      100,
			Width:  100,
			Height: 100,
		},
		velocity: rl.Vector2{X: 0, Y: garvity},
		State:    ALIVE,
	}
}

func (p *Player) update(dt float32) PLAYER_STATE {
	if p.velocity.Y < p.garvity {
		p.velocity.Y += 3_500 * dt
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		p.velocity.Y = -800
		p.rect.X += p.velocity.X * dt
	}

	p.rect.Y += p.velocity.Y * dt

	if p.rect.Y < 0 {
		p.State = DEAD
	}
	if p.rect.Y > float32(rl.GetScreenHeight()) {
		p.State = DEAD
	}

	return p.State
}

func (p *Player) draw() {
	rl.DrawRectangleRec(p.rect, rl.Blue)
}

func PrintPlayer() {
	fmt.Println("testing")
}
