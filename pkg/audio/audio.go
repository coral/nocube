package audio

import (
	"fmt"

	aubio "github.com/coral/aubio-go"
	"github.com/coral/nocube/pkg/settings"
	"github.com/gordonklaus/portaudio"
	log "github.com/sirupsen/logrus"
)

type Audio struct {
	s     *settings.Settings
	Input Input
}

type Input struct {
	ProcessedSamples int64
	Buffer           []float32
	Stream           *portaudio.Stream
}

func New(s *settings.Settings) *Audio {
	return &Audio{
		s: s,
		Input: Input{
			ProcessedSamples: 0,
		},
	}
}

//Init the audio object
func (a *Audio) Init() error {
	log.Debug("Init PortAudio")
	portaudio.Initialize()
	defer portaudio.Terminate()
	log.Debug("PortAudio init success")

	a.Input.Buffer = make([]float32, a.s.Global.Audio.BufSize)

	t, err := portaudio.OpenDefaultStream(a.s.Global.Audio.Channels,
		0,
		a.s.Global.Audio.SampleRate,
		len(a.Input.Buffer), a.Input.Buffer)

	if err != nil {
		log.Fatalln("Could not open audio stream")
	}
	fmt.Println(t)
	a.Input.Stream = t
	defer a.Input.Stream.Close()

	return nil
}

//Process starts the audio processing, this function is locking
func (a *Audio) Process() {

	fmt.Println(a.Input.Stream)

	err := a.Input.Stream.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		err := a.Input.Stream.Read()
		if err != nil {
			log.Fatal(err)
		} else {
			a.Input.ProcessedSamples += int64(len(a.Input.Buffer))
			convertedBuffer := convertTo64(a.Input.Buffer)
			aubio.NewSimpleBufferData(a.s.Global.Audio.BufSize, convertedBuffer)

		}
		//go lel(in)
	}
}
func convertTo64(ar []float32) []float64 {
	newar := make([]float64, len(ar))
	var v float32
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}
