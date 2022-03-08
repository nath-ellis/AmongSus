package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	limeWalking   []*ebiten.Image
	cyanWalking   []*ebiten.Image
	yellowWalking []*ebiten.Image
)

func Draw(screen *ebiten.Image) {
	if Player.WalkingStage >= (len(limeWalking)-1) ||
		Player.WalkingStage >= (len(cyanWalking)-1) ||
		Player.WalkingStage >= (len(yellowWalking)-1) {
		Player.WalkingStage = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	switch Player.Colour {
	case "lime":
		screen.DrawImage(limeWalking[Player.WalkingStage], op)
	case "cyan":
		screen.DrawImage(cyanWalking[Player.WalkingStage], op)
	case "yellow":
		screen.DrawImage(yellowWalking[Player.WalkingStage], op)
	}

	if Player.WalkCool < 0 {
		Player.WalkingStage += 1
		Player.WalkCool += 5
	} else {
		Player.WalkCool -= 1
	}
}
