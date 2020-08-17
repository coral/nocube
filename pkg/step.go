package pkg

import "github.com/coral/nocube/pkg/frame"

type Step interface {
	Init()
	Name() string
	Type() string
	Gen(pixels []Pixel, f *frame.F) []Pixel
}

type StepIndex interface {
	IndexStep() []string
}
