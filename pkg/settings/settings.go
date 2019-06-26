package settings

import (
	"encoding/json"
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
			OSC struct {
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
			FFTSize int
		}
		Output struct {
			ActivatedOutputs []string
			RAPA102          []struct {
				Identifier string
				TargetFPS  int
			}
		}
	}
}

func New(path string) (*Settings, error) {

	b, err := ioutil.ReadFile("../../files/settings/" + path + ".json")
	if err != nil {
		ns := Settings{}
		ns.Path = path
		saveTimer = time.NewTicker(20 * time.Minute)
		go ns.Save()
		return &ns, nil
	}

	ns := Settings{}
	err = json.Unmarshal(b, &ns)
	if err != nil {
		return nil, err
	}

	saveTimer = time.NewTicker(20 * time.Minute)
	go ns.Save()

	return &ns, nil
}

func (s *Settings) Save() {
	for {
		select {
		case <-saveTimer.C:
			data, err := json.MarshalIndent(s, "", "	")
			if err != nil {
				log.Error("could not save settings")
				log.Error(err)
				panic(err)
			}

			err = ioutil.WriteFile("../../files/settings/"+s.Path+".json", data, 0644)
			if err != nil {
				log.Error(err)
			}

			log.Info("Saved settings to " + s.Path + ".json")
		}
	}
}
