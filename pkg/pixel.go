package pkg

import "github.com/stojg/vector"

type Pixel struct {
	Active     bool
	Coordinate vector.Vector3
}

// Sides
func (p *Pixel) IsTop() bool {
	return p.Coordinate[1] == 0.0
}

func (p *Pixel) IsBottom() bool {
	return p.Coordinate[1] == 1.0
}

func (p *Pixel) IsLeft() bool {
	return p.Coordinate[0] == 0.0
}

func (p *Pixel) IsRight() bool {
	return p.Coordinate[0] == 1.0
}

func (p *Pixel) IsFront() bool {
	return p.Coordinate[2] == 1.0
}

func (p *Pixel) IsBack() bool {
	return p.Coordinate[2] == 0.0
}
