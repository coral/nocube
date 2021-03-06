package blending

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/utils"
)

func Add(op1 []pkg.Pixel, op2 []pkg.Pixel, val float64) []pkg.Pixel {
	for i, e := range op1 {
		//d := op2[i].Color.Add(&e.Color
		r := utils.Clamp01(e.Color[0] + op2[i].Color[0])
		g := utils.Clamp01(e.Color[1] + op2[i].Color[1])
		b := utils.Clamp01(e.Color[2] + op2[i].Color[2])

		op2[i].Color[0] = r
		op2[i].Color[1] = g
		op2[i].Color[2] = b
	}

	return op2

}
