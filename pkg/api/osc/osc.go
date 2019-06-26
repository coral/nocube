package osc

import (
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/pipelines"
	"github.com/hypebeast/go-osc/osc"
)

type OSC struct {
	pipelines *pipelines.Pipelines
	s         *osc.Server
	d         *data.Data
}

func New(p *pipelines.Pipelines, d *data.Data) OSC {
	return OSC{
		pipelines: p,
		d:         d,
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
	o.s.Handle("/1/push2", func(msg *osc.Message) {
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

	o.s.Handle("/1/rotary1", func(msg *osc.Message) {
		o.d.SetScopedFloat64("denis", "zebra", "speed", AssertFloat64(msg))
	})

	o.s.Handle("/1/rotary2", func(msg *osc.Message) {
		o.d.SetScopedFloat64("denis", "colorize", "hue", AssertFloat64(msg))
	})

	o.s.Handle("/1/rotary3", func(msg *osc.Message) {
		o.d.SetScopedInt64("olof", "strobe", "segment", int64(AssertFloat64(msg)))
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
