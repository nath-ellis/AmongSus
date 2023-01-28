package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/nath-ellis/AmongSus/world"
	"github.com/solarlune/resolv"
)

type Game struct{}

var (
	Player         player.PlayerData
	LeftColourBtn  world.Button
	RightColourBtn world.Button
)

func init() {
	rand.Seed(time.Now().Unix())

	space.Init(1200, 600)
	Player.Init()

	world.Init()

	leftDark, _, _ := ebitenutil.NewImageFromFile("res/prompts/left_dark.png")
	leftLight, _, _ := ebitenutil.NewImageFromFile("res/prompts/left_light.png")
	rightDark, _, _ := ebitenutil.NewImageFromFile("res/prompts/right_dark.png")
	rightLight, _, _ := ebitenutil.NewImageFromFile("res/prompts/right_light.png")

	LeftColourBtn = world.Button{
		Obj:           resolv.NewObject(5, 100, 32, 32),
		Sprite:        leftDark,
		OnHoverSprite: leftLight,
	}

	RightColourBtn = world.Button{
		Obj:           resolv.NewObject(55, 100, 32, 32),
		Sprite:        rightDark,
		OnHoverSprite: rightLight,
	}
}

func (g *Game) Update() error {
	if Player.State == "menu" {
		if LeftColourBtn.IsPressed() {
			Player.ChangeColour("left")
		} else if RightColourBtn.IsPressed() {
			Player.ChangeColour("right")
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			Player.State = "game"
		}
	} else if Player.State == "game" {
		Player.Controls()
		world.Update()
	} else if Player.State == "gameOver" {
		if LeftColourBtn.IsPressed() {
			Player.ChangeColour("left")
		} else if RightColourBtn.IsPressed() {
			Player.ChangeColour("right")
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			// Reset the game
			Player.Obj.X = 50
			Player.Obj.Y = -100

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

			world.Coins = []world.Coin{}

			Player.State = "game"
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if Player.State == "menu" {
		world.DrawMenu(screen, fmt.Sprint(Player.Coins))

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(5, 5)

		screen.DrawImage(Player.IdleSprite, op)

		LeftColourBtn.Draw(screen, true)
		RightColourBtn.Draw(screen, true)
	} else if Player.State == "game" {
		world.Draw(screen)
		Player.Draw(screen)
	} else if Player.State == "gameOver" {
		world.DrawGameOver(screen, fmt.Sprint(Player.Coins))

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(5, 5)

		screen.DrawImage(Player.IdleSprite, op)

		LeftColourBtn.Draw(screen, true)
		RightColourBtn.Draw(screen, true)
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
