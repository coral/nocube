package strobe

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
)

type Strobe struct {
}

var _ pkg.Generator = &Strobe{}

func (g *Strobe) Generate(pixels []pkg.Pixel, f *frame.F, n string, d *data.Data) (result []pkg.GeneratorResult) {

	segment := d.GetInt64(n, "segment")
	_, r := f.GetSegment(segment)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: utils.Threshold(r, 0.92),
			})

		}
	}

	return
}

func (g *Strobe) Settings() {

}
