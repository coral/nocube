package audio

import (
	aubio "github.com/coral/aubio-go"
	"github.com/coral/nocube/pkg/settings"
	"github.com/gordonklaus/portaudio"
	log "github.com/sirupsen/logrus"
)

type Audio struct {
	s        *settings.Settings
	Input    Input
	Analysis analysis

	Tempo Tempo
	FFT   []float64
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

	a.FFT = make([]float64, a.s.Global.Audio.FFTSize)
	a.Analysis.Init(a.s)

	log.Debug("Init PortAudio")
	portaudio.Initialize()
	log.Debug("PortAudio init success")

	a.Input.Buffer = make([]float32, a.s.Global.Audio.BufSize)

	t, err := portaudio.OpenDefaultStream(a.s.Global.Audio.Channels,
		0,
		a.s.Global.Audio.SampleRate,
		len(a.Input.Buffer), a.Input.Buffer)

	if err != nil {
		log.Fatalln("Could not open audio stream")
	}
	a.Input.Stream = t

	go func() {
		for {
			a.Tempo = <-a.Analysis.TempoStream

		}
	}()

	go func() {
		for {
			a.FFT = <-a.Analysis.FFTStream

		}
	}()

	return nil
}

func (a *Audio) Close() {
	portaudio.Terminate()
	a.Input.Stream.Close()
}

//Process starts the audio processing, this function is locking
func (a *Audio) Process() {

	err := a.Input.Stream.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Started audio processing")
	for {
		err := a.Input.Stream.Read()
		if err != nil {
			log.Fatal(err)
		} else {
			a.Input.ProcessedSamples += int64(len(a.Input.Buffer))
			convertedBuffer := convertTo64(a.Input.Buffer)
			b := aubio.NewSimpleBufferData(a.s.Global.Audio.BufSize, convertedBuffer)

			a.Analysis.Do(b)
		}
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
