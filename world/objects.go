package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type Object struct {
	Obj  *resolv.Object
	Type string
}

var Objects []Object

func (o *Object) Update() {
	switch o.Type {
	case "platform":
		o.Obj.X -= Speed

	case "column":
		o.Obj.X -= Speed

	case "spikes":
		o.Obj.X -= Speed

	case "turretbase":
		o.Obj.X -= Speed
	}

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

	switch Type {
	case "platform":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 124, "object", Type), Type})

	case "column":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 31, 124, "object", Type), Type})

	case "spikes":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})

	case "turretbase":
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})

	case "longplatform":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})

	case "columnspikes":
		Type = "column"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 31, 124, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+155, y, 31, 124, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+31, y+62, 124, 62, "object", Type), Type})

	case "hill":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})

	case "spikehill":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})

	case "platformspikes":
		Type = "turretbase"
		newObjects = append(newObjects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x, y-62, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		newObjects = append(newObjects, Object{resolv.NewObject(x+248, y-62, 124, 62, "object", Type), Type})

		Type = "spikes"
		newObjects = append(newObjects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
	}

	if len(newObjects) > 0 {
		for _, o := range newObjects {
			Objects = append(Objects, o)
			space.Space.Add(o.Obj)
		}
	}
}
