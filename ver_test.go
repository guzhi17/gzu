// ------------------
// User: pei
// DateTime: 2019/12/2 10:22
// Description: 
// ------------------

package gzu

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestVerToInt64(t *testing.T) {
	log.SetFlags(11)

	log.Println(VerToInt64("1.3.2"))
	log.Println(VerFromInt64(281487861612546))


	log.Println(VerToInt64("2019.12.1"))
	log.Println(VerFromInt64(Ver20191201))
	log.Println(VerFromInt64(568297982273782972))
}

func TestIsVerBiggerThan(t *testing.T) {
	log.Println(strconv.FormatFloat(float64(12121212213)/100, 'f', 1, 64))
	log.Println(strconv.FormatFloat(float64(12121212213)/100, 'f', 2, 64))
}

func TestBasicTypeName(t *testing.T) {
	type MyObj struct {
		S string
		B []byte
		I int
	}
	o := MyObj{S:"hello", B:[]byte("world"), I: 17}
	b, err := json.Marshal(o)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(o.B))
	fmt.Println(string(b))
	fmt.Println(base64.StdEncoding.EncodeToString(b))
	var x struct {
		B string
	}
	err = json.Unmarshal([]byte(`{"B":"\b\u0001\u0002\u0003\u0004\u0005\u0006\u0007\b"}`), &x)
	fmt.Println(x)



	///  \\u0001\\u0002\\u0003\\u0004\\u0005\\u0006\\u0007\\b
}