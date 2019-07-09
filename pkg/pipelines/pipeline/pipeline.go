package pipeline

import (
	"fmt"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/colorlookups"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/generators"
	"github.com/coral/nocube/pkg/mapping"
)

type Pipeline struct {
	Name      string
	Opacity   float64
	Gen       pkg.Generator
	Color     pkg.ColorLookup
	BlendMode string
}

type PipelineJSON struct {
	Name            string
	Opacity         float64
	GeneratorName   string
	ColorLookupName string
	BlendMode       string
}

func New(name string, genName string, colorName string, blendMode string) (*Pipeline, error) {
	g := generators.Generators[genName]
	if g == nil {
		return nil, fmt.Errorf("Could not find generator " + genName)
	}

	c := colorlookups.ColorLookups[colorName]
	if c == nil {
		return nil, fmt.Errorf("Could not find color loookup " + colorName)
	}
	c.Init()

	return &Pipeline{
		Name:      name,
		Opacity:   0.0,
		Gen:       g,
		Color:     c,
		BlendMode: blendMode,
	}, nil

}

func (p *Pipeline) Marshal() PipelineJSON {
	return PipelineJSON{
		Name:            p.Name,
		Opacity:         p.Opacity,
		GeneratorName:   p.Gen.Name(),
		ColorLookupName: p.Color.Name(),
		BlendMode:       p.BlendMode,
	}
}

func (p *Pipeline) Process(f *frame.F, m *mapping.Mapping, d *data.Data) []pkg.ColorLookupResult {
	g := p.Gen.Generate(m.Coordinates, f, pkg.GeneratorParameters{
		Data: d,
		Name: p.Name,
	})
	c := p.Color.Lookup(g, f, pkg.ColorLookupParameters{
		Data: d,
		Name: p.Name,
	})
	for i, d := range c {
		c[i].Color = *d.Color.Scale(p.Opacity)
	}
	return c
}

func (p *Pipeline) ChangeOpacity(newOpacity float64) {
	p.Opacity = newOpacity
}
