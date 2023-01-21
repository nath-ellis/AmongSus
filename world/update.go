package world

import (
	"math/rand"
)

var (
	ObjTicker    int = 0
	MaxObjTicker int = 250
)

func Update() {
	if ObjTicker <= 0 {
		r := rand.Intn(11)

		switch r {
		case 0:
			NewObject(1240, 352, "column")
		case 1:
			NewObject(1240, 414, "spikes")
		case 2:
			NewObject(1240, 414, "turretbase")
		case 3:
			NewObject(1240, 414, "longplatform")
		case 4:
			NewObject(1240, 352, "columnspikes")
		case 5:
			NewObject(1240, 414, "hill")
		case 6:
			NewObject(1240, 414, "spikehill")
		case 7:
			NewObject(1240, 414, "platformspikes")
		case 8:
			NewObject(1240, 414, "longplatform")
			NewObject(1240, 352, "spikehill")
		case 9:
			NewObject(1240, 414, "longplatform")
			NewObject(1240, 352, "hill")
		case 10:
			NewObject(1240, 414, "longplatform")
			NewObject(1332, 290, "columnspikes")
		}

		ObjTicker = MaxObjTicker

		if r%2 == 0 {
			chance := rand.Intn(6)
			if Speed < MaxSpeed && chance == 1 {
				Speed += 1
			}

			if Speed >= 5 {
				MaxObjTicker = 100
			} else if Speed >= 4 {
				MaxObjTicker = 150
			} else if Speed >= 3 {
				MaxObjTicker = 200
			}
		}
	} else {
		ObjTicker -= 1
	}

	for _, o := range Objects {
		o.Update()
	}
}
