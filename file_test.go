// ------------------
// User: pei
// DateTime: 2020/2/24 8:26
// Description: 
// ------------------

package gzu

import (
	"log"
	"testing"
)

func TestStringSplit(t *testing.T) {
	t.Log(StringSplit("", "#"))
	t.Log(StringSplit("#abcdefg", "#"))
	t.Log(StringSplit("abcdefg", "#"))
	t.Log(StringSplit("abcdef#g", "#"))
}

func TestExt(t *testing.T) {
	log.Println(Ext("hello.jpg"))
}