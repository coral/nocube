package layers

import (
	"github.com/coral/nocube/generators"
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/dynamic"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/layers/blending"
	"github.com/coral/nocube/pkg/mapping"
	"golang.org/x/exp/errors/fmt"
)

type Layers struct {
	Active         []Chain
	frame          *frame.F
	mapping        *mapping.Mapping
	d              *data.Data
	dynamic        *dynamic.Dynamic
	AvaliableSteps map[string]pkg.Step

	outputBuffer []pkg.Pixel
}

type Chain struct {
	Name         string
	BlendMode    string
	Opacity      float64
	Sequence     []Link
	outputBuffer []pkg.Pixel
}

func New(project string, f *frame.F, m *mapping.Mapping, d *data.Data, dynamic *dynamic.Dynamic) Layers {
	return Layers{
		frame:          f,
		mapping:        m,
		d:              d,
		dynamic:        dynamic,
		AvaliableSteps: make(map[string]pkg.Step),

		outputBuffer: make([]pkg.Pixel, len(m.Coordinates)),
	}
}

func (c *Layers) Initialize() {
	var tempStep []pkg.Step
	tempStep = append(tempStep, generators.CreateSteps()...)
	tempStep = append(tempStep, c.dynamic.CreateSteps()...)

	for _, step := range tempStep {
		c.AvaliableSteps[step.Name()] = step
	}

	for _, m := range c.AvaliableSteps {
		fmt.Println(m.Name())
	}

	ch := Chain{
		Name:      "test",
		BlendMode: "add",
		Opacity:   1.0,
		Sequence: []Link{Link{
			Step:      c.AvaliableSteps["p1.js"],
			BlendMode: "add",
			Opacity:   1.0,
		}},
		outputBuffer: make([]pkg.Pixel, len(c.mapping.Coordinates)),
	}

	c.Active = append(c.Active, ch)
}

func (c *Layers) Process(f *frame.F) []pkg.Pixel {

	///TODO: MULTITHREAD EACH LAYER

	for i := range c.outputBuffer {
		c.outputBuffer[i] = pkg.Pixel{}
	}

	for _, chain := range c.Active {

		for i := range chain.outputBuffer {
			chain.outputBuffer[i] = pkg.Pixel{}
		}

		for _, link := range chain.Sequence {
			data := link.Gen(f, c.mapping, c.d)
			blending.Opacity(data, link.Opacity)
			chain.outputBuffer = blending.BlendModes[link.BlendMode](chain.outputBuffer, data, 0.0)
		}
		blending.Opacity(chain.outputBuffer, chain.Opacity)
		c.outputBuffer = blending.BlendModes[chain.BlendMode](c.outputBuffer, chain.outputBuffer, 0.0)
	}

	return c.outputBuffer
}

//////////////
/// Link
/////////////

type Link struct {
	Step      pkg.Step
	BlendMode string
	Opacity   float64
}

func (l *Link) Gen(f *frame.F, m *mapping.Mapping, d *data.Data) []pkg.Pixel {
	c := l.Step.Gen(m.Coordinates, f)
	return c
}
