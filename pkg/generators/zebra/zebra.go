package zebra

import (
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type Zebra struct {
}

var _ pkg.Generator = &Zebra{}

func (g *Zebra) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	// quat := vector.QuaternionToTarget(&vector.Vector3{0, 1, 0}, &vector.Vector3{1, 1, 1})
	// Make identity vector
	quat45up := vector.QuaternionFromAxisAngle(&vector.Vector3{0, -1, 0}, math.Pi/4*f.Timepoint)
	quat45right := vector.QuaternionFromAxisAngle(&vector.Vector3{0, 0, 1}, math.Pi/4*f.Timepoint*0.3)
	quat := quat45up.NewMultiply(quat45right)
	// fmt.Println(quat45)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			coord := pixel.Coordinate.Clone()
			// fmt.Println("Z BEFORE:", coord[2])
			coord = coord.Sub(&vector.Vector3{0.5, 0.5, 0.5})
			coord = coord.Rotate(quat)
			// coord = coord.Add(&vector.Vector3{0.5, 0.5, 0.5})

			// fmt.Println("Z AFTER:", coord[2])
			// var v float64
			// if coord[2] > 0.5 {
			// 	v = 1
			// } else {
			// 	v = 0
			// }
			result = append(result, pkg.GeneratorResult{
				// Value: math.Mod(coord[2], 1.0),
				// Value: v,
				// Value: v,
				Intensity: coord[2],
			})

		}
	}

	return
}
