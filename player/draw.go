package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	limeWalking         []*ebiten.Image
	cyanWalking         []*ebiten.Image
	yellowWalking       []*ebiten.Image
	limeFalling, _, _   = ebitenutil.NewImageFromFile("res/Lime/jumping.png")
	cyanFalling, _, _   = ebitenutil.NewImageFromFile("res/Cyan/jumping.png")
	yellowFalling, _, _ = ebitenutil.NewImageFromFile("res/Yellow/jumping.png")
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
		if Player.Falling {
			screen.DrawImage(limeFalling, op)
		} else {
			screen.DrawImage(limeWalking[Player.WalkingStage], op)
		}
	case "cyan":
		if Player.Falling {
			screen.DrawImage(cyanFalling, op)
		} else {
			screen.DrawImage(cyanWalking[Player.WalkingStage], op)
		}
	case "yellow":
		if Player.Falling {
			screen.DrawImage(yellowFalling, op)
		} else {
			screen.DrawImage(yellowWalking[Player.WalkingStage], op)
		}
	}

	if Player.WalkCool < 0 {
		Player.WalkingStage += 1
		Player.WalkCool += 5
	} else {
		Player.WalkCool -= 1
	}
}
