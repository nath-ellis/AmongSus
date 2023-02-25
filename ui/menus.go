package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/nath-ellis/AmongSus/world"
)

func UpdateMenu() {
	if LeftColourBtn.IsPressed() {
		player.Player.ChangeColour("left")
	} else if RightColourBtn.IsPressed() {
		player.Player.ChangeColour("right")
	} else if shopBtn.IsPressed() {
		if shopOpen {
			shopOpen = false
			shopBtn.Sprite = shopLight
			shopBtn.OnHoverSprite = shopDark
		} else {
			shopOpen = true
			shopBtn.Sprite = xLight
			shopBtn.OnHoverSprite = xDark
			initItems()
		}
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !shopOpen {
		player.Player.State = "game"
	}

	if shopOpen {
		checkItems()
	}
}

func DrawMenu(screen *ebiten.Image, coins string) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.8, 0.8)
	screen.DrawImage(menubg, op)

	if shopOpen {
		for _, s := range shopItems {
			s.Draw(screen, false)
		}
	} else {
		op.GeoM.Reset()
		op.GeoM.Scale(0.5, 0.5)
		op.GeoM.Translate(340, 200)

		screen.DrawImage(title, op)

		op.GeoM.Reset()
		op.GeoM.Scale(0.3, 0.3)
		op.GeoM.Translate(490, 400)
		screen.DrawImage(click, op)

		op.GeoM.Reset()
		op.GeoM.Translate(5, 5)

		screen.DrawImage(player.Player.IdleSprite, op)

		LeftColourBtn.Draw(screen, true)
		RightColourBtn.Draw(screen, true)
	}

	drawCoinAmount(screen, coins)

	shopBtn.Draw(screen, false)
}

func UpdateGameOver() {
	// Reset the game
	reset := func() {
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

		world.Coins = []world.Coin{}
	}

	if LeftColourBtn.IsPressed() {
		player.Player.ChangeColour("left")
	} else if RightColourBtn.IsPressed() {
		player.Player.ChangeColour("right")
	} else if homeBtn.IsPressed() {
		reset()
		player.Player.State = "menu"
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		reset()
		player.Player.State = "game"
	}
}

func DrawGameOver(screen *ebiten.Image, coins string) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.8, 0.8)
	screen.DrawImage(menubg, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(350, 200)

	screen.DrawImage(gameOver, op)

	op.GeoM.Reset()
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(490, 400)
	screen.DrawImage(click, op)

	drawCoinAmount(screen, coins)

	op.GeoM.Reset()
	op.GeoM.Translate(5, 5)

	screen.DrawImage(player.Player.IdleSprite, op)

	LeftColourBtn.Draw(screen, true)
	RightColourBtn.Draw(screen, true)
	homeBtn.Draw(screen, false)
}

func drawCoinAmount(screen *ebiten.Image, coins string) {
	op := &ebiten.DrawImageOptions{}

	if shopOpen {
		op.GeoM.Translate(5, 5)
		text.Draw(screen, coins, VCROSDMono, 60, 55, color.White)
	} else {
		op.GeoM.Translate(5, 540)
		text.Draw(screen, coins, VCROSDMono, 60, 590, color.White)
	}

	screen.DrawImage(coin, op)
}
