// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"

	"periph.io/x/periph/devices/apa102"

	"github.com/gorilla/websocket"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var ba = apa102.Dev{}

func echo(w http.ResponseWriter, r *http.Request) {
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
	opts.Intensity = 30
	a, err := apa102.New(s, &opts)
	defer a.Halt()

	if err != nil {
		panic(err)
	}

	ba = *a

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
