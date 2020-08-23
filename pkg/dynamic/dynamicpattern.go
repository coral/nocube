package dynamic

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	v8 "github.com/augustoroman/v8"
	"github.com/augustoroman/v8/v8console"
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type DynamicPattern struct {
	PatternName string
	path        string
	code        string
	v8isolate   *v8.Isolate
	v8ctx       *v8.Context
	Loaded      bool

	buffer []float64
}

func CreatePattern(Name string, path string) *DynamicPattern {
	return &DynamicPattern{
		PatternName: Name,
		path:        path,
	}
}

type Data struct {
	Length int
	Pixels []V8Mapping
}

type V8Mapping struct {
	Index  int64
	Active bool
	X      float64
	Y      float64
	Z      float64
}

func (dp *DynamicPattern) Load(s *v8.Snapshot, m []pkg.Pixel) {

	dp.buffer = make([]float64, len(m))

	code, err := ioutil.ReadFile(dp.path)
	if err != nil {
		panic(err)
	}

	dp.v8isolate = v8.NewIsolateWithSnapshot(s)
	dp.v8ctx = dp.v8isolate.NewContext()

	var tm []V8Mapping
	for _, km := range m {
		tm = append(tm, V8Mapping{
			Index:  km.Index,
			Active: km.Active,
			X:      km.Coordinate[0],
			Y:      km.Coordinate[1],
			Z:      km.Coordinate[2],
		})
	}

	mapdata := Data{
		Length: len(m),
		Pixels: tm,
	}

	v8map, err := dp.v8ctx.Create(mapdata)
	if err != nil {
		panic(err)
	}

	dp.v8ctx.Global().Set("mapping", v8map)

	dp.v8ctx.Eval(string(code), dp.PatternName)

	v8console.Config{"", os.Stdout, os.Stderr, true}.Inject(dp.v8ctx)
	dp.Loaded = true
	fmt.Println("Loaded Pattern " + dp.PatternName)
}

func (dp *DynamicPattern) Gen(pixels []pkg.Pixel, f *frame.F) []pkg.Pixel {
	if dp.Loaded {

		fmt.Println(dp.Loaded)
		res, err := dp.v8ctx.Eval(`render()`, "demo.js")
		if err != nil {
			fmt.Println(err)
			//panic(err)
		}
		_ = res
		// result := dp.ineffecientBytesToFloat64Slice(res.Bytes(), len(pixels))
		// _ = result
		// fmt.Println(result)

	}

	return pixels
}

//TODO: REWRITE THIS SHIT FUNCTION INTO SOMETHING THAT ISNT SHIT

func (dp *DynamicPattern) ineffecientBytesToFloat64Slice(bytes []byte, len int) []float64 {
	dp.buffer = dp.buffer[:0]

	for i := 0; i < len; i++ {
		dp.buffer[i] = math.Float64frombits(binary.LittleEndian.Uint64(bytes[i*8 : (i+1)*8]))
	}

	return dp.buffer

}

// func BytesToFloat64(bytes []byte) []float64 {
// 	var s []float64
// 	s = make([]float64, 864)
// 	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
// 	stringHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))
// 	stringHeader.Data = sliceHeader.Data
// 	stringHeader.Len = sliceHeader.Len
// 	runtime.KeepAlive(&bytes)
// 	return s
// }

func (dp *DynamicPattern) Unload() {
	dp.Loaded = false
	dp.v8isolate.Terminate()
}

func (dp *DynamicPattern) Init() {

}

func (dp *DynamicPattern) Name() string {
	return dp.PatternName
}

func (dp *DynamicPattern) Type() string {
	return "dynamic"
}
