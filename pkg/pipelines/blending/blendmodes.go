package blending

import (
	"github.com/coral/nocube/pkg"
)

var BlendModes = map[string]bm{
	"add":      Add,
	"screen":   Screen,
	"subtract": Subtract,
}

type bm func([]pkg.ColorLookupResult, []pkg.ColorLookupResult, float64) []pkg.ColorLookupResult
