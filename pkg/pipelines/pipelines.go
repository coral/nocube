package pipelines

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/mapping"
	"github.com/coral/nocube/pkg/pipelines/blending"
	"github.com/coral/nocube/pkg/pipelines/pipeline"
)

type Pipelines struct {
	Active  []*pipeline.Pipeline
	frame   *frame.F
	mapping *mapping.Mapping
	d       *data.Data
}

func New(f *frame.F, m *mapping.Mapping, d *data.Data) *Pipelines {
	return &Pipelines{
		frame:   f,
		mapping: m,
		d:       d,
	}
}

func (p *Pipelines) Process(f *frame.F) []pkg.ColorLookupResult {
	outputBuffer := make([]pkg.ColorLookupResult, 864)
	for i := range outputBuffer {
		outputBuffer[i] = pkg.ColorLookupResult{}
	}
	for _, pipeline := range p.Active {
		data := pipeline.Process(f, p.mapping, p.d)
		outputBuffer = blending.BlendModes[pipeline.BlendMode](outputBuffer, data, 0.0)

	}
	return outputBuffer
}

func (p *Pipelines) Create(newPipeline *pipeline.Pipeline) {
	p.Active = append(p.Active, newPipeline)
}

func (p *Pipelines) GetActive() []*pipeline.Pipeline {
	return p.Active
}

func (p *Pipelines) Destroy(name string) {
	for i, e := range p.Active {
		if e.Name == name {
			p.Active = append(p.Active[:i], p.Active[i+1:]...)
		}
	}
}

func (p *Pipelines) ChangeOpacity(name string, opacity float64) {
	for _, e := range p.Active {
		if e.Name == name {
			e.ChangeOpacity(opacity)
		}
	}
}
