package dynamic

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/coral/nocube/pkg"
	"github.com/fsnotify/fsnotify"
)

type Dynamic struct {
	patternPath     string
	loadedLibraries string
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
		if s.Name() != "node_modules" && s.Name() != "system" && s.IsDir() {

			pfiles, err := ioutil.ReadDir(d.patternPath + "/" + s.Name())
			if err != nil {
				fmt.Println(err)
			}
			for _, pfile := range pfiles {
				if filepath.Ext(pfile.Name()) == ".js" {
					d.Patterns[pfile.Name()] = CreatePattern(
						pfile.Name(),
						d.patternPath+"/"+s.Name()+"/"+pfile.Name(),
						d.patternPath,
					)
					pp := pfile.Name()
					d.Patterns[pp].Load(d.mapping)
				}
			}

		}
	}

}

func (d *Dynamic) reloadPattern(name string, path string) {
	d.Patterns[name].Unload()
	d.Patterns[name].Load(d.mapping)
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
		if s.Name() != "node_modules" && s.Name() != "system" && s.IsDir() {
			fmt.Println("Watching: " + s.Name())
			err = d.watcher.Add(d.patternPath + "/" + s.Name())
			if err != nil {
				log.Fatal("PEPEGA")
			}
		}
	}
}
