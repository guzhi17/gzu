// ------------------
// User: pei
// DateTime: 2020/5/27 10:34
// Description:
// ------------------

package gzu

import (
	"os"
	"path/filepath"
)



func ProcessFullPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func ProcessFullName() (string, error) {
	return filepath.Abs(os.Args[0])
}
