package main

import (
	"flag"
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

var ba = apa102.Dev{}

func data(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		ba.Write(message)
	}
}

func main() {

	flag.Parse()
	log.SetFlags(0)

	server, err := zeroconf.Register(*bridgename, "_apabridge._tcp", "local.", *port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	ports := spireg.All()
	s, err := spireg.Open(ports[0].Name)
	if err != nil {
		panic(err)
	}

	opts := apa102.DefaultOpts
	opts.NumPixels = 32
	opts.Intensity = 255
	a, err := apa102.New(s, &opts)
	defer a.Halt()

	if err != nil {
		panic(err)
	}

	ba = *a

	http.HandleFunc("/data", data)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*port), nil))
}
