package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/nath-ellis/AmongSus/world"
)

type Game struct{}

func init() {
	space.Init(1200, 600)
	player.Init()

	world.Init()
}

func (g *Game) Update() error {
	player.Controls()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	player.Draw(screen)
	world.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1200, 600
}

func main() {
	ebiten.SetWindowSize(1200, 600)
	ebiten.SetWindowTitle("Among Sus")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal("Failed to run game: ", err)
	}
}
