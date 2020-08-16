package dynamic

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	v8 "github.com/joesonw/js8"
)

type Dynamic struct {
	patternPath     string
	loadedLibraries string
	libSnapshot     *v8.Snapshot
	watcher         *fsnotify.Watcher
}

func New(PatternPath string) *Dynamic {
	return &Dynamic{
		patternPath: PatternPath,
	}
}

func (d *Dynamic) Initialize() {

	d.startWatcher()
	d.createV8Snapshot()
	//defer watcher.Close()
}

func (d *Dynamic) createV8Snapshot() {
	files, err := ioutil.ReadDir(d.patternPath + "/libs/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		dat, err := ioutil.ReadFile(d.patternPath + "/libs/" + file.Name())
		if err != nil {
			panic(err)
		}

		d.loadedLibraries = d.loadedLibraries + string(dat)

	}

	d.libSnapshot = v8.CreateSnapshot(d.loadedLibraries)

}

func (d *Dynamic) startWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	d.watcher = watcher

	go func() {

		for {
			select {
			case event, ok := <-d.watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-d.watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	files, err := ioutil.ReadDir(d.patternPath)
	if err != nil {
		log.Fatal("FUCK THIS MUSIC")
	}

	for _, s := range files {
		if s.Name() != "libs" {
			fmt.Println("Watching: " + s.Name())
			err = d.watcher.Add(d.patternPath + "/" + s.Name())
			if err != nil {
				log.Fatal("PEPEGA")
			}
		}
	}
}
