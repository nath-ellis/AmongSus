package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	if player.Player.State == "menu" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { // FINISH
			posX, posY := ebiten.CursorPosition()

			if (posX > 5 && posX < 37) && (posY > 105 && posY < 137) {
				player.ChangeColour()
			} else if (posX > 45 && posX < 77) && (posY > 105 && posY < 137) {
				player.ChangeColour()
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
			player.Player.Obj.X = 50
			player.Player.Obj.Y = -100

			for _, o := range world.Objects {
				if o.Type != "platform" {
					tmp := []world.Object{}

					for _, O := range world.Objects {
						if o.Obj.X == O.Obj.X && o.Type == O.Type {
							continue
						}
						tmp = append(tmp, O)
					}

					space.Space.Remove(o.Obj)

					world.Objects = []world.Object{}
					world.Objects = tmp
				}
			}

			player.Player.State = "game"
		}
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
