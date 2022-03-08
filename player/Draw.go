package player

import "github.com/hajimehoshi/ebiten/v2"

var (
	LimeWalking []*ebiten.Image
)

func Draw(screen *ebiten.Image) {
	if Player.WalkingStage >= (len(LimeWalking) - 1) {
		Player.WalkingStage = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	screen.DrawImage(LimeWalking[Player.WalkingStage], op)
}
