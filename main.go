package main

import (
	"log"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
)

type Game struct{}

func init() {
	space.Init(800, 800)
	player.Init()
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Among Sus")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal("Failed to run game: ", err)
	}
}
