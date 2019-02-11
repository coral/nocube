package apaled

import (
	"errors"
	"fmt"
	"log"
	"os"

	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/devices/apa102"
	"periph.io/x/periph/host"
)

type ApaLED struct {
	NUM_LEDS  int64
	INTENSITY int64
}

func (a *ApaLED) Init(numleds int64, intensity int64) {

	a.NUM_LEDS = numleds
	a.INTENSITY = intensity

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	d, err := a.getApa102Device()
	if err != nil {
		fmt.Println("Error getting apa102 device:", err)
		os.Exit(1)
	}
	defer d.Halt()
}

func (a *ApaLED) pusher(d *apa102.Dev, in chan []byte, stop chan bool, debug chan bool) {
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
func (a *ApaLED) getApa102Device() (*apa102.Dev, error) {
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
