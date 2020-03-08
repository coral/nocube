package colorize

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/stojg/vector"
)

type Colorize struct {
}

var _ pkg.ColorLookup = &Colorize{}

func (g *Colorize) Init() {

}

func (g *Colorize) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, p pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	hue := p.Data.GetScopedFloat64(p.Name, g.Name(), "hue")
	saturation := p.Data.GetScopedFloat64(p.Name, g.Name(), "saturation")
	for _, pixel := range generatorResults {
		col := colorful.Hsv(hue*360, saturation, pixel.Intensity)
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

func (g *Colorize) Name() string {
	return "colorize"
}
