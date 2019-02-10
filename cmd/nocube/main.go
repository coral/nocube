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

// const NUM_LEDS = 920
// const INTENSITY = 30
// const PUSHER_DURATION = 5 * time.Millisecond

// var packagesPushedThisSecond uint

// func pusher(d *apa102.Dev, in chan []byte, stop chan bool, debug chan bool) {
// 	for {
// 		select {
// 		case bytes := <-in:
// 			_, err := d.Write(bytes)
// 			if err != nil {
// 				log.Fatal("xd", err)
// 			}

// 			debug <- true

// 		case <-stop:
// 			return
// 		}
// 	}
// }

// func debugger(packagePushed, stop chan bool) {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-packagePushed:
// 			packagesPushedThisSecond++
// 		case <-ticker.C:
// 			fmt.Println("Packages pushed this second:", packagesPushedThisSecond)
// 			packagesPushedThisSecond = 0

// 		case <-stop:
// 			return
// 		}
// 	}
// }

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

	render := render.New(settings)
	render.Start()

	frame := frame.New(render)
	frame.SetBeat(60.0/30.0, 0)

	Pipelines := pipelines.New(&frame, mapping)
	test := pipeline.New("zebra", "allwhite")
	Pipelines.Add(test)

	go func() {
		for v := range frame.OnUpdate {
			Pipelines.Process(v)
		}
	}()

	server := web.Server{}
	server.Init(settings, mapping)

	/*
		generatorStop := make(chan bool)
		pusherStop := make(chan bool)
		debuggerStop := make(chan bool)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for _ = range c {
				generatorStop <- true
				pusherStop <- true
				debuggerStop <- true
			}
		}()

		bytesChannel := make(chan []byte, 1)
		debuggerChannel := make(chan bool, 10)

		go pusher(d, bytesChannel, pusherStop, debuggerChannel)

		go debugger(debuggerChannel, debuggerStop)

		generator(mapping, bytesChannel, generatorStop) */
}

/* func generator(mapping *mapping.Mapping, bytesChannel chan []byte, stop chan bool) {
	ticker := time.NewTicker(PUSHER_DURATION)
	defer ticker.Stop()

	generators := map[string]pkg.Generator{
		"xd":       &xd.Xd{},
		"edgelord": &edgelord.Edgelord{},
		"zebra":    &zebra.Zebra{},
	}

	colorLookups := map[string]pkg.ColorLookup{
		"dummy":   &dummy.Dummy{},
		"palette": &palette.Palette{},
	}

	var t float64
	frame := frame.New()
	frame.SetBeat(60.0/30.0, 0)

	generator := generators["edgelord"]
	colorLookup := colorLookups["palette"]

	if generator == nil || colorLookup == nil {
		panic("No valid generator or color lookup chosen")
	}

	for {
		select {
		case <-ticker.C:
			frame.Update(t)
			res := generator.Generate(mapping.Coordinates, &frame, pkg.GeneratorParameters{})
			colorRes := colorLookup.Lookup(res, &frame, pkg.ColorLookupParameters{})

			var bytes = []byte{}
			for _, color := range colorRes {
				bytes = append(bytes, []byte{
					utils.Clamp255(color.Color[0] * 255),
					utils.Clamp255(color.Color[1] * 255),
					utils.Clamp255(color.Color[2] * 255),
				}...)
			}

			bytesChannel <- bytes

			t += PUSHER_DURATION.Seconds()

		case <-stop:
			return
		}
	}
} */
