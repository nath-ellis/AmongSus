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
	Coins          int
}

func (p *PlayerData) Init() {
	p.Obj = resolv.NewObject(50, -100, 90, 95, "player")
	p.Colour = "lime"
	p.XSpeed = 5
	p.YSpeed = 2.0
	p.YVel = 1.0
	p.JumpSpeed = 21.0
	p.Falling = false
	p.WalkingStage = 0
	p.WalkCool = 0
	p.State = "menu"
	p.Coins = 0

	space.Space.Add(p.Obj)

	p.LoadSprites()
}

func (p *PlayerData) LoadSprites() {
	// Import sprites
	p.IdleSprite, _, _ = ebitenutil.NewImageFromFile("res/" + p.Colour + "/idle.png")

	p.FallingSprite, _, _ = ebitenutil.NewImageFromFile("res/" + p.Colour + "/jumping.png")

	p.WalkingSprites = []*ebiten.Image{}

	walking1, _, _ := ebitenutil.NewImageFromFile("res/" + p.Colour + "/walking-1.png")
	walking2, _, _ := ebitenutil.NewImageFromFile("res/" + p.Colour + "/walking-2.png")
	walking3, _, _ := ebitenutil.NewImageFromFile("res/" + p.Colour + "/walking-3.png")
	walking4, _, _ := ebitenutil.NewImageFromFile("res/" + p.Colour + "/walking-4.png")

	p.WalkingSprites = append(p.WalkingSprites, walking1)
	p.WalkingSprites = append(p.WalkingSprites, walking2)
	p.WalkingSprites = append(p.WalkingSprites, walking3)
	p.WalkingSprites = append(p.WalkingSprites, walking4)
}

var Player PlayerData
