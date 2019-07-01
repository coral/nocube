package main

import (
	"fmt"

	opc "github.com/coral/go-opc"
)

func main() {

	// Create a client
	c := opc.NewClient()

	c.Connect("tcp", "localhost:7890")

	// Make a message!
	// This creates a message to send on channel 0
	// Or according to the OPC spec, a Broadcast Message.

	m := opc.NewMessage(0)
	m.SetLength(3 * 864)
	// Set pixel #1 to white.
	for i := 0; i < 864; i++ {

		m.SetPixelColor(i, 254, 254, 254)
	}

	fmt.Println(m.ByteArray())
	c.Send(m)
	// Send the message!

	// The first pixel of all registered devices should be white!
}
