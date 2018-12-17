package xd

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type Xd struct {
}

var _ pkg.Generator = &Xd{}

func (g *Xd) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	// quat := vector.QuaternionToTarget(&vector.Vector3{0, 1, 0}, &vector.Vector3{1, 1, 1})

	segment, remainder := f.GetSegment(4)

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
			case 1:
				isActive = pixel.IsFront()
			case 2:
				isActive = pixel.IsRight()
			case 3:
				isActive = pixel.IsBack()
			}

			if isActive {
				base = remainder * 4
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
