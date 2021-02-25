// ------------------
// User: pei
// DateTime: 2020/4/8 14:43
// Description: 
// ------------------

package gzu

import (
	"log"
	"math"
	"testing"
)

func TestPowerUint64(t *testing.T) {
	t.Log(PowerUint64(3, 5))
	t.Log(PowerUint64(2, 3))
	t.Log(PowerUint64(5, 3))
}

func TestFloatEqual(t *testing.T) {
	t.Log(FloatEqual(0.001, 0, 0.0001))
	t.Log(FloatEqual(0.001, 0, 0.001))
	t.Log(FloatEqual(0.001, 0, 0.01))
}

func TestDistanceM(t *testing.T) {
	t.Log(DistanceM(Location{
		Latitude:             30.691106,
		Longitude:            103.98507,
	}, Location{
		Latitude:             30.634942,
		Longitude:            104.146622,
	}))
}

func TestR3OfLocation(t *testing.T) {

	//
	var ls = []Location{
		{0, 0}, //1,0,0
		{0, -90}, //0, -1, 0
		{0, 90}, //0, 1, 0
		{0, 180}, //-1, 0, 0
		{-90, 0}, //
		{-90, -90},
		{-90, 90},
		{-90, 180},
		{90, 0},
		{90, -90},
		{90, 90},
		{90, 180},
	}

	fract := math.Pi/180
	log.Println(fract)
	for _, li := range ls{
		log.Println(R3OfLocation(li.Latitude, li.Longitude))
	}
}