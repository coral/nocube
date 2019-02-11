package pipelines

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/mapping"
	"github.com/coral/nocube/pkg/pipelines/pipeline"
)

type Pipelines struct {
	Active  map[string]*pipeline.Pipeline
	frame   *frame.F
	mapping *mapping.Mapping
}

func New(f *frame.F, m *mapping.Mapping) *Pipelines {
	return &Pipelines{
		Active:  make(map[string]*pipeline.Pipeline, 100),
		frame:   f,
		mapping: m,
	}
}

func (p *Pipelines) Process(f *frame.F) *[]pkg.ColorLookupResult {
	//var outputBuffer []pkg.ColorLookupResult

	p.frame = f
	for _, pipeline := range p.Active {
		data := pipeline.Process(p.frame, p.mapping)
		return &data
		//Intensity
		for i, d := range data {
			data[i].Color = *d.Color.Scale(0.5)
		}

	}
	return nil
}

func (p *Pipelines) Add(newPipeline *pipeline.Pipeline) {
	p.Active["first"] = newPipeline
}
