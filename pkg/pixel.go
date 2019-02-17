package pkg

import (
	//"math"

	"github.com/stojg/vector"
)

type Pixel struct {
	Index      int64          `json:"I"`
	Active     bool           `json:"A"`
	Coordinate vector.Vector3 `json:"C"`
	Normal     vector.Vector3 `json:"N"`
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
	return p.Coordinate[2] == 0.0
}

func (p *Pixel) IsBack() bool {
	return p.Coordinate[2] == 1.0
}

func (p *Pixel) PositionInTube() float64 {
	axis := p.Coordinate.Clone()
	axis[0] *= p.Normal[0]
	axis[1] *= p.Normal[1]
	axis[2] *= p.Normal[2]
	return axis.Length()
}

func (p *Pixel) PositionOnNormal(n vector.Vector3) float64 {
	axis := p.Coordinate.Clone()
	axis[0] *= n[0]
	axis[1] *= n[1]
	axis[2] *= n[2]
	res := axis.Length()

	if vector.NewVector3(1, 1, 1).Dot(&n) < 0 {
		res = 1 - res
	}

	return res
}
