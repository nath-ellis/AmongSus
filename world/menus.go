package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	title, _, _    = ebitenutil.NewImageFromFile("res/text/title.png")
	gameOver, _, _ = ebitenutil.NewImageFromFile("res/text/gameover.png")
	click, _, _    = ebitenutil.NewImageFromFile("res/text/click-to-play.png")
	menubg, _, _   = ebitenutil.NewImageFromFile("res/space.png")
)

func DrawMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(menubg, op)
	op.GeoM.Translate(375, 0)
	screen.DrawImage(menubg, op)
	op.GeoM.Translate(-375, 300)
	screen.DrawImage(menubg, op)
	op.GeoM.Translate(375, 0)
	screen.DrawImage(menubg, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(340, 200)

	screen.DrawImage(title, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(490, 400)
	screen.DrawImage(click, op)
}

func DrawGameOver(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(menubg, op)
	op.GeoM.Translate(375, 0)
	screen.DrawImage(menubg, op)
	op.GeoM.Translate(-375, 300)
	screen.DrawImage(menubg, op)
	op.GeoM.Translate(375, 0)
	screen.DrawImage(menubg, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(350, 200)

	screen.DrawImage(gameOver, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(490, 400)
	screen.DrawImage(click, op)
}
