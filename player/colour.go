package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	limeIdle, _, _   = ebitenutil.NewImageFromFile("res/Lime/idle.png")
	cyanIdle, _, _   = ebitenutil.NewImageFromFile("res/Cyan/idle.png")
	yellowIdle, _, _ = ebitenutil.NewImageFromFile("res/Yellow/idle.png")
)

func ColourSelectorCtl() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if Player.Colour == "lime" {
			Player.Colour = "cyan"
		} else if Player.Colour == "cyan" {
			Player.Colour = "yellow"
		} else if Player.Colour == "yellow" {
			Player.Colour = "lime"
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if Player.Colour == "lime" {
			Player.Colour = "yellow"
		} else if Player.Colour == "cyan" {
			Player.Colour = "lime"
		} else if Player.Colour == "yellow" {
			Player.Colour = "cyan"
		}
	}
}

func DrawColourSelector(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 5)

	if Player.Colour == "lime" {
		screen.DrawImage(limeIdle, op)
	} else if Player.Colour == "cyan" {
		screen.DrawImage(cyanIdle, op)
	} else if Player.Colour == "yellow" {
		screen.DrawImage(yellowIdle, op)
	}
}
