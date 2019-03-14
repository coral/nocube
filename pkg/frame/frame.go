package frame

import (
	"math"
	"time"

	"github.com/coral/nocube/pkg/audio"
	"github.com/coral/nocube/pkg/render"
)

type F struct {
	Timepoint float64
	Index     uint64

	BeatDuration float64
	BeatStart    float64

	Phase float64

	FFT []float64

	renderHolder *render.Render
	renderSignal chan render.Update
	audioHolder  *audio.Audio

	OnUpdate chan *F
}

func New(newR *render.Render, audio *audio.Audio) F {

	newF := F{
		Index:        0,
		BeatDuration: 60.0 / 120.0,
		BeatStart:    0.0,
		Phase:        0.0,
		renderHolder: newR,
		renderSignal: make(chan render.Update),
		audioHolder:  audio,
		OnUpdate:     make(chan *F),
		FFT:          make([]float64, 128),
	}

	newF.renderHolder.Update.Register(newF.renderSignal)

	go func(chan render.Update) {
		for {
			select {
			case v := <-newF.renderSignal:
				newF.Update(v)
			}
		}
	}(newF.renderSignal)

	return newF
}
func (f *F) Update(u render.Update) {

	if f.audioHolder.Tempo.Confidence > 0.05 {
		f.BeatDuration = 60 / f.audioHolder.Tempo.Tempo
	}
	f.Timepoint = float64(u.TimeSinceStart/time.Millisecond) / 1000

	f.Phase = math.Mod((f.Timepoint-f.BeatStart)/f.BeatDuration, 1)
	f.Index = u.FrameNumber
	f.FFT = f.audioHolder.FFT

	f.OnUpdate <- f
}

func (f *F) SetBeat(beatduration, beatstart float64) {
	f.BeatDuration = beatduration
	f.BeatStart = beatstart
}

func (f *F) GetSine(offset float64) float64 {
	return math.Sin((f.Phase + offset) * math.Pi * 2)
}

func (f *F) GetCos(offset float64) float64 {
	return math.Cos((f.Phase + offset) * math.Pi * 2)
}

//TODO implement
func (f *F) GetTriangle(offset float64) float64 {
	return 1
}

func (f *F) GetSquare() float64 {
	if f.Phase >= 0.5 {
		return 1.0
	} else {
		return 0.0
	}
}

func (f *F) GetSegment(numSegments uint64) (segmentIndex uint64, remainder float64) {
	segmentIndex = uint64(math.Floor(float64(numSegments) * f.Phase))
	remainder = math.Mod(float64(numSegments)*f.Phase, 1)

	return
}
