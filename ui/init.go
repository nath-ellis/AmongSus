package ui

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	VCROSDMono font.Face

	title     *ebiten.Image
	gameOver  *ebiten.Image
	click     *ebiten.Image
	menubg    *ebiten.Image
	coin      *ebiten.Image
	shopLight *ebiten.Image
	shopDark  *ebiten.Image
	xLight    *ebiten.Image
	xDark     *ebiten.Image
	homeLight *ebiten.Image
	homeDark  *ebiten.Image
	itemBG    *ebiten.Image
	whiteIdle *ebiten.Image
	blackIdle *ebiten.Image
	dark      *ebiten.Image

	LeftColourBtn  Button
	RightColourBtn Button
	shopBtn        Button
	homeBtn        Button

	shopOpen bool = false
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
	itemBG, _, _ = ebitenutil.NewImageFromFile("res/prompts/item_bg.png")

	leftDark, _, _ := ebitenutil.NewImageFromFile("res/prompts/left_dark.png")
	leftLight, _, _ := ebitenutil.NewImageFromFile("res/prompts/left_light.png")
	rightDark, _, _ := ebitenutil.NewImageFromFile("res/prompts/right_dark.png")
	rightLight, _, _ := ebitenutil.NewImageFromFile("res/prompts/right_light.png")

	whiteIdle, _, _ = ebitenutil.NewImageFromFile("res/white/idle.png")
	blackIdle, _, _ = ebitenutil.NewImageFromFile("res/black/idle.png")

	shopLight, _, _ = ebitenutil.NewImageFromFile("res/prompts/shop_light.png")
	shopDark, _, _ = ebitenutil.NewImageFromFile("res/prompts/shop_dark.png")
	xLight, _, _ = ebitenutil.NewImageFromFile("res/prompts/x_light.png")
	xDark, _, _ = ebitenutil.NewImageFromFile("res/prompts/x_dark.png")
	homeLight, _, _ = ebitenutil.NewImageFromFile("res/prompts/home_light.png")
	homeDark, _, _ = ebitenutil.NewImageFromFile("res/prompts/home_dark.png")

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
	shopBtn = Button{
		Obj:           resolv.NewObject(1124, 10, 66, 60),
		Sprite:        shopLight,
		OnHoverSprite: shopDark,
	}
	homeBtn = Button{
		Obj:           resolv.NewObject(1124, 10, 64, 64),
		Sprite:        homeLight,
		OnHoverSprite: homeDark,
	}

	dark = ebiten.NewImage(150, 150)
	dark.Fill(color.RGBA{0, 0, 0, 100})
}
