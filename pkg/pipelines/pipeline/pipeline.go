package pipeline

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/colorlookups"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/generators"
	"github.com/coral/nocube/pkg/mapping"
)

type Pipeline struct {
	Gen   pkg.Generator
	Color pkg.ColorLookup
}

func New(genName string, colorName string) *Pipeline {

	return &Pipeline{
		Gen:   generators.Generators[genName],
		Color: colorlookups.ColorLookups[colorName],
	}

}

func (p *Pipeline) Process(f *frame.F, m *mapping.Mapping) []pkg.ColorLookupResult {
	g := p.Gen.Generate(m.Coordinates, f, pkg.GeneratorParameters{})
	c := p.Color.Lookup(g, f, pkg.ColorLookupParameters{})

	return c
}
