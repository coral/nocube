package xd

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type Xd struct {
}

var _ pkg.Generator = &Xd{}

func (g *Xd) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	// quat := vector.QuaternionToTarget(&vector.Vector3{0, 1, 0}, &vector.Vector3{1, 1, 1})

	segment, remainder := f.GetSegment(4)
	var normal vector.Vector3

	// fmt.Println(quat45)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Value: 0,
			})
		} else {
			isActive := false
			var base float64
			// fmt.Println("Segment:", segment)
			switch segment {
			case 0:
				isActive = pixel.IsLeft()
				normal = vector.Vector3{0, 0, 1}
			case 1:
				isActive = pixel.IsFront()
				normal = vector.Vector3{-1, 0, 0}
			case 2:
				isActive = pixel.IsRight()
				normal = vector.Vector3{0, 0, -1}
			case 3:
				isActive = pixel.IsBack()
				normal = vector.Vector3{1, 0, 0}
			}

			if isActive {
				pos := pixel.PositionOnNormal(normal)
				base = remainder - pos
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
