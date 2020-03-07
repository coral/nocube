package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/ry/v8worker2"
)

func main() {

	ExampleNewWatcher()
}

func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					ExecJS()
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("ay.js")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func ExecJS() {
	d := v8worker2.New(func(msg []byte) []byte {
		fmt.Println(msg)
		return nil
	})
	ff, _ := ioutil.ReadFile("ay.js")
	err := d.Load("code.js", string(ff))
	if err != nil {
		fmt.Println(err)
	}

	err = d.SendBytes([]byte("hii"))
	if err != nil {
		fmt.Println(err)
	}

}
