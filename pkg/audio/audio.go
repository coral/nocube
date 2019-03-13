package audio

import (
	aubio "github.com/coral/aubio-go"
	"github.com/coral/nocube/pkg/settings"
	"github.com/gordonklaus/portaudio"
	log "github.com/sirupsen/logrus"
)

type Audio struct {
	s     *settings.Settings
	Input Input

	TempoStream chan T
}

type Input struct {
	ProcessedSamples int64
	Buffer           []float32
	Stream           *portaudio.Stream
	Tempo            *aubio.Tempo
}

type T struct {
	Beat       float64
	Tempo      float64
	Confidence float64
}

func New(s *settings.Settings) *Audio {
	return &Audio{
		s:           s,
		TempoStream: make(chan T),
		Input: Input{
			ProcessedSamples: 0,
		},
	}
}

//Init the audio object
func (a *Audio) Init() error {
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

	a.Input.Tempo = aubio.TempoOrDie(aubio.SpecDiff,
		a.s.Global.Audio.BufSize,
		a.s.Global.Audio.BlockSize,
		uint(a.s.Global.Audio.SampleRate))
	a.Input.Tempo.SetSilence(-70.0)

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
			a.processTempo(b)
		}
		//go lel(in)
	}
}

func (a *Audio) processTempo(b *aubio.SimpleBuffer) {
	a.Input.Tempo.Do(b)
	for _, f := range a.Input.Tempo.Buffer().Slice() {
		if f != 0 {

			a.TempoStream <- T{
				Beat:       f,
				Tempo:      a.Input.Tempo.GetBpm(),
				Confidence: a.Input.Tempo.GetConfidence(),
			}
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
