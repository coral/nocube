package main

import (
	"github.com/coral/nocube/pkg/mapping"
	"github.com/stojg/vector"
)

func main() {
	m := mapping.New("v5", 864)

	//STANDING MAP
	//0,0,0 - > 1,0,0
	//A1
	m.Insert(0, 71, vector.Vector3{0.05, 0, 0}, vector.Vector3{0.95, 0, 0})

	//1,0,0 -> 1,1,0
	//A2
	m.Insert(72, 143, vector.Vector3{1, 0.05, 0}, vector.Vector3{1, 1, 0})

	//1,1,0 -> 1,1,1
	//A3
	m.Insert(143, 215, vector.Vector3{1, 1, 0.05}, vector.Vector3{1, 1, 0.95})

	//1,1,1 -> 1,0,1
	//A4
	m.Insert(215, 287, vector.Vector3{1, 0.95, 1}, vector.Vector3{1, 0.05, 1})

	//1,0,1 -> 1,0,0
	//A5
	m.Insert(287, 359, vector.Vector3{1, 0, 0.95}, vector.Vector3{1, 0, 0.05})

	//0,0,0 -> 0,0,1
	//A6
	m.Insert(359, 431, vector.Vector3{0, 0, 0.05}, vector.Vector3{0, 0, 0.95})

	//0,0,0 -> 0,1,0
	//B1
	m.Insert(431, 503, vector.Vector3{0, 0.05, 0}, vector.Vector3{0, 0.95, 0})

	//0,1,0 -> 0,1,1
	//B2
	m.Insert(503, 575, vector.Vector3{0, 1, 0.05}, vector.Vector3{0, 1, 0.95})

	//0,1,1 -> 1,1,1
	//B3
	m.Insert(575, 647, vector.Vector3{0.05, 1, 1}, vector.Vector3{0.95, 1, 1})

	//1,0,1 -> 0,0,1
	//B4
	m.Insert(647, 719, vector.Vector3{0.95, 0, 1}, vector.Vector3{0.05, 0, 1})

	//0,0,1 -> 0,1,1
	//B5
	m.Insert(719, 791, vector.Vector3{0, 0.05, 1}, vector.Vector3{0, 0.95, 1})

	//0,1,0 -> 1,1,0
	//B6
	m.Insert(791, 863, vector.Vector3{0.05, 1, 0}, vector.Vector3{0.95, 1, 0})
	err := m.WriteFile()
	if err != nil {
		panic(err)
	}
}
