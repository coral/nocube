package generators

import (
	"github.com/coral/nocube/generators/fill"
	"github.com/coral/nocube/pkg"
)

var Generators = map[string]pkg.Step{
	"fill": &fill.Fill{},
}

func CreateSteps() []pkg.Step {
	var steps []pkg.Step
	for _, m := range Generators {
		steps = append(steps, m)
	}
	return steps
}
