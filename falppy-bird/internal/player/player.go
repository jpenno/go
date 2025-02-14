package player

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
	pos      rl.Vector2
	velocity rl.Vector2
	size     rl.Vector2
	garvity  float32
	State    PLAYER_STATE
}

func PlayerInit() Player {
	var garvity float32 = 300
	return Player{
		garvity:  garvity,
		pos:      rl.Vector2{X: 100, Y: 100},
		velocity: rl.Vector2{X: 100, Y: garvity},
		size:     rl.Vector2{X: 100, Y: 100},
		State:    ALIVE,
	}
}

func (p *Player) Update(dt float32) PLAYER_STATE {
	if p.velocity.Y < p.garvity {
		p.velocity.Y += 3_500 * dt
	}

	p.pos.Y += p.velocity.Y * dt

	if rl.IsKeyPressed(rl.KeySpace) {
		p.velocity.Y = -800
		p.pos.X += p.velocity.X * dt
	}

	if p.pos.Y < 0 {
		p.State = DEAD
	}
	if p.pos.Y > float32(rl.GetScreenHeight()) {
		p.State = DEAD
	}

	return p.State
}

func (p Player) Draw() {
	rl.DrawRectangleV(p.pos, p.size, rl.Blue)
}

func PrintPlayer() {
	fmt.Println("testing")
}
