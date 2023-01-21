package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bg, _, _         = ebitenutil.NewImageFromFile("res/bg.png")
	platform, _, _   = ebitenutil.NewImageFromFile("res/terrain/platform.png")
	column, _, _     = ebitenutil.NewImageFromFile("res/terrain/column-1.png")
	spikes, _, _     = ebitenutil.NewImageFromFile("res/terrain/spikes.png")
	turretbase, _, _ = ebitenutil.NewImageFromFile("res/terrain/turretbase.png")
)

func Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)

	for _, o := range Objects {
		o.Draw(screen)
	}
}
