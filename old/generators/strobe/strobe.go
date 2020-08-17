package strobe

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
)

type Strobe struct {
}

var _ pkg.Generator = &Strobe{}

func (g *Strobe) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {

	segment := p.Data.GetScopedInt64(p.Name, g.Name(), "segment")
	_, r := f.GetSegment(segment)
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

func (g *Strobe) Name() string {
	return "strobe"
}
