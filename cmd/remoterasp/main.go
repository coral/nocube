package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"periph.io/x/periph/devices/apa102"

	"github.com/gorilla/websocket"
	"github.com/grandcat/zeroconf"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

var port = flag.Int("port", 12500, "listen port")
var bridgename = flag.String("bridgename", "first", "name of bridge for discovery")

var upgrader = websocket.Upgrader{} // use default options

var dataline1 = apa102.Dev{}

var dc1 chan []byte

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

		dc1 <- message

	}
}

func main() {

	hsname, _ := os.Hostname()
	flag.Parse()
	log.SetFlags(0)

	dc1 = make(chan []byte)

	server, err := zeroconf.Register(hsname, "_apabridge._tcp", "local.", *port, []string{"txtv=0", "lo=1", "la=2"}, nil)
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

	go func() {
		for {
			select {
			case d := <-dc1:
				dataline1.Write(d)
			}
		}
	}()

	http.HandleFunc("/data", data)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*port), nil))
}
