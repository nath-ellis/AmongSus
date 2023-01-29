package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
)

type Button struct {
	Obj           *resolv.Object
	Sprite        *ebiten.Image
	OnHoverSprite *ebiten.Image
}

func (b Button) IsPressed() bool {
	mouseX, mouseY := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		float64(mouseX) > b.Obj.X && float64(mouseX) < b.Obj.X+b.Obj.W &&
		float64(mouseY) > b.Obj.Y && float64(mouseY) < b.Obj.Y+b.Obj.H {
		return true
	}

	return false
}

func (b Button) Draw(screen *ebiten.Image, scale bool) {
	op := &ebiten.DrawImageOptions{}
	if scale {
		op.GeoM.Scale(2, 2)
	}
	op.GeoM.Translate(b.Obj.X, b.Obj.Y)

	mouseX, mouseY := ebiten.CursorPosition()

	if float64(mouseX) > b.Obj.X && float64(mouseX) < b.Obj.X+b.Obj.W &&
		float64(mouseY) > b.Obj.Y && float64(mouseY) < b.Obj.Y+b.Obj.H {
		screen.DrawImage(b.OnHoverSprite, op)
	} else {
		screen.DrawImage(b.Sprite, op)
	}
}
