package gzu

import (
	"bytes"
	"math"
)

func AbsInt(x int) int  {
	if x < 0{
		return -x
	}
	return x
}

func AbsInt32(x int32) int32  {
	if x < 0{
		return -x
	}
	return x
}


func AbsIntX(n int64, x uint) int64 {
	y := n >> (x-1)
	return (n ^ y) - y
}

func AbsInt64(x int64) int64  {
	if x < 0{
		return -x
	}
	return x
}

func AbsInt16(x int16) int16  {
	if x < 0{
		return -x
	}
	return x
}

func AbsInt8(x int8) int8  {
	if x < 0{
		return -x
	}
	return x
}

func MinInt(v []int) (m int) {
	if len(v)<1{
		return 0
	}
	m=v[0]
	for i:=1;i<len(v);i++{
		if v[i] < m{
			m = v[i]
		}
	}
	return m
}

func Min2Int(a, b int) (int) {
	if a < b{
		return a
	}
	return b
}

func Min2Int64(a, b int64) (int64) {
	if a < b{
		return a
	}
	return b
}

func Max2Int64(a, b int64) (int64) {
	if a < b{
		return b
	}
	return a
}


func Max2Int(a, b int) (int) {
	if a < b{
		return b
	}
	return a
}

func Max2Float64(a, b float64) (float64) {
	if a < b{
		return b
	}
	return a
}

func MaxInt(v []int) (m int) {
	if len(v)<1{
		return 0
	}
	m=v[0]
	for i:=1;i<len(v);i++{
		if m < v[i]{
			m = v[i]
		}
	}
	return m
}

func Min2Bytes(a, b []byte) []byte {
	if bytes.Compare(a, b) < 0{
		return a
	}
	return b
}

func Max2Bytes(a, b []byte) []byte {
	if bytes.Compare(a, b) < 0{
		return b
	}
	return a
}

func MinInt32(v []int32) (m int32) {
	if len(v)<1{
		return 0
	}
	m=v[0]
	for i:=1;i<len(v);i++{
		if v[i] < m{
			m = v[i]
		}
	}
	return m
}

func MaxInt32(v []int32) (m int32) {
	if len(v)<1{
		return 0
	}
	m=v[0]
	for i:=1;i<len(v);i++{
		if m < v[i]{
			m = v[i]
		}
	}
	return m
}


func MaxInt64(v []int64) (m int64) {
	if len(v)<1{
		return 0
	}
	m=v[0]
	for i:=1;i<len(v);i++{
		if m < v[i]{
			m = v[i]
		}
	}
	return m
}
func SumInt64(v []int64)(m int64)  {
	for i:=0;i<len(v) ;i++  {
		m+=v[i]
	}
	return m
}

func SumInt32(v []int32)(m int32)  {
	for i:=0;i<len(v) ;i++  {
		m+=v[i]
	}
	return m
}

func SumInt(v []int)(m int)  {
	for i:=0;i<len(v) ;i++  {
		m+=v[i]
	}
	return m
}

func Max2In64(a, b int64)(int64){
	if a > b {
		return a
	}
	return b
}


func ClipInt(v, a, b int) int {
	if v < a{
		return a
	}
	if v > b{
		return b
	}
	return v
}
func ClipInt32(v, a, b int32) int32 {
	if v < a{
		return a
	}
	if v > b{
		return b
	}
	return v
}
func ClipInt64(v, a, b int64) int64 {
	if v < a{
		return a
	}
	if v > b{
		return b
	}
	return v
}

func ClipFloat64(v, a, b float64) float64 {
	if v < a{
		return a
	}
	if v > b{
		return b
	}
	return v
}

func ClipBytes(v, a, b []byte)  []byte {
	if bytes.Compare(v, a) < 0{
		return a
	}
	if bytes.Compare(b, v) < 0{
		return b
	}
	return v
}

func IsInRangeInt64(v, a, b int64) bool  {
	return v>=a && a<=b
}

func PowerUint64(a uint64, n int) (ret uint64) {
	ret = 1 //
	for n != 0 {
		if n%2 != 0 {
			ret = ret * a
		}
		n /= 2
		a = a * a
	}
	return ret
}

func FloatEqualA(a, b float64) bool {
	return math.Abs(a-b) < math.SmallestNonzeroFloat64
	//math.SmallestNonzeroFloat64
}
func FloatEqual(a, b, e float64) bool {
	return math.Abs(a-b) < e
}


const(
	EarthDiameterM 	= 12756274
	RadFactor 		= 0.0174532925
)

type Location struct {
	Latitude             float64
	Longitude            float64
}

func DistanceRadian(a Location, b Location) float64  {
	lat1, lat2 := a.Latitude * RadFactor, b.Latitude * RadFactor
	dlat, dlon := lat1-lat2, (a.Longitude - b.Longitude)*RadFactor
	return math.Asin(math.Sqrt(math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)))
}

func DistanceM(a Location, b Location) float64 {
	return EarthDiameterM*DistanceRadian(a, b)
}

//
//rad(Angle)->
//Angle*0.01745329. %math:pi()/180 = 0.017453292519943295.
//
//diameter()->12756274. %2*6378137
//% 12756274 = 2*6378137
//%L = 2R*arcsin(sqrt(sin^2((WA-WB)/2) + cos(WA)*cos(WB)*sin^2((JA-JB)/2))
//distance_sphere({LatA, LonA}, {LatB, LonB})->
//{Lat1, Lat2} = {rad(LatA), rad(LatB)},
//DLat = Lat1 - Lat2,
//DLon = rad(LonA) - rad(LonB),
//math:asin(math:sqrt(math:pow(math:sin(DLat/2), 2) + math:cos(Lat1)*math:cos(Lat2)*math:pow(math:sin(DLon/2), 2))).