package blending

import "github.com/coral/nocube/pkg"

func Screen(op1 []pkg.Pixel, op2 []pkg.Pixel, val float64) []pkg.Pixel {

	for i, e := range op1 {
		//d := op2[i].Color.Add(&e.Color
		r := 1 - (1-e.Color[0])*(1-op2[i].Color[0])
		g := 1 - (1-e.Color[1])*(1-op2[i].Color[1])
		b := 1 - (1-e.Color[2])*(1-op2[i].Color[2])

		if r > 1.0 {
			r = 1.0
		}
		if g > 1.0 {
			g = 1.0
		}
		if b > 1.0 {
			b = 1.0
		}

		op2[i].Color[0] = r
		op2[i].Color[1] = g
		op2[i].Color[2] = b
	}

	return op2

}
