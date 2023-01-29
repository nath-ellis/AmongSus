package ui

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	VCROSDMono font.Face

	title    *ebiten.Image
	gameOver *ebiten.Image
	click    *ebiten.Image
	menubg   *ebiten.Image
	coin     *ebiten.Image

	LeftColourBtn  Button
	RightColourBtn Button
)

func Init() {
	fontBytes, _ := os.ReadFile("res/text/vcrosdmono.ttf")
	parsed, _ := opentype.Parse(fontBytes)
	VCROSDMono, _ = opentype.NewFace(parsed, &opentype.FaceOptions{
		Size:    76,
		DPI:     64,
		Hinting: font.HintingFull,
	})

	title, _, _ = ebitenutil.NewImageFromFile("res/text/title.png")
	gameOver, _, _ = ebitenutil.NewImageFromFile("res/text/gameover.png")
	click, _, _ = ebitenutil.NewImageFromFile("res/text/click-to-play.png")
	menubg, _, _ = ebitenutil.NewImageFromFile("res/space.png")
	coin, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_1.png")

	leftDark, _, _ := ebitenutil.NewImageFromFile("res/prompts/left_dark.png")
	leftLight, _, _ := ebitenutil.NewImageFromFile("res/prompts/left_light.png")
	rightDark, _, _ := ebitenutil.NewImageFromFile("res/prompts/right_dark.png")
	rightLight, _, _ := ebitenutil.NewImageFromFile("res/prompts/right_light.png")

	LeftColourBtn = Button{
		Obj:           resolv.NewObject(5, 100, 32, 32),
		Sprite:        leftDark,
		OnHoverSprite: leftLight,
	}

	RightColourBtn = Button{
		Obj:           resolv.NewObject(55, 100, 32, 32),
		Sprite:        rightDark,
		OnHoverSprite: rightLight,
	}
}
