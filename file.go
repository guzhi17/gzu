// ------------------
// User: pei
// DateTime: 2020/2/24 8:10
// Description:
// ------------------

package gzu

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)


var(
	ErrDone = errors.New("done")
)


func FileLines(fn string, handler func(string)) (err error) {
	fp, err := os.Open(fn)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(fp)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimRight(line, "\n\r")
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func IsTypeOfLower(t string, ts ...string)int64{
	for i, tsi := range ts{
		if strings.Contains(tsi, strings.ToLower(t)){
			return 1 << i
		}
	}
	return 0
}

func Ext(path string) string {
	l:=len(path) - 1
	for i := l; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			if i == l{return ""}
			return path[i+1:]
		}
	}
	return ""
}
