package blending

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/utils"
)

func Subtract(op1 []pkg.ColorLookupResult, op2 []pkg.ColorLookupResult, val float64) []pkg.ColorLookupResult {
	for i, e := range op1 {

		r := utils.Clamp01(e.Color[0] - utils.Clamp01(op2[i].Color[0]))
		g := utils.Clamp01(e.Color[1] - utils.Clamp01(op2[i].Color[1]))
		b := utils.Clamp01(e.Color[2] - utils.Clamp01(op2[i].Color[2]))

		op2[i].Color[0] = r
		op2[i].Color[1] = g
		op2[i].Color[2] = b
	}

	return op2

}
