package player

import "github.com/hajimehoshi/ebiten/v2"

func Controls() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Player.Obj.Y -= Player.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Player.Obj.Y += Player.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		Player.Obj.X -= Player.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		Player.Obj.X += Player.Speed
	}

	Player.Obj.Update()

	if ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyS) ||
		ebiten.IsKeyPressed(ebiten.KeyA) ||
		ebiten.IsKeyPressed(ebiten.KeyD) {
		Player.Moving = true
	} else {
		Player.Moving = false
	}
}
