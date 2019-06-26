package beat

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type Beat struct {
}

var _ pkg.Generator = &Beat{}

func (g *Beat) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: f.Phase,
			})

		}
	}

	return
}

func (g *Beat) Name() string {
	return "beat"
}
