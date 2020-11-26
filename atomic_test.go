package gzu

import (
	"fmt"
	"log"
	"testing"
)


type I interface {
	P()
}

type S struct {
	v int
}

func (s *S)P()  {
	fmt.Println("i am v: ", s.v)
}

func Get()(r I)  {
	return nil
}

func TestAny(t *testing.T)  {
	a := Any{}

	s, ok := a.Get().([]byte)
	fmt.Println(s, ok)

	a.Set([]byte("hello world"))
	s, ok = a.Get().([]byte)

	fmt.Println(string(s), ok)


	a.Set(&S{v: 6})
	x, ok := a.Get().(I)
	if ok{
		x.P()
	}
	xx := Get()
	fmt.Println( "xx==nil: ", xx == nil)

}

func TestAny_CAS(t *testing.T) {
	a := Any{}
	a.Set(&S{v:0})

	b:=a

	a.Set(&S{v:1})
	t.Log(b.Get())

	s:=&S{v:7}
	t.Log(a.CAS(&b, s))
	var ok bool
	s, ok = a.Get().(*S)
	t.Log(s, ok)
}

func TestCas(t *testing.T)  {
	var idx Int32

	log.Println(idx.Swap(1))
	log.Println(idx.Swap(1))
	log.Println(idx.Swap(0))
	log.Println(idx.Swap(0))
}