package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	limeWalking      []*ebiten.Image
	cyanWalking      []*ebiten.Image
	yellowWalking    []*ebiten.Image
	limeIdle, _, _   = ebitenutil.NewImageFromFile("res/Lime/idle.png")
	cyanIdle, _, _   = ebitenutil.NewImageFromFile("res/Cyan/idle.png")
	yellowIdle, _, _ = ebitenutil.NewImageFromFile("res/Yellow/idle.png")
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
		if Player.Moving {
			screen.DrawImage(limeWalking[Player.WalkingStage], op)
		} else {
			screen.DrawImage(limeIdle, op)
		}
	case "cyan":
		if Player.Moving {
			screen.DrawImage(cyanWalking[Player.WalkingStage], op)
		} else {
			screen.DrawImage(cyanIdle, op)
		}
	case "yellow":
		if Player.Moving {
			screen.DrawImage(yellowWalking[Player.WalkingStage], op)
		} else {
			screen.DrawImage(yellowIdle, op)
		}
	}

	if Player.WalkCool < 0 {
		Player.WalkingStage += 1
		Player.WalkCool += 5
	} else {
		Player.WalkCool -= 1
	}
}
