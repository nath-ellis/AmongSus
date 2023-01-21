package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/world"
)

func (p *PlayerData) Controls() {
	// X Collision and movement
	xSpeed := 0.0

	if ebiten.IsKeyPressed(ebiten.KeyA) && p.Obj.X > 0 {
		xSpeed = -p.XSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && p.Obj.X < 1105 {
		xSpeed = p.XSpeed
	}

	if c := p.Obj.Check(xSpeed, 0, "object"); c != nil {
		xSpeed = c.ContactWithObject(c.Objects[0]).X()

		if objs := c.ObjectsByTags("object"); objs[0].HasTags("column", "turretbase", "spikes") {
			if p.Obj.X < objs[0].X { // Left
				p.Obj.X -= world.Speed // Move player with object
			}
		}
	}

	p.Obj.X += xSpeed

	// Y Collision and Jumping
	p.YSpeed += p.YVel

	if !p.Falling {
		if ebiten.IsKeyPressed(ebiten.KeySpace) { // jumping
			p.YSpeed = -p.JumpSpeed
			p.Falling = true
		}
	}

	p.Falling = true

	ySpeed := p.YSpeed
	ySpeed = math.Max(math.Min(ySpeed, 16), -16)

	checkDistance := ySpeed

	if checkDistance >= 0 {
		checkDistance++
	}

	if c := p.Obj.Check(0, checkDistance, "object"); c != nil {
		if objs := c.ObjectsByTags("object"); len(objs) > 0 {
			ySpeed = c.ContactWithObject(objs[0]).Y()
			p.YSpeed = 0

			// Collision with spikes
			if objs[0].HasTags("spikes") {
				p.State = "gameOver"
			}

			if objs[0].Y > p.Obj.Y {
				p.Falling = false
			}
		}
	}

	p.Obj.Y += ySpeed

	p.Obj.Update()

	if p.Obj.X <= -p.Obj.W {
		p.State = "gameOver"
	}
}
