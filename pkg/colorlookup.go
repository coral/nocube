package pkg

import "github.com/stojg/vector"

type ColorLookupParameters struct {
	Speed float64
}

type ColorLookupResult struct {
	Color vector.Vector3
}

type ColorLookup interface {
	Lookup(generatorResults []GeneratorResult, t float64, parameters ColorLookupParameters) []ColorLookupResult
}
