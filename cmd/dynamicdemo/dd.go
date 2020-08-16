package main

import (
	"time"

	"github.com/coral/nocube/pkg/dynamic"
)

func main() {

	dynamic := dynamic.New("../../dynamic/")
	dynamic.Initialize()

	for {
		for _, p := range dynamic.Patterns {
			p.Exec()
		}
		time.Sleep(1 * time.Second)
	}

	blockChan := make(chan bool)

	_ = <-blockChan
}
