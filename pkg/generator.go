package pkg

import "github.com/coral/nocube/pkg/frame"

type GeneratorParameters struct {
	values  map[string]float64
	toggles map[string]bool
}

type GeneratorResult struct {
	Intensity float64
	Phase     float64
}

type Generator interface {
	Generate(pixels []Pixel, f *frame.F, parameters GeneratorParameters) []GeneratorResult
}
