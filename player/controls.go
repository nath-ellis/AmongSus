package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/world"
)

func Controls() {
	if c := Player.Obj.Check(Player.XSpeed, Player.YSpeed, "object"); c != nil {
		checkObjects(c.Objects[0].X, c.Objects[0].Y, c.Objects[0].W)
	}

	xSpeed := Player.XSpeed

	if ebiten.IsKeyPressed(ebiten.KeyA) && Player.Obj.X > 0 {
		if c := Player.Obj.Check(-xSpeed, 0, "object"); c != nil {
			xSpeed = c.ContactWithCell(c.Cells[0]).X()
		}
		Player.Obj.X -= xSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) && Player.Obj.X < 1105 {
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

	if Player.Obj.X < -Player.Obj.W {
		Player.State = "gameOver"
	}
}

func checkObjects(objectX float64, objectY float64, objectW float64) {
	for _, o := range world.Objects {
		if o.Obj.X < 0 || o.Obj.X > 1200 { // Skip if off screen
			continue
		}

		if o.Type == "platform" {
			continue
		}

		if o.Type == "column" || o.Type == "turretbase" {
			if Player.Obj.X < objectX || Player.Obj.Y > objectY { // Left and above
				Player.Obj.X -= world.Speed // Move player with object
				Player.Obj.Update()
			}
			continue
		}

		if o.Type == "spikes" {
			if Player.Obj.X < objectX { // Left
				Player.Obj.X -= world.Speed // Move player with object
				Player.Obj.Update()
				continue
			} else if Player.Obj.X > (objectX + objectW) { // Right
				continue
			}
		}

		if o.Obj.X == objectX && o.Obj.Y == objectY {
			Player.State = "gameOver"
		}
	}
}
