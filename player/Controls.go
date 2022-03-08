package player

import "github.com/hajimehoshi/ebiten/v2"

func Controls() {
	if c := Player.Obj.Check(0, -Player.Speed); c == nil {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			Player.Obj.Y -= Player.Speed
		}
	}

	if c := Player.Obj.Check(0, Player.Speed); c == nil {
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			Player.Obj.Y += Player.Speed
		}
	}

	if c := Player.Obj.Check(-Player.Speed, 0); c == nil {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			Player.Obj.X -= Player.Speed
		}
	}

	if c := Player.Obj.Check(Player.Speed, 0); c == nil {
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			Player.Obj.X += Player.Speed
		}
	}

	Player.Obj.Update()
}
