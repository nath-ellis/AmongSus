package world

var objTicker int = 0

func Update() {
	if objTicker <= 0 {
		//r := rand.Intn(5) // more types later

		NewObject(1240, 352, "column")

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

				Objects = []Object{}
				Objects = tmp
			}
		}

		o.Obj.Update()
	}
}
