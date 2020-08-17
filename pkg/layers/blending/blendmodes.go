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

func Opacity(c []pkg.Pixel, opacity float64) []pkg.Pixel {
	for i, d := range c {
		c[i].Color = *d.Color.Scale(opacity)
	}

	return c
}
