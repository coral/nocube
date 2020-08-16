package main

import "github.com/coral/nocube/pkg/dynamic"

func main() {

	dynamic := dynamic.New("../../dynamic/")
	dynamic.Initialize()

	blockChan := make(chan bool)

	_ = <-blockChan
}
