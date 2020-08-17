package pipelines

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/mapping"
	"github.com/coral/nocube/pkg/pipelines/blending"
	"github.com/coral/nocube/pkg/pipelines/pipeline"
	log "github.com/sirupsen/logrus"
)

type Pipelines struct {
	Name    string
	Active  []*pipeline.Pipeline
	frame   *frame.F
	mapping *mapping.Mapping
	d       *data.Data
}

func New(name string, f *frame.F, m *mapping.Mapping, d *data.Data) *Pipelines {

	return &Pipelines{
		Name:    name,
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
		if pipeline.Opacity > 0.0 {
			data := pipeline.Process(f, p.mapping, p.d)
			outputBuffer = blending.BlendModes[pipeline.BlendMode](outputBuffer, data, 0.0)
		}
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

func (p *Pipelines) LoadPipelines() error {
	b, err := ioutil.ReadFile("../../files/settings/" + p.Name + ".json")
	if err != nil {
		return err
	}

	var lp []pipeline.PipelineJSON
	err = json.Unmarshal(b, &lp)
	if err != nil {
		return err
	}

	for _, np := range lp {
		pipe, err := pipeline.New(np.Name, np.GeneratorName, np.ColorLookupName, np.BlendMode)
		if err != nil {
			log.Error("Could not load pipeline")
			log.Error(err)
		} else {
			pipe.ChangeOpacity(np.Opacity)
			p.Create(pipe)
			log.WithFields(log.Fields{
				"Name":         np.Name,
				"Generator":    np.GeneratorName,
				"Color Lookup": np.ColorLookupName,
				"Opacity":      np.Opacity,
				"Blend Mode":   np.BlendMode,
			}).Info("Created pipeline")
		}
	}

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for _ = range ticker.C {
			p.SavePipelines()
		}
	}()

	return nil
}

func (p *Pipelines) SavePipelines() {
	var s []pipeline.PipelineJSON
	for _, pipe := range p.Active {
		s = append(s, pipe.Marshal())
	}
	data, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		log.Error("could not save pipeline settings")
		log.Error(err)
		panic(err)
	}

	err = ioutil.WriteFile("../../files/settings/"+p.Name+".json", data, 0644)
	if err != nil {
		log.Error(err)
	}

	log.Info("Saved pipelines to " + p.Name + ".json")
}
