package main

import (
	"flag"

	"github.com/coral/nocube/pkg/audio"
	"github.com/coral/nocube/pkg/data"
	"github.com/coral/nocube/pkg/dynamic"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/layers"
	"github.com/coral/nocube/pkg/mapping"
	"github.com/coral/nocube/pkg/output"
	"github.com/coral/nocube/pkg/render"
	"github.com/coral/nocube/pkg/settings"

	log "github.com/sirupsen/logrus"
)

func main() {

	var benchmark = flag.Bool("benchmark", false, "Print system FPS")
	flag.Parse()

	//Initialize logger
	log.SetLevel(log.DebugLevel)
	log.Info("Starting nocube")

	//Load settings
	settings, err := settings.New("start")
	if err != nil {
		panic(err)
	}

	db := data.New()
	db.Init()

	log.WithFields(log.Fields{
		"Settings": settings.Path + ".json",
		"Mapping":  settings.Global.Mapping.Path + ".json",
		"Web":      settings.Global.Control.Web.Listen,
	}).Info("Loaded settings")

	log.Debug(settings)

	a := audio.New(settings)
	a.Init()
	defer a.Close()
	go a.Process()

	mapping, err := mapping.LoadNewFromFile(settings.Global.Mapping.Path)
	if err != nil {
		panic(err)
	}

	output := output.New(settings)
	output.Init()

	rend := render.New(settings.Global.Render.InternalTargetFPS, *benchmark)
	rend.Start()

	frame := frame.New(rend, a)
	frame.SetBeat(60.0/30.0, 0)

	dn := dynamic.New("../../dynamic/")
	dn.Initialize(mapping.Coordinates)

	layer := layers.New("msdn", &frame, mapping, &db, dn)
	layer.Initialize()

	go func() {

		for {
			select {
			case v := <-frame.OnUpdate:
				p := layer.Process(v)
				go output.Write(p)
			}
		}

	}()
	blockChan := make(chan bool)

	_ = <-blockChan

	//api := api.New(mapping, Pipelines, &db)
	//api.Init(settings)

}
