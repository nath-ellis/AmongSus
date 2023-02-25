package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/nath-ellis/AmongSus/player"
	"github.com/solarlune/resolv"
)

type ShopItem struct {
	Button
	Name  string
	Owned bool
	Price int
}

func (s ShopItem) Draw(screen *ebiten.Image, scale bool) { // add frame
	op := &ebiten.DrawImageOptions{}
	if scale {
		op.GeoM.Scale(2, 2)
	}
	op.GeoM.Translate(s.Obj.X, s.Obj.Y)

	screen.DrawImage(itemBG, op)
	op.GeoM.Translate(30, 30)

	mouseX, mouseY := ebiten.CursorPosition()

	if float64(mouseX) > s.Obj.X && float64(mouseX) < s.Obj.X+s.Obj.W &&
		float64(mouseY) > s.Obj.Y && float64(mouseY) < s.Obj.Y+s.Obj.H {
		screen.DrawImage(s.OnHoverSprite, op)

		op.GeoM.Reset()
		op.GeoM.Translate(s.Obj.X, s.Obj.Y)
		screen.DrawImage(dark, op)

		// Draw price
		op.GeoM.Reset()
		op.GeoM.Translate(s.Obj.X+10, s.Obj.Y+50)

		screen.DrawImage(coin, op)

		if s.Owned {
			text.Draw(screen, "--", VCROSDMono, int(s.Obj.X+55), int(s.Obj.Y+100), color.White)
		} else {
			text.Draw(screen, fmt.Sprint(s.Price), VCROSDMono, int(s.Obj.X+55), int(s.Obj.Y+100), color.White)
		}
	} else {
		screen.DrawImage(s.Sprite, op)
	}
}

var shopItems []ShopItem

func initItems() {
	shopItems = []ShopItem{}

	x, y := float64(50), float64(100)

	// Add shop items
	for _, i := range player.SavedData.Items {
		var sprite *ebiten.Image

		switch i.Name {
		case "white":
			sprite = whiteIdle
		case "black":
			sprite = blackIdle
		}

		shopItems = append(shopItems, ShopItem{
			Button: Button{
				Obj:           resolv.NewObject(x, y, 150, 150),
				Sprite:        sprite,
				OnHoverSprite: sprite,
			},
			Name:  i.Name,
			Owned: i.Owned,
			Price: i.Price,
		})

		x += 150

		if x >= 1000 {
			x = 50
			y += 150
		}
	}
}

func checkItems() {
	for _, s := range shopItems {
		if s.IsPressed() && !s.Owned {
			for i, k := range player.SavedData.Items {
				if k.Name == s.Name && player.Player.Coins >= k.Price {
					// purchase item
					player.SavedData.Items[i].Owned = true
					player.Player.Coins -= k.Price
					player.SavedData.Save()
				}
			}
			initItems() // re-initialise the shop
		}
	}
}
