package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type pipe struct {
	rect     rl.Rectangle
	velocity rl.Vector2
	Active   bool
}

func spawnPipePair(gape rl.Rectangle) []pipe {
	pipes := make([]pipe, 2)
	var posX float32 = 800

	rect := rl.Rectangle{
		X:      posX,
		Y:      0,
		Width:  50,
		Height: gape.Y,
	}
	pipes[0] = pipeInit(rect)

	rect = rl.Rectangle{
		X:      posX,
		Y:      gape.Y + gape.Height,
		Width:  50,
		Height: float32(rl.GetScreenHeight()),
	}
	pipes[1] = pipeInit(rect)

	return pipes
}

func pipeInit(rect rl.Rectangle) pipe {
	return pipe{
		rect:     rect,
		velocity: rl.Vector2{X: -200.0, Y: 0.0},
		Active:   true,
	}
}

func (p *pipe) update(dt float32) {
	p.rect.X += p.velocity.X * dt

	if p.rect.X+p.rect.Width < 0 {
		p.Active = false
	}
}

func (p *pipe) draw() {
	rl.DrawRectangleRec(p.rect, rl.Green)
}
