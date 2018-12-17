package pkg

import "github.com/stojg/vector"

type Pixel struct {
	Active     bool
	Coordinate vector.Vector3
}

type GeneratorParameters struct {
	Speed float64
}

type GeneratorResult struct {
	Value float64
}

type Generator interface {
	Generate(pixels []Pixel, t float64, parameters GeneratorParameters) []GeneratorResult
}
