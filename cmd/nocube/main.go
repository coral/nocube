package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/colorlookups/dummy"
	"github.com/coral/nocube/pkg/generators/zebra"
	"github.com/coral/nocube/pkg/utils"
	"github.com/stojg/vector"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/devices/apa102"
	"periph.io/x/periph/host"
)

const NUM_LEDS = 604

var pixelCoordinates []pkg.Pixel

func init() {
	pixelCoordinates = make([]pkg.Pixel, NUM_LEDS)
}

func vectorLerp(a, b vector.Vector3, f float64) vector.Vector3 {
	l := b.NewSub(&a)

	asd2 := l.Scale(f)

	return *a.NewAdd(asd2)
}

func insertCoordinates(startIndex, stopIndex int, startVector, stopVector vector.Vector3) {
	length := stopIndex - startIndex

	for index := 0; index <= length; index++ {
		val := float64(index) / float64(length)

		pixelCoordinates[index+startIndex].Active = true
		pixelCoordinates[index+startIndex].Coordinate = vectorLerp(startVector, stopVector, val)
	}
}

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	insertCoordinates(9, 80, vector.Vector3{0, 1, 0}, vector.Vector3{0, 0, 0})
	insertCoordinates(87, 158, vector.Vector3{0, 0, 0}, vector.Vector3{1, 0, 0})
	insertCoordinates(165, 237, vector.Vector3{1, 0, 0}, vector.Vector3{1, 1, 0})
	insertCoordinates(246, 305, vector.Vector3{1, 1, 0}, vector.Vector3{1, 1, 1})
	insertCoordinates(313, 385, vector.Vector3{1, 1, 1}, vector.Vector3{0, 1, 1})
	insertCoordinates(393, 464, vector.Vector3{0, 1, 1}, vector.Vector3{0, 1, 0})
	insertCoordinates(472, 544, vector.Vector3{0, 1, 0}, vector.Vector3{1, 1, 0})

	d := getLEDs()

	const duration = 50 * time.Millisecond

	ticker := time.NewTicker(duration)
	stopTicker := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			stopTicker <- true
		}
	}()

	zebra := zebra.Zebra{}
	dummy := dummy.Dummy{}

	var t float64
	for {
		select {
		case <-ticker.C:
			res := zebra.Generate(pixelCoordinates, t, pkg.GeneratorParameters{})
			colorRes := dummy.Lookup(res, t, pkg.ColorLookupParameters{})

			var bytes = []byte{}
			for _, color := range colorRes {
				bytes = append(bytes, []byte{
					utils.Clamp255(color.Color[0] * 255),
					utils.Clamp255(color.Color[1] * 255),
					utils.Clamp255(color.Color[2] * 255),
				}...)
			}

			// img := image.NewNRGBA(d.Bounds())
			// for x := 0; x < img.Rect.Max.X; x++ {
			// 	img.SetNRGBA(x, 0, colorWheel(float64(x)+offset/float64(img.Rect.Max.X)+offset))
			// }
			// if err := d.Draw(d.Bounds(), img, image.Point{}); err != nil {
			// 	log.Fatal("Error drawing:", err)
			// }

			_, err := d.Write(bytes)
			if err != nil {
				log.Fatal("xd", err)
			}

			t += 0.05

		case <-stopTicker:
			ticker.Stop()
			d.Halt()
			return
		}
	}
}

// getLEDs returns an *apa102.Dev, or fails back to *screen.Dev if no SPI port
// is found.
func getLEDs() *apa102.Dev {
	s, err := spireg.Open("")
	if err != nil {
		fmt.Printf("Failed to find a SPI port, printing at the console:\n")
		return nil
	}

	// Change the option values to see their effects.
	opts := apa102.DefaultOpts
	opts.NumPixels = NUM_LEDS
	opts.Intensity = 255
	d, err := apa102.New(s, &opts)
	if err != nil {
		log.Fatal(err)

	}
	return d
}

// colorWheel returns a HSV color wheel.
func colorWheel(h float64) color.NRGBA {
	h *= 6
	switch {
	case h < 1.:
		return color.NRGBA{R: 255, G: byte(255 * h), A: 255}
	case h < 2.:
		return color.NRGBA{R: byte(255 * (2 - h)), G: 255, A: 255}
	case h < 3.:
		return color.NRGBA{G: 255, B: byte(255 * (h - 2)), A: 255}
	case h < 4.:
		return color.NRGBA{G: byte(255 * (4 - h)), B: 255, A: 255}
	case h < 5.:
		return color.NRGBA{R: byte(255 * (h - 4)), B: 255, A: 255}
	default:
		return color.NRGBA{R: 255, B: byte(255 * (6 - h)), A: 255}

	}

}
