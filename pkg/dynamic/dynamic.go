package dynamic

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	v8 "github.com/augustoroman/v8"
	"github.com/coral/nocube/pkg"
	"github.com/fsnotify/fsnotify"
)

type Dynamic struct {
	patternPath     string
	loadedLibraries string
	libSnapshot     *v8.Snapshot
	watcher         *fsnotify.Watcher
	Patterns        map[string]*DynamicPattern
	mapping         []pkg.Pixel
}

func New(PatternPath string) *Dynamic {
	return &Dynamic{
		patternPath: PatternPath,
		Patterns:    make(map[string]*DynamicPattern),
	}
}

func (d *Dynamic) Initialize(m []pkg.Pixel) {

	d.mapping = m

	d.startWatcher()
	d.createV8Snapshot()
	d.loadPatterns()
	//defer watcher.Close()
}

func (d *Dynamic) CreateSteps() []pkg.Step {
	var steps []pkg.Step
	for _, m := range d.Patterns {
		steps = append(steps, m)
	}
	return steps
}

func (d *Dynamic) loadPatterns() {

	files, err := ioutil.ReadDir(d.patternPath)
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range files {
		if s.Name() != "libs" {
			pfiles, err := ioutil.ReadDir(d.patternPath + "/" + s.Name())
			if err != nil {
				fmt.Println(err)
			}
			for _, pfile := range pfiles {
				d.Patterns[pfile.Name()] = CreatePattern(
					pfile.Name(),
					d.patternPath+"/"+s.Name()+"/"+pfile.Name(),
				)
				pp := pfile.Name()
				d.Patterns[pp].Load(d.libSnapshot, d.mapping)
			}

		}
	}

}

func (d *Dynamic) reloadPattern(name string, path string) {
	d.Patterns[name].Unload()
	d.Patterns[name].Load(d.libSnapshot, d.mapping)
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
					d.reloadPattern(filepath.Base(event.Name), event.Name)
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
