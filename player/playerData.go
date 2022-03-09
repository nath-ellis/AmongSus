package player

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type PlayerData struct {
	Obj          *resolv.Object
	Colour       string
	XSpeed       float64
	YSpeed       float64
	YVel         float64
	JumpSpeed    float64
	Falling      bool
	WalkingStage int
	WalkCool     int
	State        string
}

var Player PlayerData

func Init() {
	Player.Obj = resolv.NewObject(50, 50, 90, 95, "player")
	Player.Colour = "lime"
	Player.XSpeed = 5
	Player.YSpeed = 2.0
	Player.YVel = 1.0
	Player.JumpSpeed = 21.5
	Player.Falling = false
	Player.WalkingStage = 0
	Player.WalkCool = 0
	Player.State = "game"

	space.Space.Add(Player.Obj)

	// Imports sprites
	limeWalking1, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-1.png")
	limeWalking2, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-2.png")
	limeWalking3, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-3.png")
	limeWalking4, _, _ := ebitenutil.NewImageFromFile("res/Lime/walking-4.png")

	limeWalking = append(limeWalking, limeWalking1)
	limeWalking = append(limeWalking, limeWalking2)
	limeWalking = append(limeWalking, limeWalking3)
	limeWalking = append(limeWalking, limeWalking4)

	cyanWalking1, _, _ := ebitenutil.NewImageFromFile("res/Cyan/walking-1.png")
	cyanWalking2, _, _ := ebitenutil.NewImageFromFile("res/Cyan/walking-2.png")
	cyanWalking3, _, _ := ebitenutil.NewImageFromFile("res/Cyan/walking-3.png")
	cyanWalking4, _, _ := ebitenutil.NewImageFromFile("res/Cyan/walking-4.png")

	cyanWalking = append(cyanWalking, cyanWalking1)
	cyanWalking = append(cyanWalking, cyanWalking2)
	cyanWalking = append(cyanWalking, cyanWalking3)
	cyanWalking = append(cyanWalking, cyanWalking4)

	yellowWalking1, _, _ := ebitenutil.NewImageFromFile("res/Yellow/walking-1.png")
	yellowWalking2, _, _ := ebitenutil.NewImageFromFile("res/Yellow/walking-2.png")
	yellowWalking3, _, _ := ebitenutil.NewImageFromFile("res/Yellow/walking-3.png")
	yellowWalking4, _, _ := ebitenutil.NewImageFromFile("res/Yellow/walking-4.png")

	yellowWalking = append(yellowWalking, yellowWalking1)
	yellowWalking = append(yellowWalking, yellowWalking2)
	yellowWalking = append(yellowWalking, yellowWalking3)
	yellowWalking = append(yellowWalking, yellowWalking4)
}
