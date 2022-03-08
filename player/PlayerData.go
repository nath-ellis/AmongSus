package player

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type PlayerData struct {
	Obj          *resolv.Object
	WalkingStage int
}

var Player PlayerData

func Init() {
	Player.Obj = resolv.NewObject(50, 50, 102, 102, "player")
	Player.WalkingStage = 0

	space.Space.Add(Player.Obj)

	// Imports sprites
	limeWalking1, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-1.png")
	limeWalking2, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-2.png")
	limeWalking3, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-3.png")
	limeWalking4, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-4.png")

	LimeWalking = append(LimeWalking, limeWalking1)
	LimeWalking = append(LimeWalking, limeWalking2)
	LimeWalking = append(LimeWalking, limeWalking3)
	LimeWalking = append(LimeWalking, limeWalking4)
}
