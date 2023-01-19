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
	leftDark, _, _   = ebitenutil.NewImageFromFile("res/prompts/left_dark.png")
	leftLight, _, _  = ebitenutil.NewImageFromFile("res/prompts/left_light.png")
	rightDark, _, _  = ebitenutil.NewImageFromFile("res/prompts/right_dark.png")
	rightLight, _, _ = ebitenutil.NewImageFromFile("res/prompts/right_light.png")
)

func ChangeColour(key ebiten.Key) {
	if key == ebiten.KeyRight || key == ebiten.Key(ebiten.MouseButtonRight) {
		if Player.Colour == "lime" {
			Player.Colour = "cyan"
		} else if Player.Colour == "cyan" {
			Player.Colour = "yellow"
		} else if Player.Colour == "yellow" {
			Player.Colour = "lime"
		}
	}

	if key == ebiten.KeyLeft || key == ebiten.Key(ebiten.MouseButtonLeft) {
		if Player.Colour == "lime" {
			Player.Colour = "yellow"
		} else if Player.Colour == "cyan" {
			Player.Colour = "lime"
		} else if Player.Colour == "yellow" {
			Player.Colour = "cyan"
		}
	}
}

func ColourSelectorCtl() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		ChangeColour(ebiten.KeyRight)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		ChangeColour(ebiten.KeyLeft)
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
