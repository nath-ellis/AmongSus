package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/nath-ellis/AmongSus/world"
)

type Game struct{}

func init() {
	rand.Seed(time.Now().Unix())

	space.Init(1200, 600)
	player.Init()

	world.Init()
}

func (g *Game) Update() error {
	if player.Player.State == "menu" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			posX, posY := ebiten.CursorPosition()

			if (posX > 5 && posX < 37) && (posY > 105 && posY < 137) { // Left Button
				player.ChangeColour(ebiten.Key(ebiten.MouseButtonLeft))
			} else if (posX > 45 && posX < 77) && (posY > 105 && posY < 137) { // Right Button
				player.ChangeColour(ebiten.Key(ebiten.MouseButtonRight))
			} else {
				player.Player.State = "game"
			}
		}
		player.ColourSelectorCtl()
	} else if player.Player.State == "game" {
		player.Controls()
		world.Update()
	} else if player.Player.State == "gameOver" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			posX, posY := ebiten.CursorPosition()

			if (posX > 5 && posX < 37) && (posY > 105 && posY < 137) { // Left Button
				player.ChangeColour(ebiten.Key(ebiten.MouseButtonLeft))
			} else if (posX > 45 && posX < 77) && (posY > 105 && posY < 137) { // Right Button
				player.ChangeColour(ebiten.Key(ebiten.MouseButtonRight))
			} else {
				// Reset the game
				player.Player.Obj.X = 50
				player.Player.Obj.Y = -100

				world.Speed = 2
				world.ObjTicker = 0
				world.MaxObjTicker = 250

				// Remove all objects
				for _, o := range world.Objects {
					space.Space.Remove(o.Obj)
				}

				world.Objects = []world.Object{}
				world.NewObject(0, 476, "platform")
				world.NewObject(124, 476, "platform")
				world.NewObject(248, 476, "platform")
				world.NewObject(372, 476, "platform")
				world.NewObject(496, 476, "platform")
				world.NewObject(620, 476, "platform")
				world.NewObject(744, 476, "platform")
				world.NewObject(868, 476, "platform")
				world.NewObject(992, 476, "platform")
				world.NewObject(1116, 476, "platform")
				world.NewObject(1240, 476, "platform")

				player.Player.State = "game"
			}
		}
		player.ColourSelectorCtl()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if player.Player.State == "menu" {
		world.DrawMenu(screen)
		player.DrawColourSelector(screen)
	} else if player.Player.State == "game" {
		world.Draw(screen)
		player.Draw(screen)
	} else if player.Player.State == "gameOver" {
		world.DrawGameOver(screen)
		player.DrawColourSelector(screen)
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
