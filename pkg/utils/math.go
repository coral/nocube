package utils

import "math/rand"

func Clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}

	return v
}

func Clamp01(v float64) float64 {
	if v < 0.0 {
		return 0.0
	}
	if v > 1.0 {
		return 1.0
	}

	return v
}

func Clamp255(v float64) byte {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}

	return byte(v)
}

//Crush will return 1 if over the threshold
func Crush(v float64, threshold float64) float64 {
	if v > threshold {
		return 1
	} else {
		return v
	}
}

func Threshold(v float64, threshold float64) float64 {
	if v > threshold {
		return v
	} else {
		return 0
	}
}

func Trigger(v float64, trigger bool) float64 {
	if trigger {
		return v
	} else {
		return 0.0
	}
}

func Random(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
