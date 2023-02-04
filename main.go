package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/nath-ellis/AmongSus/ui"
	"github.com/nath-ellis/AmongSus/world"
)

type Game struct{}

func init() {
	space.Init(1200, 600)
	player.Player.Init()

	world.Init()
	ui.Init()
}

func (g *Game) Update() error {
	switch player.Player.State {
	case "menu":
		ui.UpdateMenu()

	case "game":
		player.Player.Controls()
		world.Update()

	case "gameOver":
		ui.UpdateGameOver()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch player.Player.State {
	case "menu":
		ui.DrawMenu(screen, fmt.Sprint(player.Player.Coins))

	case "game":
		world.Draw(screen)
		player.Player.Draw(screen)

	case "gameOver":
		ui.DrawGameOver(screen, fmt.Sprint(player.Player.Coins))
	}
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
