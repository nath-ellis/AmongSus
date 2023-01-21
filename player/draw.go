package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (p *PlayerData) Draw(screen *ebiten.Image) {
	if p.WalkingStage >= (len(p.WalkingSprites) - 1) {
		p.WalkingStage = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Obj.X, p.Obj.Y)

	if p.Falling {
		screen.DrawImage(p.FallingSprite, op)
	} else {
		screen.DrawImage(p.WalkingSprites[p.WalkingStage], op)
	}

	if p.WalkCool < 0 {
		p.WalkingStage += 1
		p.WalkCool += 5
	} else {
		p.WalkCool -= 1
	}
}
