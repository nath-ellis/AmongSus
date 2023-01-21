package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	leftDark, _, _   = ebitenutil.NewImageFromFile("res/prompts/left_dark.png")
	leftLight, _, _  = ebitenutil.NewImageFromFile("res/prompts/left_light.png")
	rightDark, _, _  = ebitenutil.NewImageFromFile("res/prompts/right_dark.png")
	rightLight, _, _ = ebitenutil.NewImageFromFile("res/prompts/right_light.png")
)

func (p *PlayerData) ChangeColour(key ebiten.Key) {
	if key == ebiten.KeyRight || key == ebiten.Key(ebiten.MouseButtonRight) {
		if p.Colour == "lime" {
			p.Colour = "cyan"
		} else if p.Colour == "cyan" {
			p.Colour = "yellow"
		} else if p.Colour == "yellow" {
			p.Colour = "lime"
		}
	}

	if key == ebiten.KeyLeft || key == ebiten.Key(ebiten.MouseButtonLeft) {
		if p.Colour == "lime" {
			p.Colour = "yellow"
		} else if p.Colour == "cyan" {
			p.Colour = "lime"
		} else if p.Colour == "yellow" {
			p.Colour = "cyan"
		}
	}

	p.LoadSprites()
}

func (p *PlayerData) ColourSelectorCtl() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.ChangeColour(ebiten.KeyRight)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.ChangeColour(ebiten.KeyLeft)
	}
}

func (p *PlayerData) DrawColourSelector(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 5)

	screen.DrawImage(p.IdleSprite, op)

	posX, posY := ebiten.CursorPosition()

	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(0, 100)

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || (posX > 5 && posX < 37) && (posY > 105 && posY < 137) {
		screen.DrawImage(leftLight, op)
	} else {
		screen.DrawImage(leftDark, op)
	}

	op.GeoM.Translate(40, 0)

	if ebiten.IsKeyPressed(ebiten.KeyRight) || (posX > 45 && posX < 77) && (posY > 105 && posY < 137) {
		screen.DrawImage(rightLight, op)
	} else {
		screen.DrawImage(rightDark, op)
	}
}
