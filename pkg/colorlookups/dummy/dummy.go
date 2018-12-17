package dummy

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type Dummy struct {
}

var _ pkg.ColorLookup = &Dummy{}

func (g *Dummy) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for _, pixel := range generatorResults {
		r := 0.0
		g := 0.0
		b := 0.0
		if pixel.Value < 0 {
			r = 255
			g = 255
			b = 255
		} else {
			r = 0
			g = 0
			b = 0
			// g = pixel.Value
		}
		// clampedValue := clamp01(pixel.Value)
		results = append(results, pkg.ColorLookupResult{
			Color: vector.Vector3{
				r,
				g,
				b,
			},
		})
	}

	return
}
