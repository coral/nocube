package main

import (
	"fmt"

	"github.com/ry/v8worker2"
)

func main() {
	d := v8worker2.New(func(msg []byte) []byte {
		fmt.Println(msg)
		return nil
	})
	err := d.Load("code.js", `V8Worker2.print("JS");
	
		V8Worker2.recv(function(msg) {
			V8Worker2.print("TestBasic recv byteLength", msg.byteLength);
		});
	`)
	if err != nil {
		panic(err)
	}

	err = d.SendBytes([]byte("hii"))
	if err != nil {
		panic(err)
	}
}
