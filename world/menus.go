package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/solarlune/resolv"
)

var (
	title, _, _    = ebitenutil.NewImageFromFile("res/text/title.png")
	gameOver, _, _ = ebitenutil.NewImageFromFile("res/text/gameover.png")
	click, _, _    = ebitenutil.NewImageFromFile("res/text/click-to-play.png")
	menubg, _, _   = ebitenutil.NewImageFromFile("res/space.png")
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

func DrawMenu(screen *ebiten.Image, coins string) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.8, 0.8)
	screen.DrawImage(menubg, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(340, 200)

	screen.DrawImage(title, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(490, 400)
	screen.DrawImage(click, op)

	drawCoinAmount(screen, coins)
}

func DrawGameOver(screen *ebiten.Image, coins string) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.8, 0.8)
	screen.DrawImage(menubg, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(350, 200)

	screen.DrawImage(gameOver, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(490, 400)
	screen.DrawImage(click, op)

	drawCoinAmount(screen, coins)
}

func drawCoinAmount(screen *ebiten.Image, coins string) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(5, 540)
	screen.DrawImage(coinFrame1, op)
	text.Draw(screen, coins, VCROSDMono, 60, 590, color.White)
}
