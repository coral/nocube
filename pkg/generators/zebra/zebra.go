package zebra

import (
	"fmt"
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/stojg/vector"
)

type Zebra struct {
}

var _ pkg.Generator = &Zebra{}

func (g *Zebra) Generate(pixels []pkg.Pixel, t float64, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	fmt.Println("Zebra generating with t", t)
	// quat := vector.QuaternionToTarget(&vector.Vector3{0, 1, 0}, &vector.Vector3{1, 1, 1})

	// Make identity vector
	// quat := &vector.Quaternion{1, 0, 0, 0}
	// quat.RotateByVector(&vector.Vector3{0, 1, 0})
	quat45up := vector.QuaternionFromAxisAngle(&vector.Vector3{0, -1, 0}, math.Pi/4*t)
	quat45right := vector.QuaternionFromAxisAngle(&vector.Vector3{1, 0, 0}, math.Pi/4*t*0.3)
	quat := quat45up.NewMultiply(quat45right)
	// fmt.Println(quat45)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Value: 0,
			})
		} else {
			coord := pixel.Coordinate.Clone()
			// fmt.Println("Z BEFORE:", coord[2])
			coord = coord.Rotate(quat)
			// fmt.Println("Z AFTER:", coord[2])
			result = append(result, pkg.GeneratorResult{
				Value: math.Mod(coord[2]*5, 1.0),
			})
		}
	}

	return
}
