package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/coral/nocube/pkg/utils"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "10.0.1.67:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/data"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()

	var p float64 = 0.0
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			_ = t
			var bytes = []byte{}
			for i := 0; i < 32; i++ {
				bytes = append(bytes, []byte{
					utils.Clamp255(p),
					utils.Clamp255(0),
					utils.Clamp255(0),
				}...)
			}
			if p <= 255 {
				p = p + 1.0
			} else {
				p = 0.0
			}
			fmt.Println(p)

			err := c.WriteMessage(websocket.BinaryMessage, bytes)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
