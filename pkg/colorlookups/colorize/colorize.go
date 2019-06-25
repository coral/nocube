package colorize

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/stojg/vector"
)

type Colorize struct {
	Hue        float64
	Saturation float64
	Lightness  float64
}

var _ pkg.ColorLookup = &Colorize{
	Hue:        270.0,
	Saturation: 1.0,
	Lightness:  1.0,
}

func (g *Colorize) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for _, pixel := range generatorResults {
		col := colorful.Hsl(0.0, 1.0, pixel.Intensity*0.5)
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
