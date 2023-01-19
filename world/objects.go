package world

import (
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

type Object struct {
	Obj  *resolv.Object
	Type string
}

var Objects []Object

func NewObject(x float64, y float64, Type string) {
	switch Type {
	case "platform":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 124, "object", Type), Type})
	case "column":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 31, 124, "object", Type), Type})

		for _, o := range Objects {
			if o.Obj.X == x && o.Obj.Y == y {
				space.Space.Add(o.Obj)
			}
		}
	case "spikes":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})

		for _, o := range Objects {
			if o.Obj.X == x && o.Obj.Y == y {
				space.Space.Add(o.Obj)
			}
		}
	case "turretbase":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})

		for _, o := range Objects {
			if o.Obj.X == x && o.Obj.Y == y {
				space.Space.Add(o.Obj)
			}
		}
	case "longplatform":
		Type = "turretbase"
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})

		for _, o := range Objects {
			if (o.Obj.X == x || o.Obj.X == x+124 || o.Obj.X == x+248) && o.Obj.Y == y {
				space.Space.Add(o.Obj)
			}
		}
	case "columnspikes":
		Type = "column"
		Objects = append(Objects, Object{resolv.NewObject(x, y, 31, 124, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+155, y, 31, 124, "object", Type), Type})
		Type = "spikes"
		Objects = append(Objects, Object{resolv.NewObject(x+31, y+62, 124, 62, "object", Type), Type})

		for _, o := range Objects {
			if (o.Obj.X == x || o.Obj.X == x+31 || o.Obj.X == x+155) && (o.Obj.Y == y || o.Obj.Y == y+62) {
				space.Space.Add(o.Obj)
			}
		}
	case "hill":
		Type = "turretbase"
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})

		for _, o := range Objects {
			if (o.Obj.X == x || o.Obj.X == x+124 || o.Obj.X == x+248) && (o.Obj.Y == y || o.Obj.Y == y-62) {
				space.Space.Add(o.Obj)
			}
		}
	case "spikehill":
		Type = "turretbase"
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+124, y, 124, 62, "object", Type), Type})
		Objects = append(Objects, Object{resolv.NewObject(x+248, y, 124, 62, "object", Type), Type})
		Type = "spikes"
		Objects = append(Objects, Object{resolv.NewObject(x+124, y-62, 124, 62, "object", Type), Type})

		for _, o := range Objects {
			if (o.Obj.X == x || o.Obj.X == x+124 || o.Obj.X == x+248) && (o.Obj.Y == y || o.Obj.Y == y-62) {
				space.Space.Add(o.Obj)
			}
		}
	}
}
