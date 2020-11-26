// ------------------
// User: pei
// DateTime: 2020/5/6 19:09
// Description: 
// ------------------

package gzu

import (
	"log"
	"testing"
	"time"
)

func TestNewTimeZone(t *testing.T) {
	z := NewTimeZone(8)
	t.Log(z.YearZeroMs(time.Now()))
	t.Log(z.NowSDayZeroToEndTTL(0))
}

func TestTimeZone_StartOf(t *testing.T) {
	now := time.Now()
	z := NewTimeZone(8)

	var xs= []PeriodTypes{PeriodTypeMinute,
		PeriodTypeHour,
		PeriodTypeDay,
		PeriodTypeWeek,
		PeriodTypeMonth,
		PeriodTypeYear,
	}

	for _, p := range xs{
		t.Log(p, z.StartOf(now, p, 0))
	}

	//t.Log(PeriodTypeDay, z.StartOf(now, PeriodTypeDay, 0))
	//t.Log(PeriodTypeWeek, z.StartOf(now, PeriodTypeWeek, 0))
	//t.Log(PeriodTypeMonth, z.StartOf(now, PeriodTypeMonth, 0))
	//t.Log(PeriodTypeYear, z.StartOf(now, PeriodTypeYear, 0))
	t.Log(PeriodTypeHour, z.StartOf(now, PeriodTypeHour, 2))
	t.Log(PeriodTypeMinute, z.StartOf(now, PeriodTypeMinute, 2))
}

func TestTimeZone_ToDayEnd(t *testing.T) {
	z := NewTimeZone(8)
	now := time.Now()
	t.Log(z.ToDayEnd(now, 0))
	t.Log(z.ToDayEnd(now, time.Second))
}

func TestParseDuration(t *testing.T) {
	t.Log(ParseDuration("36500d"))
	t.Log(ParseDuration("365d"))
}

func TestTimeZone_StartOf2(t *testing.T) {
	z := NewTimeZone(8)
	now := time.Now()
	tm := z.StartOf(now, PeriodTypeWeek, 1)
	t.Log(tm)
	t.Log(now.ISOWeek())
	t.Log(tm.ISOWeek())
	t.Log(tm.Weekday())
}

func TestTimeDayOf(t *testing.T) {
	now := time.Now()
	log.Println(TimeDayOf(now))
	log.Println(TimeDayOf(now.Add(time.Hour*24)))
}