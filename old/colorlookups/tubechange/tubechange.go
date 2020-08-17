package tubechange

import (
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type Tubechange struct {
}

var _ pkg.ColorLookup = &Tubechange{}

func (g *Tubechange) Init() {

}

func (g *Tubechange) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for i, pixel := range generatorResults {

		d := math.Ceil(float64(i+1) / 72)
		col := colorful.Hsl(d*30, 1.0, pixel.Intensity*0.5)
		//d := utils.Crush(pixel.Intensity, 0.1)
		r := col.R
		g := col.G
		b := col.B

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

func (g *Tubechange) Name() string {
	return "tubechange"
}
