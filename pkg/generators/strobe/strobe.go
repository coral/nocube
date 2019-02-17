package strobe

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type Strobe struct {
}

var _ pkg.Generator = &Strobe{}

func (g *Strobe) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: f.GetSquare(),
			})

		}
	}

	return
}
