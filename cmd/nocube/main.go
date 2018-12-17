package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/colorlookups/dummy"
	"github.com/coral/nocube/pkg/colorlookups/palette"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/generators/edgelord"
	"github.com/coral/nocube/pkg/generators/xd"
	"github.com/coral/nocube/pkg/generators/zebra"
	"github.com/coral/nocube/pkg/mapping"
	"github.com/coral/nocube/pkg/utils"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/devices/apa102"
	"periph.io/x/periph/host"
)

const NUM_LEDS = 920
const INTENSITY = 30
const PUSHER_DURATION = 5 * time.Millisecond

var packagesPushedThisSecond uint

func pusher(d *apa102.Dev, in chan []byte, stop chan bool, debug chan bool) {
	for {
		select {
		case bytes := <-in:
			_, err := d.Write(bytes)
			if err != nil {
				log.Fatal("xd", err)
			}

			debug <- true

		case <-stop:
			return
		}
	}
}

func debugger(packagePushed, stop chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-packagePushed:
			packagesPushedThisSecond++
		case <-ticker.C:
			fmt.Println("Packages pushed this second:", packagesPushedThisSecond)
			packagesPushedThisSecond = 0

		case <-stop:
			return
		}
	}
}

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	mapping := mapping.New("v1", NUM_LEDS)
	err := mapping.LoadFile()
	if err != nil {
		panic(err)
	}

	d, err := getApa102Device()
	if err != nil {
		fmt.Println("Error getting apa102 device:", err)
		os.Exit(1)
	}
	defer d.Halt()

	generatorStop := make(chan bool)
	pusherStop := make(chan bool)
	debuggerStop := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			generatorStop <- true
			pusherStop <- true
			debuggerStop <- true
		}
	}()

	bytesChannel := make(chan []byte, 1)
	debuggerChannel := make(chan bool, 10)

	go pusher(d, bytesChannel, pusherStop, debuggerChannel)

	go debugger(debuggerChannel, debuggerStop)

	generator(mapping, bytesChannel, generatorStop)
}

func generator(mapping *mapping.Mapping, bytesChannel chan []byte, stop chan bool) {
	ticker := time.NewTicker(PUSHER_DURATION)
	defer ticker.Stop()

	generators := map[string]pkg.Generator{
		"xd":       &xd.Xd{},
		"edgelord": &edgelord.Edgelord{},
		"zebra":    &zebra.Zebra{},
	}

	colorLookups := map[string]pkg.ColorLookup{
		"dummy":   &dummy.Dummy{},
		"palette": &palette.Palette{},
	}

	var t float64
	frame := frame.New()
	frame.SetBeat(60.0/30.0, 0)

	generator := generators["edgelord"]
	colorLookup := colorLookups["palette"]

	if generator == nil || colorLookup == nil {
		panic("No valid generator or color lookup chosen")
	}

	for {
		select {
		case <-ticker.C:
			frame.Update(t)
			res := generator.Generate(mapping.Coordinates, &frame, pkg.GeneratorParameters{})
			colorRes := colorLookup.Lookup(res, &frame, pkg.ColorLookupParameters{})

			var bytes = []byte{}
			for _, color := range colorRes {
				bytes = append(bytes, []byte{
					utils.Clamp255(color.Color[0] * 255),
					utils.Clamp255(color.Color[1] * 255),
					utils.Clamp255(color.Color[2] * 255),
				}...)
			}

			bytesChannel <- bytes

			t += PUSHER_DURATION.Seconds()

		case <-stop:
			return
		}
	}
}

func getApa102Device() (*apa102.Dev, error) {
	s, err := spireg.Open("")
	if err != nil {
		return nil, errors.New("Unable to find SPI port")
	}

	// Change the option values to see their effects.
	opts := apa102.DefaultOpts
	opts.NumPixels = NUM_LEDS
	opts.Intensity = INTENSITY
	return apa102.New(s, &opts)
}
