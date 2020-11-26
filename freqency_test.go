// ------------------
// User: pei
// DateTime: 2020/9/29 15:37
// Description: 
// ------------------

package gzu

import (
	"log"
	"testing"
	"time"
)

func TestFrequency_Call(t *testing.T) {
	var x = NewFrequency(time.Second/3)

	for i:=0;i<100;i++{
		x.Call(func() {
			log.Println(i)
		})
	}
	time.Sleep(time.Second/2)
	for i:=0;i<100;i++{
		x.Call(func() {
			log.Println(i)
		})
	}
}