package runner

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type Runner struct {
}

var _ pkg.Generator = &Runner{}

func (g *Runner) Generate(pixels []pkg.Pixel, f *frame.F, p pkg.GeneratorParameters) (result []pkg.GeneratorResult) {

	for _, pixel := range pixels {
		_ = pixel
		result = append(result, pkg.GeneratorResult{
			Intensity: 1,
		})

	}
	return
}

func (g *Runner) Name() string {
	return "runner"
}
