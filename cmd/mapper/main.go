package main

import (
	"github.com/coral/nocube/pkg/mapping"
	"github.com/stojg/vector"
)

func main() {
	m := mapping.New("v3", 870)

	//0,0,0 - > 1,0,0
	m.Insert(1, 72, vector.Vector3{0.05, 0, 0}, vector.Vector3{0.95, 0, 0})

	//1,0,0 -> 1,0,1
	m.Insert(73, 144, vector.Vector3{1, 0, 0.05}, vector.Vector3{1, 0, 0.95})

	//1,0,1 -> 0,0,1
	m.Insert(144, 216, vector.Vector3{0.95, 0, 1}, vector.Vector3{0.05, 0, 1})

	//0,0,1 -> 0,0,0
	m.Insert(217, 288, vector.Vector3{0, 0, 0.95}, vector.Vector3{0, 0, 0.05})

	//0,0,0 -> 0,1,0
	m.Insert(289, 360, vector.Vector3{0, 0.05, 0}, vector.Vector3{0, 0.95, 0})

	//0,1,0 -> 0,1,1
	m.Insert(361, 432, vector.Vector3{0, 1, 0.05}, vector.Vector3{0, 1, 0.95})

	//0,1,1 -> 0,0,1
	m.Insert(433, 504, vector.Vector3{0, 0.95, 1}, vector.Vector3{0, 0.05, 1})

	//0,1,1 -> 1,1,1
	m.Insert(505, 576, vector.Vector3{0.05, 1, 1}, vector.Vector3{0.95, 1, 1})

	//1,1,1 -> 1,1,0
	m.Insert(577, 648, vector.Vector3{1, 1, 0.95}, vector.Vector3{1, 1, 0.05})

	//1,1,0 -> 0,1,0
	m.Insert(649, 720, vector.Vector3{0.95, 1, 0}, vector.Vector3{0.05, 1, 0})

	//1,1,1 -> 1,0,1
	m.Insert(721, 792, vector.Vector3{1, 0.95, 1}, vector.Vector3{1, 0.05, 1})

	//1,1,0 -> 1,0,0
	m.Insert(793, 864, vector.Vector3{1, 0.95, 0}, vector.Vector3{1, 0.05, 0})

	// m.Insert(0, 31, vector.Vector3{0, 0, 0}, vector.Vector3{0, 1, 0})
	err := m.WriteFile()
	if err != nil {
		panic(err)
	}
}
