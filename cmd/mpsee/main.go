package main

import (
	"fmt"
	"log"

	"periph.io/x/extra/hostextra/d2xx"

	"github.com/coral/nocube/pkg/utils"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/devices/apa102"
	"periph.io/x/periph/host"
)

//https://github.com/periph/extra/blob/master/hostextra/d2xx/d2xxsmoketest/d2xxsmoketest.go
//https://periph.io/device/ftdi/

//https://godoc.org/periph.io/x/extra/hostextra/d2xx#Allhttps://godoc.org/periph.io/x/extra/hostextra/d2xx#All

//https://learn.adafruit.com/adafruit-ft232h-breakout/mac-osx-setup
//https://cpldcpu.wordpress.com/2014/08/27/apa102/

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d2xx.Version())

	ports := spireg.All()

	s, err := spireg.Open(ports[0].Name)
	if err != nil {
		panic(err)
	}

	// Change the option values to see their effects.
	opts := apa102.DefaultOpts
	opts.NumPixels = 32
	opts.Intensity = 30
	a, err := apa102.New(s, &opts)
	defer a.Halt()

	if err != nil {
		panic(err)
	}

	var bytes = []byte{}
	for i := 0; i < 32; i++ {
		bytes = append(bytes, []byte{
			utils.Clamp255(127),
			utils.Clamp255(127),
			utils.Clamp255(127),
		}...)
	}

	for {
		a.Write(bytes)
	}

}

/* func main() {
	if _, err := host.Init(); err != nil {
		panic(err)
	}
	all := d2xx.All()
	fmt.Println(all)

	major, minor, build := d2xx.Version()
	log.Printf("Using library %d.%d.%d\n", major, minor, build)

	if len(all) == 0 {
		panic("no spi device")
	}

	dev := all[0]
	d := dev.(*d2xx.FT232H)

	d.SetSpeed(4 * physic.MegaHertz)

	fmt.Printf("  SPI functionality:\n")
	dd, err := d.SPI()
	if err != nil {
		panic(err)
	}

	s, err := dd.Connect(4*physic.MegaHertz, 0, 8)

	opts := apa102.DefaultOpts
	opts.NumPixels = 32
	opts.Intensity = 100
	r, err := apa102.New(s, &opts)

	if err != nil {
		panic(err)
	}

	var bytes = []byte{}
	bytes = append(bytes, []byte{
		utils.Clamp255(255),
		utils.Clamp255(255),
		utils.Clamp255(255),
	}...)

	fmt.Println(r.Write(bytes))

	if err = s.Close(); err != nil {
		panic(err)
	}
}
*/
