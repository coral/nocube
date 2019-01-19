package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

var saveTimer *time.Ticker

type Settings struct {
	Path   string
	Global struct {
		Mapping struct {
			Path string
		}
		Control struct {
			Web struct {
				Listen string
			}
		}
		Render struct {
			InternalTargetFPS int
		}
		Audio struct {
			SampleRate float64
			Channels   int
			BufSize    uint
			BlockSize  uint
			Beat       struct {
				Silence   float32
				Threshold float32
			}
		}
	}
}

func New(path string) (*Settings, error) {

	b, err := ioutil.ReadFile("../../files/settings/" + path + ".json")
	if err != nil {
		ns := Settings{}
		ns.Path = path
		saveTimer = time.NewTicker(5 * time.Minute)
		go ns.Save()
		return &ns, nil
	}

	ns := Settings{}
	err = json.Unmarshal(b, &ns)
	if err != nil {
		return nil, err
	}

	saveTimer = time.NewTicker(5 * time.Minute)
	go ns.Save()

	return &ns, nil
}

func (s *Settings) Save() {
	for {
		select {
		case <-saveTimer.C:
			data, err := json.MarshalIndent(s, "", "	")
			if err != nil {
				// TODO
				// Replace with real logging library
				fmt.Println("COULD NOT SAVE")
				fmt.Println(err)
				panic(err)
			}

			err = ioutil.WriteFile("../../files/settings/"+s.Path+".json", data, 0644)
			if err != nil {
				// TODO
				//Replace with real logging
				fmt.Println(err)
			}

			log.Info("Saved settings to " + s.Path + ".json")
		}
	}
}
