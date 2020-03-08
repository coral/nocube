package generators

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/generators/beat"
	"github.com/coral/nocube/pkg/generators/beatstrobe"
	"github.com/coral/nocube/pkg/generators/edgelord"
	"github.com/coral/nocube/pkg/generators/runner"
	"github.com/coral/nocube/pkg/generators/sense"
	"github.com/coral/nocube/pkg/generators/solid"
	"github.com/coral/nocube/pkg/generators/strobe"
	"github.com/coral/nocube/pkg/generators/xd"
	"github.com/coral/nocube/pkg/generators/zebra"
)

var Generators = map[string]pkg.Generator{
	"xd":         &xd.Xd{},
	"edgelord":   &edgelord.Edgelord{},
	"zebra":      &zebra.Zebra{},
	"strobe":     &strobe.Strobe{},
	"beat":       &beat.Beat{},
	"solid":      &solid.Solid{},
	"beatstrobe": &beatstrobe.BeatStrobe{},
	"sense":      &sense.Sense{},
	"runner":     &runner.Runner{},
}
