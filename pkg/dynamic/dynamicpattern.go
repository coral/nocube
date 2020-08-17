package dynamic

import (
	"fmt"
	"io/ioutil"

	v8 "github.com/augustoroman/v8"
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
}

func CreatePattern(Name string, path string) *DynamicPattern {
	return &DynamicPattern{
		PatternName: Name,
		path:        path,
	}
}
func (dp *DynamicPattern) Load(s *v8.Snapshot) {
	code, err := ioutil.ReadFile(dp.path)
	if err != nil {
		panic(err)
	}

	dp.v8isolate = v8.NewIsolateWithSnapshot(s)
	dp.v8ctx = dp.v8isolate.NewContext()
	dp.v8ctx.Eval(string(code), dp.PatternName)
	dp.Loaded = true
	fmt.Println("Loaded Pattern " + dp.PatternName)
}

func (dp *DynamicPattern) Gen(pixels []pkg.Pixel, f *frame.F) []pkg.Pixel {
	if dp.Loaded {
		res, _ := dp.v8ctx.Eval(`render()`, "demo.js")
		fmt.Println("snapshotdemo =", res.String())
	}

	return nil
}

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
