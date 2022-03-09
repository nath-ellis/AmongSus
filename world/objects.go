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
		Objects = append(Objects, Object{resolv.NewObject(x, y, 124, 124, "object"), Type})
	case "column":
		Objects = append(Objects, Object{resolv.NewObject(x, y, 31, 124, "object"), Type})
	}

	for _, o := range Objects {
		space.Space.Add(o.Obj)
	}
}
