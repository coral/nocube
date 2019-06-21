package strobe

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
)

type Strobe struct {
}

var _ pkg.Generator = &Strobe{}

func (g *Strobe) Generate(pixels []pkg.Pixel, f *frame.F, parameters pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	_, r := f.GetSegment(4)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: utils.Threshold(r, 0.90),
			})

		}
	}

	return
}

func (g *Strobe) Settings() {

}
