package frame

import "math"

//GetSine returns a 0-1 sine wave based on the phase of the frame.
func (f *F) GetSine(offset float64) float64 {
	return math.Sin((f.Phase + offset) * math.Pi * 2)
}

//GetCos returns a 0-1 cosine wave based on the phase of the frame.
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

func (f *F) GetSegment(numSegments int64) (segmentIndex int64, remainder float64) {
	segmentIndex = int64(math.Floor(float64(numSegments) * f.Phase))
	remainder = math.Mod(float64(numSegments)*f.Phase, 1)

	return
}

//GetBeat tries to return a bool based on if you are in the cycle of a beat
func (f *F) GetBeat(beat uint64, offset uint64) bool {

	d := math.Mod(float64(f.audioHolder.BeatNumber), 4)
	if d == float64(offset) {
		return true
	} else {
		return false
	}

}
