package dummy

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
	"github.com/stojg/vector"
)

type Dummy struct {
}

var _ pkg.ColorLookup = &Dummy{}

func (g *Dummy) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for _, pixel := range generatorResults {

		d := utils.Crush(pixel.Intensity, 0.1)
		r := d
		g := d
		b := d

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
