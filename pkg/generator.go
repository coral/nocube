package pkg

import "github.com/coral/nocube/pkg/frame"

type GeneratorParameters struct {
	Speed float64
}

type GeneratorResult struct {
	Value float64
}

type Generator interface {
	Generate(pixels []Pixel, f *frame.F, parameters GeneratorParameters) []GeneratorResult
}
