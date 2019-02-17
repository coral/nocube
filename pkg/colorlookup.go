package pkg

import (
	"github.com/coral/nocube/pkg/frame"
	"github.com/stojg/vector"
)

type ColorLookupParameters struct {
	Speed float64
}

type ColorLookupResult struct {
	Color vector.Vector3 `json:"C"`
}

type ColorLookup interface {
	Lookup(generatorResults []GeneratorResult, f *frame.F, parameters ColorLookupParameters) []ColorLookupResult
}
