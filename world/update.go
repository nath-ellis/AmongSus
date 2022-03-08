package world

func Update() {
	for _, o := range Objects {
		o.Obj.X -= Speed

		if o.Obj.X <= -124 {
			o.Obj.X = 1240
		}

		o.Obj.Update()
	}
}
