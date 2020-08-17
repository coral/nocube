package blending

import (
	"github.com/coral/nocube/pkg"
)

var BlendModes = map[string]bm{
	"add":      Add,
	"screen":   Screen,
	"subtract": Subtract,
}

type bm func([]pkg.Pixel, []pkg.Pixel, float64) []pkg.Pixel
