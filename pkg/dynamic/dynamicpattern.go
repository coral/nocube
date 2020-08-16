package dynamic

import (
	"fmt"
	"io/ioutil"

	v8 "github.com/augustoroman/v8"
)

type DynamicPattern struct {
	Name      string
	path      string
	code      string
	v8isolate *v8.Isolate
	v8ctx     *v8.Context
	Loaded    bool
}

func CreatePattern(Name string, path string) *DynamicPattern {
	return &DynamicPattern{
		Name: Name,
		path: path,
	}
}
func (dp *DynamicPattern) Load(s *v8.Snapshot) {
	code, err := ioutil.ReadFile(dp.path)
	if err != nil {
		panic(err)
	}

	dp.v8isolate = v8.NewIsolateWithSnapshot(s)
	dp.v8ctx = dp.v8isolate.NewContext()
	dp.v8ctx.Eval(string(code), dp.Name)
	dp.Loaded = true
	fmt.Println("Loaded Pattern " + dp.Name)
}

func (dp *DynamicPattern) Exec() {
	if dp.Loaded {
		res, _ := dp.v8ctx.Eval(`render()`, "demo.js")
		fmt.Println("snapshotdemo =", res.String())
	}
}

func (dp *DynamicPattern) Unload() {
	dp.Loaded = false
	dp.v8isolate.Terminate()
}
