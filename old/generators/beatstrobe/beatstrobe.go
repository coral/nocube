package beatstrobe

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
)

type BeatStrobe struct {
}

var _ pkg.Generator = &BeatStrobe{}

func (g *BeatStrobe) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {
	_, r := f.GetSegment(2)
	isbeat := f.GetBeat(1, 0)
	for _, pixel := range pixels {
		if !pixel.Active {
			result = append(result, pkg.GeneratorResult{
				Intensity: 0,
			})
		} else {
			result = append(result, pkg.GeneratorResult{
				Intensity: utils.Trigger(utils.Threshold(r, 0.92), isbeat),
			})

		}
	}

	return
}

func (g *BeatStrobe) Settings() {

}

func (g *BeatStrobe) Name() string {
	return "beatstrobe"
}
