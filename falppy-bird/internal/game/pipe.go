package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pipe struct {
	rect     rl.Rectangle
	pos      rl.Vector2
	size     rl.Vector2
	velocity rl.Vector2
}

func pipeInit(rect rl.Rectangle) Pipe {
	return Pipe{
		rect:     rect,
		velocity: rl.Vector2{X: -200.0, Y: 0.0},
	}
}

func (p *Pipe) update(dt float32) {
	p.rect.X += p.velocity.X * dt
}

func (p *Pipe) draw() {
	rl.DrawRectangleRec(p.rect, rl.Green)
}
