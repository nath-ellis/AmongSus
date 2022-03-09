package world

func Update() {
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
