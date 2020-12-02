// ------------------
// User: pei
// DateTime: 2020/4/8 14:34
// Description: 
// ------------------

package gzu

import (
	"fmt"
	"math"
	"testing"
)

//utils.NewRandomStringGenerator(62, 62^3, 62^7)

var (
	gRandStringGen = NewRandomStringGeneratorLen(62, 3, 7)
)

func TestNewRandomStringGenerator(t *testing.T) {
	t.Log(fmt.Sprintf("BAT%s", gRandStringGen.Gen()))
	t.Log(fmt.Sprintf("BAT%s", gRandStringGen.Gen()))
	t.Log(fmt.Sprintf("BAT%s", gRandStringGen.Gen()))
	t.Log(fmt.Sprintf("BAT%s", gRandStringGen.Gen()))
	t.Log(fmt.Sprintf("BAT%s", gRandStringGen.Gen()))
	t.Log(fmt.Sprintf("BAT%s", gRandStringGen.Gen()))
}

func TestRandInt32(t *testing.T) {
	var (
		r = NewRandomStringGenerator(64, math.MaxUint32, math.MaxUint64)
	)
	t.Log(r.Gen())
	t.Log(r.Gen())
	t.Log(r.Gen())
	t.Log(r.Gen())
	t.Log(r.Gen())
	t.Log(r.Gen())
	t.Log(r.Gen())
	t.Log(r.Gen())
}