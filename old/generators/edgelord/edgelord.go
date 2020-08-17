package edgelord

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
	"github.com/stojg/vector"
)

type Edgelord struct {
}

var _ pkg.Generator = &Edgelord{}

func (g *Edgelord) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	// quat := vector.QuaternionToTarget(&vector.Vector3{0, 1, 0}, &vector.Vector3{1, 1, 1})

	circle := vector.NewVector3(f.GetSine(0), 0, f.GetCos(0))
	phaseCircle := vector.NewVector3(f.GetSine(0.75), 0, f.GetCos(0.75))

	// fmt.Println(quat45)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
				Phase:     0,
			})
		} else {
			circlePos := pixel.Coordinate.Clone()
			circlePos[1] = 0
			circlePos.Sub(vector.NewVector3(0.5, 0, 0.5))
			circlePos.Normalize()

			distance := circlePos.NewSub(circle).Length()
			phase := circlePos.NewSub(phaseCircle).Length()

			distance = utils.Clamp01(1 - distance)

			if circlePos.Dot(circle) > 0 {
				distance *= -1
			}

			if circlePos.Dot(phaseCircle) > 0 {
				phase *= -1
			}

			result = append(result, pkg.GeneratorResult{
				Intensity: distance,
				Phase:     phase,
			})
		}
	}

	return
}

func (g *Edgelord) Name() string {
	return "edgelord"
}
