package v1

import (
	"os"
	"path"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../..") // change to suit test file location
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
