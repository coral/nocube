package allwhite

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type AllWhite struct {
}

var _ pkg.ColorLookup = &AllWhite{}

func (g *AllWhite) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for _, _ = range generatorResults {

		results = append(results, pkg.ColorLookupResult{
			Color: vector.Vector3{
				1,
				1,
				1,
			},
		})
	}

	return
}
