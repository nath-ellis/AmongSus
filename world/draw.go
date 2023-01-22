package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bg, _, _ = ebitenutil.NewImageFromFile("res/bg.png")
)

func Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)

	for _, o := range Objects {
		o.Draw(screen)
	}

	// Update the coin animations
	for i, c := range Coins {
		if c.AnimationFrame >= (len(c.Sprites) - 1) {
			Coins[i].AnimationFrame = 0
		}

		if c.AnimationFPS < 0 {
			Coins[i].AnimationFrame += 1
			Coins[i].AnimationFPS += 5
		} else {
			Coins[i].AnimationFPS -= 1
		}
	}

	for _, c := range Coins {
		c.Draw(screen)
	}
}
