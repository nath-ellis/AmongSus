package player

func (p *PlayerData) ChangeColour(direction string) {
	var blackUnlocked, whiteUnlocked bool

	for _, i := range SavedData.Items {
		if i.Name == "white" {
			whiteUnlocked = i.Owned
		}
		if i.Name == "black" {
			blackUnlocked = i.Owned
		}
	}

	if direction == "right" {
		switch p.Colour {
		case "lime":
			p.Colour = "cyan"
		case "cyan":
			p.Colour = "yellow"
		case "yellow":
			if whiteUnlocked {
				p.Colour = "white"
			} else if blackUnlocked {
				p.Colour = "black"
			} else {
				p.Colour = "lime"
			}
		case "white":
			if blackUnlocked {
				p.Colour = "black"
			} else {
				p.Colour = "lime"
			}
		case "black":
			p.Colour = "lime"
		}
	}

	if direction == "left" {
		switch p.Colour {
		case "lime":
			if blackUnlocked {
				p.Colour = "black"
			} else if whiteUnlocked {
				p.Colour = "white"
			} else {
				p.Colour = "yellow"
			}
		case "black":
			if whiteUnlocked {
				p.Colour = "white"
			} else {
				p.Colour = "yellow"
			}
		case "white":
			p.Colour = "yellow"
		case "yellow":
			p.Colour = "cyan"
		case "cyan":
			p.Colour = "lime"
		}
	}

	p.LoadSprites()
}
