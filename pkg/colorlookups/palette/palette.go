package palette

import (
	"image"
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/color"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/utils"
	"github.com/stojg/vector"
)

type Palette struct {
}

var _ pkg.ColorLookup = &Palette{}
var palette image.Image
var x int
var y int

func init() {
	var err error
	palette, err = color.LoadPaletteFromImage("aetgrot")
	if err != nil {
		panic(err)
	}
	size := palette.Bounds().Size()
	x = size.X
	y = size.Y
}

func (g *Palette) Lookup(generatorResults []pkg.GeneratorResult, f *frame.F, parameters pkg.ColorLookupParameters) (results []pkg.ColorLookupResult) {
	for _, pixel := range generatorResults {
		p := math.Floor(utils.Clamp01(math.Abs(pixel.Phase)) * float64(x))
		rr, gg, bb, _ := palette.At(int(p), 0).RGBA()
		r := float64(rr) / 255
		g := float64(gg) / 255
		b := float64(bb) / 255
		// clampedValue := clamp01(pixel.Value)
		results = append(results, pkg.ColorLookupResult{
			Color: vector.Vector3{
				r,
				g,
				b,
			},
		})
	}

	return
}
