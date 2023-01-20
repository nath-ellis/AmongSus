package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/world"
)

func Controls() {
	// X Collision and movement
	xSpeed := 0.0

	if ebiten.IsKeyPressed(ebiten.KeyA) && Player.Obj.X > 0 {
		xSpeed = -Player.XSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && Player.Obj.X < 1105 {
		xSpeed = Player.XSpeed
	}

	if c := Player.Obj.Check(xSpeed, 0, "object"); c != nil {
		xSpeed = c.ContactWithObject(c.Objects[0]).X()

		if objs := c.ObjectsByTags("object"); objs[0].HasTags("column", "turretbase", "spikes") {
			if Player.Obj.X < objs[0].X { // Left
				Player.Obj.X -= world.Speed // Move player with object
			}
		}
	}

	Player.Obj.X += xSpeed

	// Y Collision and Jumping
	Player.YSpeed += Player.YVel

	if !Player.Falling {
		if ebiten.IsKeyPressed(ebiten.KeySpace) { // jumping
			Player.YSpeed = -Player.JumpSpeed
			Player.Falling = true
		}
	}

	Player.Falling = true

	ySpeed := Player.YSpeed
	ySpeed = math.Max(math.Min(ySpeed, 16), -16)

	checkDistance := ySpeed

	if checkDistance >= 0 {
		checkDistance++
	}

	if c := Player.Obj.Check(0, checkDistance, "object"); c != nil {
		if objs := c.ObjectsByTags("object"); len(objs) > 0 {
			ySpeed = c.ContactWithObject(objs[0]).Y()
			Player.YSpeed = 0

			// Collision with spikes
			if objs[0].HasTags("spikes") {
				Player.State = "gameOver"
			}

			if objs[0].Y > Player.Obj.Y {
				Player.Falling = false
			}
		}
	}

	Player.Obj.Y += ySpeed

	Player.Obj.Update()

	if Player.Obj.X <= -Player.Obj.W {
		Player.State = "gameOver"
	}
}
