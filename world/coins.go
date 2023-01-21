package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type Coin struct {
	Obj            *resolv.Object
	Sprites        []*ebiten.Image
	AnimationFrame int
	AnimationFPS   int
}

var (
	Coins            []Coin
	coinFrame1, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_1.png")
	coinFrame2, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_2.png")
	coinFrame3, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_3.png")
	coinFrame4, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_4.png")
	coinFrame5, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_5.png")
	coinFrame6, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_6.png")
	coinFrame7, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_7.png")
	coinFrame8, _, _ = ebitenutil.NewImageFromFile("res/coins/coin_8.png")
)

func (c *Coin) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.Obj.X, c.Obj.Y)

	screen.DrawImage(c.Sprites[c.AnimationFrame], op)
}

func (c *Coin) Update() {
	c.Obj.X -= Speed
	c.Obj.Update()

	if c.Obj.X <= -c.Obj.W {
		c.Remove()
	}
}

func (c *Coin) Remove() {
	tmp := []Coin{}

	for _, C := range Coins {
		if c.Obj.X == C.Obj.X && c.Obj.Y == C.Obj.Y {
			continue
		}
		tmp = append(tmp, C)
	}

	space.Space.Remove(c.Obj)

	Coins = []Coin{}
	Coins = tmp
}

func NewCoin(x float64, y float64) {
	newCoin := Coin{
		resolv.NewObject(x, y, 45, 48, "coin"),
		[]*ebiten.Image{
			coinFrame1,
			coinFrame2,
			coinFrame3,
			coinFrame4,
			coinFrame5,
			coinFrame6,
			coinFrame7,
			coinFrame8,
		},
		0,
		5,
	}

	Coins = append(Coins, newCoin)
	space.Space.Add(newCoin.Obj)
}
