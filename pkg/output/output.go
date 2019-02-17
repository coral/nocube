package output

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/output/rapa102"
	"github.com/coral/nocube/pkg/output/ws"
	"github.com/coral/nocube/pkg/settings"
)

type Output interface {
	Init()
	ModuleName() string
	Write([]pkg.ColorLookupResult)

	SetTargetFrameRate(int)
	GetTargetFrameRate() int
}

type Controller struct {
	ActivatedOutputs []string
	PixelStream      chan []pkg.ColorLookupResult

	o []Output
	s *settings.Settings
}

func New(s *settings.Settings) *Controller {

	availableModules := []Output{rapa102.New(), ws.New()}
	var loadedModules []Output
	for _, m := range availableModules {
		for _, cm := range s.Global.Output.ActivatedOutputs {
			if m.ModuleName() == cm {
				loadedModules = append(loadedModules, m)
			}
		}
	}

	return &Controller{
		ActivatedOutputs: s.Global.Output.ActivatedOutputs,
		o:                loadedModules,
		s:                s,
	}
}

func (l *Controller) Init() {
	for _, o := range l.o {
		go o.Init()
	}
}

//TODO make it actually resolve mapping here
func (l *Controller) Write(d []pkg.ColorLookupResult) {
	for _, ao := range l.o {
		ao.Write(d)
	}
}
