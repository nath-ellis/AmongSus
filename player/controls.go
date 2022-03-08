package player

import "github.com/hajimehoshi/ebiten/v2"

func Controls() {
	if c := Player.Obj.Check(-Player.XSpeed, 0); c == nil {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			Player.Obj.X -= Player.XSpeed
		}
	}

	if c := Player.Obj.Check(Player.XSpeed, 0); c == nil {
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			Player.Obj.X += Player.XSpeed
		}
	}

	if c := Player.Obj.Check(0, Player.YSpeed); c == nil { // gravity
		Player.Obj.Y += Player.YSpeed
		Player.YSpeed += Player.YVel
	} else {
		Player.YSpeed = 5
	}

	Player.Obj.Update()
}
