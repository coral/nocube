package mapping

import (
	"encoding/json"
	"io/ioutil"
	"math"

	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/utils"
	"github.com/stojg/vector"
)

type Mapping struct {
	Coordinates []pkg.Pixel

	path string
}

func New(path string, numPixels uint) *Mapping {
	return &Mapping{
		path:        path,
		Coordinates: make([]pkg.Pixel, numPixels),
	}
}

func (m *Mapping) WriteFile() error {
	data, err := m.MarshalJSON()
	if err != nil {
		return err
	}
	ioutil.WriteFile("../../files/"+m.path+".json", data, 0644)

	return nil
}

func (m *Mapping) LoadFile() error {
	b, err := ioutil.ReadFile("../../files/" + m.path + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &m.Coordinates)
}

func (m Mapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Coordinates)
}

func (m Mapping) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m)
}

func (m *Mapping) Insert(startIndex, stopIndex int, startVector, stopVector vector.Vector3) {
	length := stopIndex - startIndex

	dir := stopVector.NewSub(&startVector)
	dir[0] = math.Abs(dir[0])
	dir[1] = math.Abs(dir[1])
	dir[2] = math.Abs(dir[2])

	dir.Normalize()

	for index := 0; index <= length; index++ {
		val := float64(index) / float64(length)

		m.Coordinates[index+startIndex].Active = true
		m.Coordinates[index+startIndex].Coordinate = utils.VectorLerp(startVector, stopVector, val)
		m.Coordinates[index+startIndex].Normal = *dir
	}
}
