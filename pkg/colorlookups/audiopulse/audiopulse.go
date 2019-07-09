package audiopulse

import (
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/stojg/vector"
)

type AudioPulse struct {
	values  []float64
	pixels  []float64
	pos     float64
	lastVal float64
	PIC     PIController
}

type PIController struct {
	Kp    float64
	Ki    float64
	Start float64
	Min   float64
	Max   float64
}

func (p *PIController) Calculate(val float64) float64 {
	p.Start = utils.Clamp(p.Start+val, p.Min, p.Max)
	return math.Max(p.Kp+val+p.Ki+p.Start, 0.3)
}

var _ pkg.ColorLookup = &AudioPulse{}

func (g *AudioPulse) Init() {
	g.pixels = make([]float64, 864)
	g.values = make([]float64, 864)
	g.PIC = PIController{
		Kp:    0.5,
		Ki:    0.35,
		Start: 30,
		Min:   0,
		Max:   400,
	}

}

func (g *AudioPulse) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, p pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {

	pixelcount := len(generatorResults)
	sensitivity := g.PIC.Calculate(0.5 - g.lastVal)
	g.pos = math.Mod(float64(g.pos)+f.Delta*.05, float64(pixelcount))

	g.lastVal = math.Pow(f.PeakFrequencyValue*sensitivity, 2)
	g.values[int(g.pos)] = g.lastVal
	g.pixels[int(g.pos)] = float64(f.PeakFrequency) / 500

	for index, _ := range generatorResults {

		m := pixelcount - index
		i := int(math.Mod(float64(index+int(g.pos)+0), float64(pixelcount)))
		v := g.values[i]
		v = v * v
		h := g.pixels[i] + float64(m)/float64(pixelcount) + f.Phase
		col := colorful.Hsv(h, 1, v)

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

func (g *AudioPulse) Name() string {
	return "audiopulse"
}
