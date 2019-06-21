package pipelines

import (
	"fmt"

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
	outputBuffer := make([]pkg.ColorLookupResult, 864)
	for i := range outputBuffer {
		outputBuffer[i] = pkg.ColorLookupResult{}
	}

	for _, pipeline := range p.Active {
		data := pipeline.Process(p.frame, p.mapping)
		fmt.Println(data)
		outputBuffer = BlendModes[pipeline.BlendMode](outputBuffer, data, 0.0)

	}
	return &outputBuffer
}

func (p *Pipelines) Add(newPipeline *pipeline.Pipeline) {
	p.Active[newPipeline.Name] = newPipeline
}
