package sparkling

import (
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/stojg/vector"
)

type Sparkling struct {
	numSparks       int
	sparkHue        float64
	sparkSaturation float64
	decay           float64
	maxSpeed        float64
	newThreshold    float64
	sparks          []float64
	sparkX          []float64
	pixels          []float64
}

var _ pkg.ColorLookup = &Sparkling{}

func (g *Sparkling) Init() {
	numsparks := 88
	g.numSparks = numsparks
	g.sparks = make([]float64, numsparks)
	g.sparkX = make([]float64, numsparks)
	g.pixels = make([]float64, 864)

	g.sparkHue = 0.2
	g.sparkSaturation = 1
	g.decay = 0.99
	g.maxSpeed = 0.4
	g.newThreshold = 0.01
}

func (g *Sparkling) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, p pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	pixelCount := float64(len(generatorResults))

	g.maxSpeed = p.Data.GetScopedFloat64(p.Name, g.Name(), "speed")
	g.sparkHue = p.Data.GetScopedFloat64(p.Name, g.Name(), "hue")
	g.sparkSaturation = p.Data.GetScopedFloat64(p.Name, g.Name(), "saturation")

	for i, e := range g.pixels {
		g.pixels[i] = e * 0.9
	}

	for i := 0; i < g.numSparks; i++ {
		if g.sparks[i] >= -g.newThreshold && g.sparks[i] <= g.newThreshold {
			g.sparks[i] = (g.maxSpeed / 2) - utils.Random(0, g.maxSpeed)
			g.sparkX[i] = utils.Random(0, pixelCount)
		}

		g.sparks[i] *= g.decay
		g.sparkX[i] += g.sparks[i] * f.Delta

		if g.sparkX[i] >= pixelCount {
			g.sparkX[i] = 0
		}

		if g.sparkX[i] < 0 {
			g.sparkX[i] = pixelCount - 1
		}

		g.pixels[int(math.Floor(g.sparkX[i]))] += g.sparks[i]
	}

	for i, _ := range generatorResults {

		val := g.pixels[i]
		col := colorful.Hsv(g.sparkHue*360, g.sparkSaturation, val*val*10)

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

func (g *Sparkling) Name() string {
	return "sparkling"
}
