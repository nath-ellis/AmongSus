package player

func (p *PlayerData) ChangeColour(direction string) {
	if direction == "right" {
		if p.Colour == "lime" {
			p.Colour = "cyan"
		} else if p.Colour == "cyan" {
			p.Colour = "yellow"
		} else if p.Colour == "yellow" {
			p.Colour = "lime"
		}
	}

	if direction == "left" {
		if p.Colour == "lime" {
			p.Colour = "yellow"
		} else if p.Colour == "cyan" {
			p.Colour = "lime"
		} else if p.Colour == "yellow" {
			p.Colour = "cyan"
		}
	}

	p.LoadSprites()
}
