package pkg

import (
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type ColorLookupParameters struct {
	Data *data.Data
	Name string
}

type ColorLookupResult struct {
	Color vector.Vector3 `json:"C"`
}

type ColorLookup interface {
	Init()
	Lookup(generatorResults []GeneratorResult, f *frame.F, parameters ColorLookupParameters) []ColorLookupResult
	Name() string
}
