package solid

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type Solid struct {
}

var _ pkg.Generator = &Solid{}

func (g *Solid) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0.2,
			})

		}
	}

	return
}

func (g *Solid) Settings() {

}
