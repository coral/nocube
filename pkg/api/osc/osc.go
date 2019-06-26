package osc

import (
	"github.com/coral/nocube/pkg/pipelines"
	"github.com/hypebeast/go-osc/osc"
)

type OSC struct {
	pipelines *pipelines.Pipelines
	s         *osc.Server
}

func New(p *pipelines.Pipelines) OSC {
	return OSC{
		pipelines: p,
	}
}

func (o *OSC) Init(addr string) {
	o.s = &osc.Server{Addr: addr}

	o.s.Handle("/1/fader1", func(msg *osc.Message) {
		o.pipelines.ChangeOpacity("denis", AssertFloat64(msg))
	})
	o.s.Handle("/1/fader2", func(msg *osc.Message) {
		o.pipelines.ChangeOpacity("olof", AssertFloat64(msg))
	})
	o.s.Handle("/1/fader3", func(msg *osc.Message) {
		o.pipelines.ChangeOpacity("solid", AssertFloat64(msg))
	})
	o.s.Handle("/1/fader4", func(msg *osc.Message) {
		o.pipelines.ChangeOpacity("beatstrobe", AssertFloat64(msg))
	})
	o.s.Handle("/1/fader5", func(msg *osc.Message) {
		o.pipelines.ChangeOpacity("xd", AssertFloat64(msg))
	})

	o.s.ListenAndServe()

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
