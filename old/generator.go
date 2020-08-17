package pkg

import (
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/frame"
)

type GeneratorParameters struct {
	Data *data.Data
	Name string
}

type GeneratorResult struct {
	Intensity float64
	Phase     float64
}

type Generator interface {
	Generate(pixels []Pixel, f *frame.F, p GeneratorParameters) []GeneratorResult
	Name() string
}
