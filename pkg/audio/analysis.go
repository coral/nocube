package audio

import (
	"time"

	aubio "github.com/coral/aubio-go"
	"github.com/coral/nocube/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type analysis struct {
	Tempo       *aubio.Tempo
	Onset       *aubio.Onset
	BeatTracker *aubio.BeatTracker
	FFT         *aubio.FFT

	TempoStream chan Tempo
	FFTStream   chan []float64
}

type Tempo struct {
	Beat       float64
	Tempo      float64
	Confidence float64
	Timepoint  time.Time
}

func (a *analysis) Init(s *settings.Settings) {

	log.Debug("Loaded Aubio")
	a.TempoStream = make(chan Tempo)
	a.FFTStream = make(chan []float64)

	a.Tempo = aubio.TempoOrDie(aubio.SpecDiff,
		s.Global.Audio.BufSize,
		s.Global.Audio.BlockSize,
		uint(s.Global.Audio.SampleRate))
	a.Tempo.SetSilence(-70.0)

	a.Onset = aubio.OnsetOrDie(aubio.SpecDiff,
		s.Global.Audio.BufSize,
		s.Global.Audio.BlockSize,
		uint(s.Global.Audio.SampleRate))
	a.Onset.SetSilence(-70.0)
	a.Onset.SetThreshold(-1.0)

	a.FFT = aubio.NewFFT(uint(s.Global.Audio.FFTSize))

}

func (a *analysis) Do(b *aubio.SimpleBuffer) {
	a.processTempo(b)
	a.processOnset(b)
	a.processFFT(b)
}

func (a *analysis) processTempo(b *aubio.SimpleBuffer) {
	a.Tempo.Do(b)
	for _, f := range a.Tempo.Buffer().Slice() {

		if f != 0 {
			log.WithFields(log.Fields{
				"Tempo":      a.Tempo.GetBpm(),
				"Confidence": a.Tempo.GetConfidence(),
			}).Debug("Found beat!")

			a.TempoStream <- Tempo{
				Beat:       f,
				Tempo:      a.Tempo.GetBpm(),
				Confidence: a.Tempo.GetConfidence(),
				Timepoint:  time.Now(),
			}
		}
	}
}

func (a *analysis) processOnset(b *aubio.SimpleBuffer) {
	a.Onset.Do(b)
	for _, f := range a.Onset.Buffer().Slice() {
		if f != 0 {

		}
	}
}

func (a *analysis) processFFT(b *aubio.SimpleBuffer) {
	a.FFT.Do(b)
	a.FFTStream <- a.FFT.Buffer().Norm()
}
