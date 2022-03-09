package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func Controls() {
	xSpeed := Player.XSpeed

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if c := Player.Obj.Check(-xSpeed, 0, "object"); c != nil {
			xSpeed = c.ContactWithCell(c.Cells[0]).X()
		}
		Player.Obj.X -= xSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		if c := Player.Obj.Check(xSpeed, 0, "object"); c != nil {
			xSpeed = c.ContactWithCell(c.Cells[0]).X()
		}
		Player.Obj.X += xSpeed
	}

	Player.YSpeed += Player.YVel

	if !Player.Falling {
		if ebiten.IsKeyPressed(ebiten.KeySpace) { // jumping
			Player.YSpeed = -Player.JumpSpeed
			Player.Falling = true
		}
	}

	ySpeed := Player.YSpeed

	Player.Falling = true

	if c := Player.Obj.Check(0, ySpeed, "object"); c != nil {
		if objs := c.ObjectsByTags("object"); len(objs) > 0 {
			ySpeed = c.ContactWithObject(objs[0]).Y()
			Player.YSpeed = 0

			if objs[0].Y > Player.Obj.Y {
				Player.Falling = false
			}
		}
	}

	Player.Obj.Y += ySpeed

	Player.Obj.Update()
}
