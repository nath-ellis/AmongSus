package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	platform, _, _ = ebitenutil.NewImageFromFile("res/terrain/platform.png")
)

func Draw(screen *ebiten.Image) {
	for _, o := range Objects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(o.Obj.X, o.Obj.Y)

		switch o.Type {
		case "platform":
			screen.DrawImage(platform, op)
		}
	}
}