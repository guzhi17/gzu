// ------------------
// User: pei
// DateTime: 2020/9/29 15:30
// Description: 
// ------------------

package gzu

import (
	"time"
)

type Frequency struct {
	dt int64 //in ms
	next Int64 //last time in ms
}

func NewFrequency(dt time.Duration) *Frequency {
	return &Frequency{dt:dt.Milliseconds()}
}

func (s *Frequency)Call(f func())  {
	if s.dt < 1{
		f()
		return
	}
	now := Now()
	v := s.next.Load()
	if v > now{return}
	if s.next.CAS(v, now+s.dt){
		f()
	}
}