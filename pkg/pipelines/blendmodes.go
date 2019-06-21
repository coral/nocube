package pipelines

import (
	"github.com/coral/nocube/pkg"
)

var BlendModes = map[string]bm{
	"add": Add,
}

type bm func([]pkg.ColorLookupResult, []pkg.ColorLookupResult, float64) []pkg.ColorLookupResult

func Add(op1 []pkg.ColorLookupResult, op2 []pkg.ColorLookupResult, val float64) []pkg.ColorLookupResult {

	for i, e := range op1 {
		//d := op2[i].Color.Add(&e.Color
		r := e.Color[0] + op2[i].Color[0]
		g := e.Color[1] + op2[i].Color[1]
		b := e.Color[2] + op2[i].Color[2]

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
		op2[i].Color[0] = g
		op2[i].Color[0] = b
	}

	return op2

}
