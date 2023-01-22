package world

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type Object struct {
	Obj  *resolv.Object
	Type string
}

var (
	Objects          []Object
	platform, _, _   = ebitenutil.NewImageFromFile("res/terrain/platform.png")
	column, _, _     = ebitenutil.NewImageFromFile("res/terrain/column-1.png")
	spikes, _, _     = ebitenutil.NewImageFromFile("res/terrain/spikes.png")
	leftspikes, _, _ = ebitenutil.NewImageFromFile("res/terrain/left-spikes.png")
	turretbase, _, _ = ebitenutil.NewImageFromFile("res/terrain/turretbase.png")
)

func (o *Object) Update() {
	o.Obj.X -= Speed

	if o.Obj.X <= -o.Obj.W {
		if o.Type == "platform" {
			o.Obj.X = 1240
		} else {
			o.Remove()
		}
	}

	o.Obj.Update()
}

func (o Object) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.Obj.X, o.Obj.Y)

	switch o.Type {
	case "platform":
		screen.DrawImage(platform, op)

	case "column":
		screen.DrawImage(column, op)

	case "spikes":
		screen.DrawImage(spikes, op)

	case "leftspikes":
		screen.DrawImage(leftspikes, op)

	case "turretbase":
		screen.DrawImage(turretbase, op)
	}
}

func (o *Object) Remove() {
	tmp := []Object{}

	for _, O := range Objects {
		if o.Obj.X == O.Obj.X && o.Type == O.Type {
			continue
		}
		tmp = append(tmp, O)
	}

	space.Space.Remove(o.Obj)

	Objects = []Object{}
	Objects = tmp
}

func NewObject(x float64, y float64, Type string) {
	newObjects := []Object{}
	coinChance := rand.Intn(4)

	switch Type {
	case "platform":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 124, "object", Type), Type})

	case "column":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 31, 124, "object", Type), Type})

		if coinChance == 1 || coinChance == 2 {
			NewCoin(x-2, y-310)
		}

	case "spikes":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})

		if coinChance == 1 || coinChance == 2 {
			NewCoin(x+39.5, y-168)
		}

	case "leftspikes":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 62, 124, "object", Type), Type})

		coinChance = 1
		if coinChance == 1 {
			NewCoin(x-62, y+6)
			NewCoin(x-62, y+68)
		}

	case "turretbase":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})

	case "longplatform":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})

		coinChance = rand.Intn(10)

		if coinChance == 1 {
			NewCoin(x+8, y-48)
			NewCoin(x+8, y-108)
			NewCoin(x+70, y-48)
			NewCoin(x+70, y-108)
			NewCoin(x+132, y-48)
			NewCoin(x+132, y-108)
			NewCoin(x+194, y-48)
			NewCoin(x+194, y-108)
			NewCoin(x+256, y-48)
			NewCoin(x+256, y-108)
			NewCoin(x+318, y-48)
			NewCoin(x+318, y-108)
		}

	case "columnspikes":
		Type = "column"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 31, 124, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+155, y, 31, 124, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+31, y+62, 124, 62, "object", Type), Type})

		if coinChance == 2 {
			NewCoin(x+70.5, y+5)
		}

	case "hill":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})

		if coinChance == 3 {
			NewCoin(x+165, y-117)
		}

	case "spikehill":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})

		if coinChance == 3 {
			NewCoin(x+165, y-175)
		}

	case "platformspikes":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y-62, 124, 62, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})

		if coinChance == 3 {
			NewCoin(x+138, y-52)
			NewCoin(x+188, y-52)
		}

	case "highspikeshill":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y-62, 124, 62, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-124, 124, 62, "object", Type), Type})

		if coinChance == 2 {
			NewCoin(x+165, y-285)
		}

	case "tallhill":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-124, 124, 62, "object", Type), Type})

		if coinChance == 3 {
			NewCoin(x+38, y-114)
			NewCoin(x+162, y-176)
			NewCoin(x+286, y-114)
		}

	case "platformcolumnspikes":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})

		Type = "column"
		newObjects = append(newObjects, Object{resolv.NewObject(x+92, y-124, 31, 124, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+247, y-124, 31, 124, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+123, y-62, 124, 62, "object", Type), Type})

		if coinChance == 2 {
			NewCoin(x+162.5, y-119)
		}
	}

	if len(newObjects) > 0 {
		for _, o := range newObjects {
			Objects = append(Objects, o)
			space.Space.Add(o.Obj)
		}
	}
}
