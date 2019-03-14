package colorlookups

import (
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/colorlookups/allwhite"
	"github.com/coral/nocube/pkg/colorlookups/dummy"
	"github.com/coral/nocube/pkg/colorlookups/fft"
	"github.com/coral/nocube/pkg/colorlookups/palette"
)

var ColorLookups = map[string]pkg.ColorLookup{
	"dummy":    &dummy.Dummy{},
	"palette":  &palette.Palette{},
	"allwhite": &allwhite.AllWhite{},
	"fft":      &fft.FFT{},
}
