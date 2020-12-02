package gzu

import (
	"math/big"
	"math/rand"
	"time"
)

var random *rand.Rand

func init(){
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt32(l, r int32) int32 {
	if r > l{
		return l + (int32)(random.Int31n(r-l))
	}else if r == l{
		return r
	}
	return r + (int32)(random.Int31n(l-r))
}

func RandInt64(l, r int64) int64 {
	if r > l{
		return l + (random.Int63n(r-l))
	}else if r == l{
		return r
	}
	return r + (random.Int63n(l-r))
}

func RandInt63() int64 {
	return random.Int63()
}

//%(0, 1<<63].
func RandInt63NoZero() int64 {
	return random.Int63()+1
}


func RandUint64(l, r uint64) uint64 {
	if r > l{
		return l + uint64(random.Int63n(int64(r-l)))
	}else if r == l{
		return r
	}
	return r + uint64(random.Int63n(int64(l-r)))
}

func RandFloat32()float32{
	return random.Float32()
}

func RandFloat64()float64{
	return random.Float64()
}

func RandInt(l, r int) int {
	if r > l{
		return l + (random.Intn(r-l))
	}else if r == l{
		return r
	}
	return r + (random.Intn(l-r))
}



const Letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Numbers = "0123456789"
var(
	LenLetters = len(Letters)
	LenNumbers = len(Numbers)
)

func RandString(n int) []byte{
	var buffer = make([]byte, n)
	for i := 0; i < n; i++ {
		buffer[i] = Letters[random.Intn(LenLetters)]
	}
	return buffer
}

func RandNumberString(n int) []byte{
	var buffer = make([]byte, n)
	for i := 0; i < n; i++ {
		buffer[i] = Numbers[random.Intn(LenNumbers)]
	}
	return buffer
}

type RandomStringGenerator struct {
	base int //2-62
	min, max uint64
	random *rand.Rand
}

func Uint64n(r *rand.Rand, n uint64) uint64 {
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Uint64() & (n - 1)
	}
	max := 0xFFFFFFFFFFFFFFFF - 0xFFFFFFFFFFFFFFFF%n //uint64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Uint64()
	for v > max {
		v = r.Uint64()
	}
	return v % n
}

func randUint64(rd *rand.Rand, l, r uint64) uint64 {
	if r > l{
		return l + Uint64n(rd, r-l)
		//return l + uint64(rd.Int63n(int64(r-l)))
	}else if r == l{
		return r
	}
	return l + Uint64n(rd, l-r) //r + uint64(rd.Int63n(int64(l-r)))
}

func NewRandomStringGenerator(base int, min, max uint64) *RandomStringGenerator {
	return &RandomStringGenerator{
		base:ClipInt(base, 2, 62),
		min: min,
		max: max,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func NewRandomStringGeneratorLen(base int, min, max int) *RandomStringGenerator {
	base = ClipInt(base, 2, 62)
	return NewRandomStringGenerator(base, PowerUint64(uint64(base), min), PowerUint64(uint64(base), max))
}

func (r RandomStringGenerator)Gen() string {
	return (&big.Int{}).SetUint64(randUint64(r.random, r.min, r.max)).Text(r.base)
}