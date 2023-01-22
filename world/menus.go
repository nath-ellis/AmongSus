package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	title, _, _    = ebitenutil.NewImageFromFile("res/text/title.png")
	gameOver, _, _ = ebitenutil.NewImageFromFile("res/text/gameover.png")
	click, _, _    = ebitenutil.NewImageFromFile("res/text/click-to-play.png")
	menubg, _, _   = ebitenutil.NewImageFromFile("res/space.png")
)

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
