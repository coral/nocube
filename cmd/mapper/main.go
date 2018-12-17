package main

import (
	"github.com/coral/nocube/pkg/mapping"
	"github.com/stojg/vector"
)

func main() {
	m := mapping.New("v1", 920)

	m.Insert(9, 80, vector.Vector3{0, 0.95, 0}, vector.Vector3{0, 0.05, 0})
	m.Insert(87, 158, vector.Vector3{0.05, 0, 0}, vector.Vector3{0.95, 0, 0})
	m.Insert(165, 237, vector.Vector3{1, 0.05, 0}, vector.Vector3{1, 0.95, 0})
	m.Insert(246, 305, vector.Vector3{1, 1, 0.05}, vector.Vector3{1, 1, 0.95})
	m.Insert(313, 385, vector.Vector3{0.95, 1, 1}, vector.Vector3{0.05, 1, 1})
	m.Insert(393, 464, vector.Vector3{0, 1, 0.95}, vector.Vector3{0, 1, 0.05})
	m.Insert(472, 544, vector.Vector3{0.05, 1, 0}, vector.Vector3{0.95, 1, 0})

	m.Insert(623, 694, vector.Vector3{1, 0.95, 1}, vector.Vector3{1, 0.05, 1})

	m.Insert(701, 772, vector.Vector3{0.95, 0, 1}, vector.Vector3{0.05, 0, 1})

	m.Insert(779, 850, vector.Vector3{0, 0, 0.95}, vector.Vector3{0, 0, 0.05})

	err := m.WriteFile()
	if err != nil {
		panic(err)
	}
}
