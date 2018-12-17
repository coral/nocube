package utils

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
