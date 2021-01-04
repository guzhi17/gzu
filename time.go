package gzu

import (
	"errors"
	"time"
)

const(
	TIME_1S_MS		= 1000
	//TIME_5S_MS		= TIME_1S_MS*5//5000
	//TIME_10S_MS		= TIME_1S_MS*10 //10000
	//TIME_20S_MS		= TIME_1S_MS*20 //20000
	//TIME_30S_MS		= TIME_1S_MS*30 //30000
	TIME_1M_MS		= TIME_1S_MS*60 //60000

	TIME_1H_MS		= TIME_1M_MS*60 //

	TIME_1D_MS		= TIME_1H_MS*24 //86400000
	TIME_7D_MS		= TIME_1D_MS*7 //604800000
	TIME_1Y_MS 		= TIME_1D_MS*365

	TIME_1M_SEC		= 60
	//TIME_10M_SEC	= TIME_1M_SEC*10 //600
	TIME_1H_SEC		= TIME_1M_SEC*60 //3600
	//TIME_2H_SEC		= TIME_1H_SEC*2 //7200
	//TIME_6H_SEC		= TIME_1H_SEC*6 //21600
	//TIME_10H_SEC	= TIME_1H_SEC*10 //36000
	//TIME_12H_SEC	= TIME_1H_SEC*12 //43200
	TIME_1D_SEC		= TIME_1H_SEC*24 //86400
	//TIME_1W_SEC		= TIME_1D_SEC*7 //604800
	TIME_7D_SEC		= TIME_1D_SEC*7 //604800
	TIME_1Y_SEC 	= TIME_1D_SEC*365
)


type TimeZone struct {
	*time.Location
}


func Now() int64 {
	return time.Now().UnixNano() / 1000000
}
func TimeMsOf(t time.Time)int64{
	return t.UnixNano() / 1000000
}
func TimeSOf(t time.Time)int64{
	return t.Unix()
}
func TimeMinOf(t time.Time)int64{
	return t.Unix()/60
}
func TimeHourOf(t time.Time)int64{
	return t.Unix()/3600
}

func TimeDayOf(t time.Time)int64{
	return t.Unix()/(3600*24)
}



func NewTimeZone(n int) *TimeZone  {
	return &TimeZone{
		Location: time.FixedZone("UTC", 60*60*n),
	}
}

//24:00
func (z TimeZone)DateZeroMs(now time.Time, ddays int) ( msday int64) {
	y, m, d := now.Date()
	msday = time.Date(y, m, d+ddays, 0,0,0,0, z.Location).UnixNano() / 1000000
	return
}


func (z TimeZone)YearZeroMs(now time.Time) ( msday int64) {
	y, _, _ := now.Date()
	msday = time.Date(y, 1, 1, 0,0,0,0, z.Location).UnixNano() / 1000000
	return
}
//24:00
func (z TimeZone)NowAndDateZeroMs(ddays int) (ms int64, msday int64) {
	now := time.Now()
	ms = now.UnixNano() / 1000000
	y, m, d := now.Date()
	msday = time.Date(y, m, d+ddays, 0,0,0,0, z.Location).UnixNano() / 1000000
	return
}

func (z TimeZone)NowAndDateZero(ddays int) (ms int64, day int) {
	now := time.Now()
	ms = now.UnixNano() / 1000000
	y, m, d := now.Date()
	day = int(time.Date(y, m, d+ddays, 0,0,0,0, z.Location).Unix() / TIME_1D_SEC)
	return
}

func (z TimeZone)ToDayEnd(tm time.Time, duration time.Duration) time.Duration {
	y, m, d := tm.Date()
	return time.Date(y, m, d+1, 0,0,0,0, z.Location).Add(duration).Sub(tm)
}


func (z TimeZone)NowSDayZeroToEndTTL(ddays int) (s int64, day int, ttl int64) {
	now := time.Now()
	s = now.Unix()
	y, m, d := now.Date()
	ds := time.Date(y, m, d+ddays, 0,0,0,0, z.Location).Unix()

	return s, int(ds/TIME_1D_SEC), ds + TIME_1D_SEC - s
}

type PeriodTypes int
const(
	PeriodTypeNone PeriodTypes = 0x00
	PeriodTypeMinute = 0x01
	PeriodTypeHour = 0x02
	PeriodTypeDay  = 0x03
	PeriodTypeWeek = 0x04
	PeriodTypeMonth = 0x05
	PeriodTypeYear = PeriodTypeNone
)

func (p PeriodTypes)String() string{
	switch p {
	default:
		return "PeriodTypeUnknown"
	case PeriodTypeDay:
		return "PeriodTypeDay"
	case PeriodTypeWeek:
		return "PeriodTypeWeek"
	case PeriodTypeMonth:
		return "PeriodTypeMonth"
	case PeriodTypeYear:
		return "PeriodTypeYear"
	case PeriodTypeHour:
		return "PeriodTypeHour"
	case PeriodTypeMinute:
		return "PeriodTypeMinute"
	}
}

func (z TimeZone)StartOf(t time.Time, p PeriodTypes, dt int) (s time.Time) {
	y, m, d := t.Date()
	switch p {
	default: fallthrough
	case PeriodTypeDay:
		s = time.Date(y, m, d+dt, 0,0,0,0, z.Location)
	case PeriodTypeWeek:
		var w = (int(t.Weekday())+6)%7
		s = time.Date(y, m, d-w+7*dt, 0,0,0,0, z.Location)
	case PeriodTypeMonth:
		s = time.Date(y, m+time.Month(dt), 1, 0,0,0,0, z.Location)
	case PeriodTypeYear:
		s = time.Date(y+dt, 1, 1, 0,0,0,0, z.Location)
	case PeriodTypeHour:
		s = time.Date(y, m, d, dt,0,0,0, z.Location)
	case PeriodTypeMinute:
		s = time.Date(y, m, d, t.Hour(), dt,0,0, z.Location)
	}
	return
}





var errLeadingInt = errors.New("time: bad [0-9]*") // never printed

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x int64, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > (1<<63-1)/10 {
			// overflow
			return 0, "", errLeadingInt
		}
		x = x*10 + int64(c) - '0'
		if x < 0 {
			// overflow
			return 0, "", errLeadingInt
		}
	}
	return x, s[i:], nil
}

// leadingFraction consumes the leading [0-9]* from s.
// It is used only for fractions, so does not return an error on overflow,
// it just stops accumulating precision.
func leadingFraction(s string) (x int64, scale float64, rem string) {
	i := 0
	scale = 1
	overflow := false
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if overflow {
			continue
		}
		if x > (1<<63-1)/10 {
			// It's possible for overflow to give a positive number, so take care.
			overflow = true
			continue
		}
		y := x*10 + int64(c) - '0'
		if y < 0 {
			overflow = true
			continue
		}
		x = y
		scale *= 10
	}
	return x, scale, s[i:]
}

var unitMap = map[string]int64{
	"ns": int64(time.Nanosecond),
	"us": int64(time.Microsecond),
	"µs": int64(time.Microsecond), // U+00B5 = micro symbol
	"μs": int64(time.Microsecond), // U+03BC = Greek letter mu
	"ms": int64(time.Millisecond),
	"s":  int64(time.Second),
	"m":  int64(time.Minute),
	"h":  int64(time.Hour),
	"d":  int64(time.Hour*24),
}

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ParseDuration(s string) (time.Duration, error) {
	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	var d int64
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("time: invalid duration " + orig)
	}
	for s != "" {
		var (
			v, f  int64       // integers before, after decimal point
			scale float64 = 1 // value = v + f/scale
		)

		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return 0, errors.New("time: invalid duration " + orig)
		}
		// Consume [0-9]*
		pl := len(s)
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, errors.New("time: invalid duration " + orig)
		}
		pre := pl != len(s) // whether we consumed anything before a period

		// Consume (\.[0-9]*)?
		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			f, scale, s = leadingFraction(s)
			post = pl != len(s)
		}
		if !pre && !post {
			// no digits (e.g. ".s" or "-.s")
			return 0, errors.New("time: invalid duration " + orig)
		}

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || '0' <= c && c <= '9' {
				break
			}
		}
		if i == 0 {
			return 0, errors.New("time: missing unit in duration " + orig)
		}
		u := s[:i]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, errors.New("time: unknown unit " + u + " in duration " + orig)
		}
		if v > (1<<63-1)/unit {
			// overflow
			return 0, errors.New("time: invalid duration " + orig)
		}
		v *= unit
		if f > 0 {
			// float64 is needed to be nanosecond accurate for fractions of hours.
			// v >= 0 && (f*unit/scale) <= 3.6e+12 (ns/h, h is the largest unit)
			v += int64(float64(f) * (float64(unit) / scale))
			if v < 0 {
				// overflow
				return 0, errors.New("time: invalid duration " + orig)
			}
		}
		d += v
		if d < 0 {
			// overflow
			return 0, errors.New("time: invalid duration " + orig)
		}
	}

	if neg {
		d = -d
	}
	return time.Duration(d), nil
}



const WeekSeconds = 7*24*60*60
func TimeWeeks(t time.Time) int64 {
	return t.Unix() / WeekSeconds
}