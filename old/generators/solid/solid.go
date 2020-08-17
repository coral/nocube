package solid

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type Solid struct {
}

var _ pkg.Generator = &Solid{}

func (g *Solid) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: 1.0,
			})

		}
	}

	return
}

func (g *Solid) Settings() {

}

func (g *Solid) Name() string {
	return "solid"
}
