package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var ()

func Draw(screen *ebiten.Image) {
	if Player.WalkingStage >= (len(Player.WalkingSprites) - 1) {
		Player.WalkingStage = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	if Player.Falling {
		screen.DrawImage(Player.FallingSprite, op)
	} else {
		screen.DrawImage(Player.WalkingSprites[Player.WalkingStage], op)
	}

	if Player.WalkCool < 0 {
		Player.WalkingStage += 1
		Player.WalkCool += 5
	} else {
		Player.WalkCool -= 1
	}
}
