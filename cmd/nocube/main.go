package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"time"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/colorlookups/dummy"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/generators/zebra"
	"github.com/coral/nocube/pkg/utils"
	"github.com/stojg/vector"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/devices/apa102"
	"periph.io/x/periph/host"
)

const NUM_LEDS = 920
const INTENSITY = 30
const PUSHER_DURATION = 5 * time.Millisecond

var pixelCoordinates []pkg.Pixel
var packagesPushedThisSecond uint

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

	dir := stopVector.NewSub(&startVector)
	dir[0] = math.Abs(dir[0])
	dir[1] = math.Abs(dir[1])
	dir[2] = math.Abs(dir[2])

	dir.Normalize()

	for index := 0; index <= length; index++ {
		val := float64(index) / float64(length)

		pixelCoordinates[index+startIndex].Active = true
		pixelCoordinates[index+startIndex].Coordinate = vectorLerp(startVector, stopVector, val)
		pixelCoordinates[index+startIndex].Normal = *dir
	}
}

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
			packagesPushedThisSecond = 0
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

	insertCoordinates(9, 80, vector.Vector3{0, 0.95, 0}, vector.Vector3{0, 0.05, 0})
	insertCoordinates(87, 158, vector.Vector3{0.05, 0, 0}, vector.Vector3{0.95, 0, 0})
	insertCoordinates(165, 237, vector.Vector3{1, 0.05, 0}, vector.Vector3{1, 0.95, 0})
	insertCoordinates(246, 305, vector.Vector3{1, 1, 0.05}, vector.Vector3{1, 1, 0.95})
	insertCoordinates(313, 385, vector.Vector3{0.95, 1, 1}, vector.Vector3{0.05, 1, 1})
	insertCoordinates(393, 464, vector.Vector3{0, 1, 0.95}, vector.Vector3{0, 1, 0.05})
	insertCoordinates(472, 544, vector.Vector3{0.05, 1, 0}, vector.Vector3{0.95, 1, 0})

	insertCoordinates(623, 694, vector.Vector3{1, 0.95, 1}, vector.Vector3{1, 0.05, 1})

	insertCoordinates(701, 772, vector.Vector3{0.95, 0, 1}, vector.Vector3{0.05, 0, 1})

	insertCoordinates(779, 850, vector.Vector3{0, 0, 0.95}, vector.Vector3{0, 0, 0.05})

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

	generator(generatorStop)
}

func generator(stop chan bool) {
	ticker := time.NewTicker(PUSHER_DURATION)
	defer ticker.Stop()

	// zebra := xd.Xd{}
	zebra := zebra.Zebra{}
	dummy := dummy.Dummy{}

	var t float64
	frame := frame.New()
	frame.SetBeat(60.0/30.0, 0)

	for {
		select {
		case <-ticker.C:
			frame.Update(t)
			res := zebra.Generate(pixelCoordinates, &frame, pkg.GeneratorParameters{})
			colorRes := dummy.Lookup(res, &frame, pkg.ColorLookupParameters{})

			var bytes = []byte{}
			for _, color := range colorRes {
				bytes = append(bytes, []byte{
					utils.Clamp255(color.Color[0] * 255),
					utils.Clamp255(color.Color[1] * 255),
					utils.Clamp255(color.Color[2] * 255),
				}...)
			}

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
