package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type PlayerData struct {
	Obj            *resolv.Object
	Colour         string
	XSpeed         float64
	YSpeed         float64
	YVel           float64
	JumpSpeed      float64
	Falling        bool
	IdleSprite     *ebiten.Image
	FallingSprite  *ebiten.Image
	WalkingSprites []*ebiten.Image
	WalkingStage   int
	WalkCool       int
	State          string
}

var Player PlayerData

func Init() {
	Player.Obj = resolv.NewObject(50, -100, 90, 95, "player")
	Player.Colour = "lime"
	Player.XSpeed = 5
	Player.YSpeed = 2.0
	Player.YVel = 1.0
	Player.JumpSpeed = 21.0
	Player.Falling = false
	Player.WalkingStage = 0
	Player.WalkCool = 0
	Player.State = "menu"

	space.Space.Add(Player.Obj)

	loadSprites()
}

func loadSprites() {
	// Imports sprites
	switch Player.Colour {
	case "lime":
		Player.IdleSprite, _, _ = ebitenutil.NewImageFromFile("res/lime/idle.png")

		Player.FallingSprite, _, _ = ebitenutil.NewImageFromFile("res/lime/jumping.png")

		Player.WalkingSprites = []*ebiten.Image{}

		walking1, _, _ := ebitenutil.NewImageFromFile("res/lime/walking-1.png")
		walking2, _, _ := ebitenutil.NewImageFromFile("res/lime/walking-2.png")
		walking3, _, _ := ebitenutil.NewImageFromFile("res/lime/walking-3.png")
		walking4, _, _ := ebitenutil.NewImageFromFile("res/lime/walking-4.png")

		Player.WalkingSprites = append(Player.WalkingSprites, walking1)
		Player.WalkingSprites = append(Player.WalkingSprites, walking2)
		Player.WalkingSprites = append(Player.WalkingSprites, walking3)
		Player.WalkingSprites = append(Player.WalkingSprites, walking4)
	case "cyan":
		Player.IdleSprite, _, _ = ebitenutil.NewImageFromFile("res/cyan/idle.png")

		Player.FallingSprite, _, _ = ebitenutil.NewImageFromFile("res/cyan/jumping.png")

		Player.WalkingSprites = []*ebiten.Image{}

		walking1, _, _ := ebitenutil.NewImageFromFile("res/cyan/walking-1.png")
		walking2, _, _ := ebitenutil.NewImageFromFile("res/cyan/walking-2.png")
		walking3, _, _ := ebitenutil.NewImageFromFile("res/cyan/walking-3.png")
		walking4, _, _ := ebitenutil.NewImageFromFile("res/cyan/walking-4.png")

		Player.WalkingSprites = append(Player.WalkingSprites, walking1)
		Player.WalkingSprites = append(Player.WalkingSprites, walking2)
		Player.WalkingSprites = append(Player.WalkingSprites, walking3)
		Player.WalkingSprites = append(Player.WalkingSprites, walking4)
	case "yellow":
		Player.IdleSprite, _, _ = ebitenutil.NewImageFromFile("res/yellow/idle.png")

		Player.FallingSprite, _, _ = ebitenutil.NewImageFromFile("res/yellow/jumping.png")

		Player.WalkingSprites = []*ebiten.Image{}

		walking1, _, _ := ebitenutil.NewImageFromFile("res/yellow/walking-1.png")
		walking2, _, _ := ebitenutil.NewImageFromFile("res/yellow/walking-2.png")
		walking3, _, _ := ebitenutil.NewImageFromFile("res/yellow/walking-3.png")
		walking4, _, _ := ebitenutil.NewImageFromFile("res/yellow/walking-4.png")

		Player.WalkingSprites = append(Player.WalkingSprites, walking1)
		Player.WalkingSprites = append(Player.WalkingSprites, walking2)
		Player.WalkingSprites = append(Player.WalkingSprites, walking3)
		Player.WalkingSprites = append(Player.WalkingSprites, walking4)
	}
}
