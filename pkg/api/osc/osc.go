package osc

import (
	"encoding/json"
	"io/ioutil"

	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/pipelines"
	"github.com/hypebeast/go-osc/osc"
)

type OSC struct {
	pipelines *pipelines.Pipelines
	s         *osc.Server
	d         *data.Data
}

type OSCMap struct {
	Path        string `json:"p"`
	Type        string `json:"t"`
	Destination string `json:"d"`
	Function    string `json:"f"`
	Parameter   string `json:"e"`
}

func New(p *pipelines.Pipelines, d *data.Data) OSC {
	return OSC{
		pipelines: p,
		d:         d,
	}
}

func (o *OSC) Init(addr string) error {
	b, err := ioutil.ReadFile("../../files/settings/oscmap.json")
	if err != nil {
		return err
	}

	var oscmapping []OSCMap
	err = json.Unmarshal(b, &oscmapping)
	if err != nil {
		return err
	}
	o.s = &osc.Server{Addr: addr}

	for _, d := range oscmapping {
		m := d
		o.s.Handle(m.Path, func(msg *osc.Message) {

			if m.Function == "opacity" {
				o.pipelines.ChangeOpacity(m.Destination, AssertFloat64(msg))
				return
			}

			if m.Type == "f" {
				o.d.SetScopedFloat64(
					m.Destination,
					m.Function,
					m.Parameter,
					AssertFloat64(msg),
				)
			}

			if m.Type == "i" {
				o.d.SetScopedInt64(
					m.Destination,
					m.Function,
					m.Parameter,
					int64(AssertFloat64(msg)),
				)
			}
		})
	}

	o.s.ListenAndServe()

	return nil

}

func AssertFloat64(msg *osc.Message) float64 {
	//var m = d[0]
	s, _ := msg.TypeTags()
	if s != ",f" {
		return 0.0
	}
	for _, s := range msg.Arguments {
		switch i := s.(type) {
		case float64:
			return i
		case float32:
			return float64(i)
		case int64:
			return float64(i)
		}
	}

	return 0.0
}
