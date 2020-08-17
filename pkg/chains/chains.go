package chains

import (
	"github.com/coral/nocube/generators"
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/dynamic"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/mapping"
	"golang.org/x/exp/errors/fmt"
)

type Chains struct {
	Active         []Chain
	frame          *frame.F
	mapping        *mapping.Mapping
	d              *data.Data
	dynamic        *dynamic.Dynamic
	AvaliableSteps []pkg.Step
}

type Chain struct {
	Sequence []Link
}

type Link struct {
	Step      pkg.Step
	BlendMode string
}

func New(project string, f *frame.F, m *mapping.Mapping, d *data.Data, dynamic *dynamic.Dynamic) Chains {
	return Chains{
		frame:   f,
		mapping: m,
		d:       d,
		dynamic: dynamic,
	}
}

func (c *Chains) Initialize() {
	c.AvaliableSteps = append(c.AvaliableSteps, generators.CreateSteps()...)
	c.AvaliableSteps = append(c.AvaliableSteps, c.dynamic.CreateSteps()...)

	fmt.Println(c.AvaliableSteps)
}
