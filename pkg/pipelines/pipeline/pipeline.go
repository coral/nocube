package pipeline

import (
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

func New(name string, opacity float64, genName string, colorName string, blendMode string) *Pipeline {

	return &Pipeline{
		Name:      name,
		Opacity:   opacity,
		Gen:       generators.Generators[genName],
		Color:     colorlookups.ColorLookups[colorName],
		BlendMode: blendMode,
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
