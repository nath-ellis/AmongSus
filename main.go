package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/nath-ellis/AmongSus/ui"
	"github.com/nath-ellis/AmongSus/world"
)

type Game struct{}

func init() {
	rand.Seed(time.Now().Unix())

	space.Init(1200, 600)
	player.Player.Init()

	world.Init()
	ui.Init()
}

func (g *Game) Update() error {
	if player.Player.State == "menu" {
		ui.UpdateMenu()
	} else if player.Player.State == "game" {
		player.Player.Controls()
		world.Update()
	} else if player.Player.State == "gameOver" {
		ui.UpdateGameOver()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if player.Player.State == "menu" {
		ui.DrawMenu(screen, fmt.Sprint(player.Player.Coins))
	} else if player.Player.State == "game" {
		world.Draw(screen)
		player.Player.Draw(screen)
	} else if player.Player.State == "gameOver" {
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
