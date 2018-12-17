package dummy

import (
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/stojg/vector"
)

type Dummy struct {
}

var _ pkg.ColorLookup = &Dummy{}

func clamp01(v float64) float64 {
	if v < 0.0 {
		return 0.0
	}
	if v > 1.0 {
		return 1.0
	}

	return v
}

func (g *Dummy) Lookup(generatorResults []pkg.GeneratorResult, t float64, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for _, pixel := range generatorResults {
		clampedValue := clamp01(pixel.Value)
		results = append(results, pkg.ColorLookupResult{
			Color: vector.Vector3{
				clampedValue,
				clampedValue * math.Sin(t),
				clampedValue,
			},
		})
	}

	return
}
