package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"periph.io/x/periph/devices/apa102"

	"github.com/gorilla/websocket"
	"github.com/grandcat/zeroconf"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

var port = flag.Int("port", 12500, "listen port")
var bridgename = flag.String("bridgename", "korv", "name of bridge for discovery")

var upgrader = websocket.Upgrader{} // use default options

var dataline1 = apa102.Dev{}
var dataline2 = apa102.Dev{}

var dc1 chan []byte
var dc2 chan []byte

func data(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	var message []byte
	var erro error
	for {
		_, message, erro = c.ReadMessage()
		if erro != nil {
			log.Println("read:", erro)
			erro = nil
			break
		}
		if message[len(message)-1] == 0x00 {
			d := message[:len(message)-1]
			dc1 <- d
		} else {
			d := message[:len(message)-1]
			dc2 <- d
		}
	}
}

func main() {

	flag.Parse()
	log.SetFlags(0)

	dc1 = make(chan []byte)
	dc2 = make(chan []byte)

	server, err := zeroconf.Register(*bridgename, "_apabridge._tcp", "local.", *port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	ports := spireg.All()
	for _, element := range ports {
		fmt.Println(element.Name)
	}
	//dataline 1
	s1, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		panic(err)
	}

	opts := apa102.DefaultOpts
	opts.NumPixels = 432
	opts.Intensity = 255
	a, err := apa102.New(s1, &opts)
	defer a.Halt()

	if err != nil {
		panic(err)
	}

	dataline1 = *a

	//dataline 2
	s2, err := spireg.Open("/dev/spidev1.0")
	if err != nil {
		panic(err)
	}

	b, err := apa102.New(s2, &opts)
	defer b.Halt()

	if err != nil {
		panic(err)
	}

	dataline2 = *b

	go func() {
		for {
			select {
			case d := <-dc1:
				dataline1.Write(d)
			}
		}
	}()
	go func() {
		for {
			select {
			case d := <-dc2:
				dataline2.Write(d)
			}
		}
	}()

	http.HandleFunc("/data", data)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*port), nil))
}
