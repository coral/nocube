package sense

import (
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
)

type Sense struct {
}

var _ pkg.Generator = &Sense{}

func (g *Sense) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	lowValue := utils.Clamp01(0)
	highValue := utils.Clamp01(0.1)
	numFFT := len(f.FFT)

	low := int64(math.Floor(float64(numFFT) * lowValue))
	high := int64(math.Floor(float64(numFFT) * highValue))

	var total float64
	for _, value := range f.FFT[low:high] {

		total += value

	}

	sum := total / float64(len(f.FFT[low:high]))

	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: sum / 40,
			})

		}
	}

	return
}

func (g *Sense) Name() string {
	return "sense"
}
