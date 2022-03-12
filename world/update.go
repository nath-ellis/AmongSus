package world

import (
	"math/rand"

	"github.com/nath-ellis/AmongSus/space"
)

var objTicker int = 0

func Update() {
	if objTicker <= 0 {
		r := rand.Intn(3) // more types later

		switch r {
		case 0:
			NewObject(1240, 352, "column")
		case 1:
			NewObject(1240, 414, "spikes")
		case 2:
			NewObject(1240, 414, "turretbase")
		}

		objTicker = 300
	} else {
		objTicker -= 1
	}

	for _, o := range Objects {
		switch o.Type {
		case "platform":
			o.Obj.X -= Speed

			if o.Obj.X <= -124 {
				o.Obj.X = 1240
			}
		case "column":
			o.Obj.X -= Speed

			if o.Obj.X <= -32 { // removes it
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
		case "spikes":
			o.Obj.X -= Speed

			if o.Obj.X <= -125 { // removes it
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
		case "turretbase":
			o.Obj.X -= Speed

			if o.Obj.X <= -250 { // removes it
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
		}

		o.Obj.Update()
	}
}
