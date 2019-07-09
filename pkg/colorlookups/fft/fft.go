package fft

import (
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type FFT struct {
}

var _ pkg.ColorLookup = &FFT{}

func (g *FFT) Init() {

}

func (g *FFT) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	p := 0.0

	numPix := len(generatorResults)
	numFFT := len(f.FFT)
	steps := float64(numFFT) / float64(numPix)

	for a, pixel := range generatorResults {
		index := int(math.Round(float64(a) * steps))
		mi := 0.0
		dv := 40
		if a > 30 {
			dv = 15
		}
		if a > 100 {
			dv = 7
		}
		if a > 250 {
			dv = 1
		}
		if index < numFFT {
			mi = f.FFT[index] / float64(dv)
		}
		d := pixel.Intensity
		var r = 0.0
		var g = d * mi
		var b = 0.0

		if f.Phase > 0.90 && f.Confidence > 0.1 {
			r = 1.0 * mi
			g = 1.0 * mi
			b = 1.0 * mi
		}
		//fmt.Println(mi)
		results = append(results, pkg.ColorLookupResult{
			Color: vector.Vector3{
				r,
				g,
				b,
			},
		})
		p = p + 0.001
	}

	return
}

func (g *FFT) Name() string {
	return "fft"
}
