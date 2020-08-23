package dynamic

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	v8 "github.com/augustoroman/v8"
	"github.com/augustoroman/v8/v8console"
	"github.com/coral/nocube/pkg"
	"github.com/coral/nocube/pkg/frame"
)

type DynamicPattern struct {
	PatternName string
	path        string
	modulepath  string
	v8isolate   *v8.Isolate
	v8ctx       *v8.Context
	Loaded      bool
	Console     bool

	buffer       []float64
	outputBuffer []pkg.Pixel
}

func CreatePattern(Name string, path string, modulepath string) *DynamicPattern {
	return &DynamicPattern{
		PatternName: Name,
		path:        path,
		modulepath:  modulepath,
		Console:     true,
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

func (dp *DynamicPattern) Load(m []pkg.Pixel) {

	dp.buffer = make([]float64, len(m)*3)
	dp.outputBuffer = make([]pkg.Pixel, len(m))

	mp, err := filepath.Abs(dp.modulepath)
	if err != nil {
		fmt.Println("Could not get absolute path of js file")
	}

	ap, err := filepath.Abs(dp.path)
	if err != nil {
		fmt.Println("Could not get absolute path of js file")
	}

	systemLibs, err := ioutil.ReadFile(mp + "/system/system.js")
	if err != nil {
		log.WithFields(log.Fields{
			"System Lib Path": mp + "/system/system.js",
			"Error":           err,
		}).Error("Could not load system libs for dynamic patterns")
	}

	patternCode, err := ioutil.ReadFile(ap)
	if err != nil {
		log.WithFields(log.Fields{
			"Pattern Path": ap,
			"Error":        err,
		}).Error("Could not load system libs for dynamic patterns")
	}

	process := string(systemLibs) + string(patternCode)

	cmd := exec.Command("/usr/local/bin/node", mp+"/node_modules/browserify/bin/cmd.js", "-", "--s", "pattern")
	cmd.Dir = mp
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Error(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, process)
	}()

	code, err := cmd.CombinedOutput()

	if err != nil {
		log.WithFields(log.Fields{
			"Error":  err,
			"Stderr": string(code),
			"File":   ap,
		}).Error("Could not execute browserify")

		return
	}

	dp.v8isolate = v8.NewIsolate()
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
		log.Error(err)
	}

	dp.v8ctx.Global().Set("mapping", v8map)

	dp.v8ctx.Eval(string(code), dp.PatternName)

	if dp.Console {
		v8console.Config{"", os.Stdout, os.Stderr, true}.Inject(dp.v8ctx)
	}
	dp.Loaded = true
	fmt.Println("Loaded Pattern " + dp.PatternName)
}

func (dp *DynamicPattern) Gen(mapping []pkg.Pixel, f *frame.F) []pkg.Pixel {

	if dp.Loaded {

		res, err := dp.v8ctx.Eval(`pattern.render();`, "demo.js")
		if err != nil {
			fmt.Println(err)
			//panic(err)
		}
		_ = res
		converted := dp.ineffecientBytesToFloat64Slice(res.Bytes(), len(mapping)*3)
		for i, _ := range mapping {
			dp.outputBuffer[i].Color[0] = converted[i*3]
			dp.outputBuffer[i].Color[1] = converted[i*3+1]
			dp.outputBuffer[i].Color[2] = converted[i*3+2]
		}

	}

	return dp.outputBuffer
}

//TODO: REWRITE THIS SHIT FUNCTION INTO SOMETHING THAT ISNT SHIT

func (dp *DynamicPattern) ineffecientBytesToFloat64Slice(bytes []byte, length int) []float64 {
	for i := 0; i < length; i++ {
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
