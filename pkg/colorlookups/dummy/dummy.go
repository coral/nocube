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
		if pixel.Intensity > 0 && pixel.Phase > 0 {
			r = pixel.Phase
			g = pixel.Intensity
			b = pixel.Intensity
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
