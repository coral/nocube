package generators

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/generators/edgelord"
	"github.com/coral/nocube/pkg/generators/xd"
	"github.com/coral/nocube/pkg/generators/zebra"
)

var Generators = map[string]pkg.Generator{
	"xd":       &xd.Xd{},
	"edgelord": &edgelord.Edgelord{},
	"zebra":    &zebra.Zebra{},
}
