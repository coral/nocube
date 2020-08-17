package fill

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type Fill struct {
}

var _ pkg.Step = &Fill{}

func (g *Fill) Init() {

}

func (g *Fill) Name() string {
	return "fill"
}

func (g *Fill) Type() string {
	return "static"
}

func (g *Fill) Gen(pixels []pkg.Pixel, f *frame.F) []pkg.Pixel {

	for i, _ := range pixels {
		pixels[i].Color = vector.Vector3{1, 1, 1}
	}

	return pixels
}
