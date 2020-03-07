package audio

import (
	"time"

	aubio "github.com/coral/aubio-go"
	"github.com/coral/nocube/pkg/settings"
	"github.com/gordonklaus/portaudio"
	log "github.com/sirupsen/logrus"
)

type Audio struct {
	s        *settings.Settings
	Input    Input
	Analysis analysis

	Tempo      Tempo
	LastBeat   time.Time
	BeatNumber int64
	FFT        []float64
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
	err := portaudio.Initialize()
	if err != nil {
		log.Panic("Could not initialize Portaudio")
	}
	log.Debug("PortAudio init success")

	dev, _ := portaudio.Devices()

	var stParam portaudio.StreamParameters

	for _, device := range dev {
		if device.Name == a.s.Global.Audio.PreferredDevice {
			log.Debug("Found requested audio device: " + device.Name)
			var mock *portaudio.DeviceInfo
			stParam = portaudio.LowLatencyParameters(device, mock)
		}
	}

	a.Input.Buffer = make([]float32, a.s.Global.Audio.BufSize)

	var tempStream *portaudio.Stream

	if stParam.Input.Device != nil {

		log.Debug("Opening requested audio device")
		stParam.Input.Channels = a.s.Global.Audio.Channels
		stParam.Output.Channels = 0
		stParam.SampleRate = a.s.Global.Audio.SampleRate
		stParam.FramesPerBuffer = len(a.Input.Buffer)

		tempStream, err = portaudio.OpenStream(stParam, a.Input.Buffer)

	} else {

		log.Debug("Opening Default audio device")
		tempStream, err = portaudio.OpenDefaultStream(a.s.Global.Audio.Channels,
			0,
			a.s.Global.Audio.SampleRate,
			len(a.Input.Buffer), a.Input.Buffer)
	}

	if err != nil {
		log.Fatalln("Could not open audio stream")
	}
	a.Input.Stream = tempStream

	go func() {
		for {
			a.Tempo = <-a.Analysis.TempoStream
			a.LastBeat = time.Now()
			a.BeatNumber = a.BeatNumber + 1

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
