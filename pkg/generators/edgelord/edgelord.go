package edgelord

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type Edgelord struct {
}

var _ pkg.Generator = &Edgelord{}

func (g *Edgelord) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	// quat := vector.QuaternionToTarget(&vector.Vector3{0, 1, 0}, &vector.Vector3{1, 1, 1})


	x := f.GetSine(0)
	y := f.GetCos(0)

	circle := vector.NewVector3(x, 0, y)

	// fmt.Println(quat45)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Value: 0,
			})
		} else {
			isActive := false
			var base float64

			circlePos := pixel.Coordinate.NewAdd(vector.NewVector3(0.5, 0.5, 0.5))
			circlePos.Normalize();

			if circlePos.NewSub(circle).Length() < 0.5 {
				isActive = true
			}

			if isActive {
				base = 1
			}

			result = append(result, pkg.GeneratorResult{
				// Value: math.Mod(coord[2], 1.0),
				// Value: v,
				// Value: v,
				Value: base,
			})
		}
	}

	return
}
