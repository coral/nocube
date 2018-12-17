package frame

import (
	"fmt"
	"math"
)

type F struct {
	Timepoint float64
	Index     uint64

	BeatDuration float64
	BeatStart    float64

	Phase float64
}

func New() F {
	newF := F{
		Index:        0,
		BeatDuration: 60.0 / 120.0,
		BeatStart:    0.0,
	}

	fmt.Println(newF)
	return newF
}
func (f *F) Update(timepoint float64) {
	f.Timepoint = timepoint
	f.Phase = math.Mod((f.Timepoint-f.BeatStart)/f.BeatDuration, 1)
	f.Index++
}

func (f *F) SetBeat(beatduration, beatstart float64) {
	f.BeatDuration = beatduration
	f.BeatStart = beatstart
}

func (f *F) GetSine(offset float64) float64 {
	return math.Sin(f.Phase * math.Pi * 2)
}

func (f *F) GetCos(offset float64) float64 {
	return math.Cos(f.Phase * math.Pi * 2)
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
	remainder = math.Mod(float64(numSegments) * f.Phase, 1)

	return
}
