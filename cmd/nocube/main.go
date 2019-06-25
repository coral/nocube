package main

import (
	"github.com/coral/nocube/pkg/audio"
	"github.com/coral/nocube/pkg/control/web"
	"github.com/coral/nocube/pkg/frame"
	"github.com/coral/nocube/pkg/mapping"
	"github.com/coral/nocube/pkg/output"
	"github.com/coral/nocube/pkg/pipelines"
	"github.com/coral/nocube/pkg/pipelines/pipeline"
	"github.com/coral/nocube/pkg/render"
	"github.com/coral/nocube/pkg/settings"

	log "github.com/sirupsen/logrus"
)

func main() {

	//Initialize logger
	log.SetLevel(log.DebugLevel)
	log.Info("Starting nocube")

	//Load settings
	settings, err := settings.New("start")
	if err != nil {
		panic(err)
	}

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

	rend := render.New(settings)
	rend.Start()

	/* 	t := make(chan render.Update)
	   	rend.Update.Register(t) */

	frame := frame.New(rend, a)
	frame.SetBeat(60.0/30.0, 0)

	Pipelines := pipelines.New(&frame, mapping)
	cp := pipeline.New("olof", 1.0, "strobe", "colorize", "add")
	Pipelines.Add(cp)

	test := pipeline.New("denis", 0.0, "zebra", "dummy", "screen")
	Pipelines.Add(test)

	go func() {

		for {
			select {
			case v := <-frame.OnUpdate:
				p := Pipelines.Process(v)
				go output.Write(p)
			}
		}

	}()

	server := web.Server{}
	server.Init(settings, mapping)

}
