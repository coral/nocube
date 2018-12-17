package utils

import "github.com/stojg/vector"

func VectorLerp(a, b vector.Vector3, f float64) vector.Vector3 {
	l := b.NewSub(&a)

	asd2 := l.Scale(f)

	return *a.NewAdd(asd2)
}
