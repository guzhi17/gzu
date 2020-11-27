// ------------------
// User: pei
// DateTime: 2019/10/16 10:10
// Description: 
// ------------------

package gzu

import (
	"math"
	"sync/atomic"
	"time"
	"unsafe"
)

// Int32 is an atomic wrapper around an int32.
type AtomInt32 int32

// Load atomically loads the wrapped value.
func (i *AtomInt32) Load() int32 {
	return atomic.LoadInt32((*int32)(i))
}

// Add atomically adds to the wrapped int32 and returns the new value.
func (i *AtomInt32) Add(n int32) int32 {
	return atomic.AddInt32((*int32)(i), n)
}

// Sub atomically subtracts from the wrapped int32 and returns the new value.
func (i *AtomInt32) Sub(n int32) int32 {
	return atomic.AddInt32((*int32)(i), -n)
}

// Inc atomically increments the wrapped int32 and returns the new value.
func (i *AtomInt32) Inc() int32 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int32 and returns the new value.
func (i *AtomInt32) Dec() int32 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomInt32) CAS(old, new int32) bool {
	return atomic.CompareAndSwapInt32((*int32)(i), old, new)
}

// Store atomically stores the passed value.
func (i *AtomInt32) Store(n int32) {
	atomic.StoreInt32((*int32)(i), n)
}

// Swap atomically swaps the wrapped int32 and returns the old value.
func (i *AtomInt32) Swap(n int32) int32 {
	return atomic.SwapInt32((*int32)(i), n)
}

// AtomInt64 is an atomic wrapper around an int64.
type AtomInt64 int64

// Load atomically loads the wrapped value.
func (i *AtomInt64) Load() int64 {
	return atomic.LoadInt64((*int64)(i))
}

// Add atomically adds to the wrapped int64 and returns the new value.
func (i *AtomInt64) Add(n int64) int64 {
	return atomic.AddInt64((*int64)(i), n)
}

// Sub atomically subtracts from the wrapped int64 and returns the new value.
func (i *AtomInt64) Sub(n int64) int64 {
	return atomic.AddInt64((*int64)(i), -n)
}

// Inc atomically increments the wrapped int64 and returns the new value.
func (i *AtomInt64) Inc() int64 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int64 and returns the new value.
func (i *AtomInt64) Dec() int64 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomInt64) CAS(old, new int64) bool {
	return atomic.CompareAndSwapInt64((*int64)(i), old, new)
}

// Store atomically stores the passed value.
func (i *AtomInt64) Store(n int64) {
	atomic.StoreInt64((*int64)(i), n)
}

// Swap atomically swaps the wrapped int64 and returns the old value.
func (i *AtomInt64) Swap(n int64) int64 {
	return atomic.SwapInt64((*int64)(i), n)
}

// AtomUint32 is an atomic wrapper around an uint32.
type AtomUint32  uint32


// Load atomically loads the wrapped value.
func (i *AtomUint32) Load() uint32 {
	return atomic.LoadUint32((*uint32)(i))
}

// Add atomically adds to the wrapped uint32 and returns the new value.
func (i *AtomUint32) Add(n uint32) uint32 {
	return atomic.AddUint32((*uint32)(i), n)
}

// Sub atomically subtracts from the wrapped uint32 and returns the new value.
func (i *AtomUint32) Sub(n uint32) uint32 {
	return atomic.AddUint32((*uint32)(i), ^(n - 1))
}

// Inc atomically increments the wrapped uint32 and returns the new value.
func (i *AtomUint32) Inc() uint32 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int32 and returns the new value.
func (i *AtomUint32) Dec() uint32 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomUint32) CAS(old, new uint32) bool {
	return atomic.CompareAndSwapUint32((*uint32)(i), old, new)
}

// Store atomically stores the passed value.
func (i *AtomUint32) Store(n uint32) {
	atomic.StoreUint32((*uint32)(i), n)
}

// Swap atomically swaps the wrapped uint32 and returns the old value.
func (i *AtomUint32) Swap(n uint32) uint32 {
	return atomic.SwapUint32((*uint32)(i), n)
}

// AtomUint64 is an atomic wrapper around a uint64.
type AtomUint64 uint64


// Load atomically loads the wrapped value.
func (i *AtomUint64) Load() uint64 {
	return atomic.LoadUint64((*uint64)(i))
}

// Add atomically adds to the wrapped uint64 and returns the new value.
func (i *AtomUint64) Add(n uint64) uint64 {
	return atomic.AddUint64((*uint64)(i), n)
}

// Sub atomically subtracts from the wrapped uint64 and returns the new value.
func (i *AtomUint64) Sub(n uint64) uint64 {
	return atomic.AddUint64((*uint64)(i), ^(n - 1))
}

// Inc atomically increments the wrapped uint64 and returns the new value.
func (i *AtomUint64) Inc() uint64 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped uint64 and returns the new value.
func (i *AtomUint64) Dec() uint64 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomUint64) CAS(old, new uint64) bool {
	return atomic.CompareAndSwapUint64((*uint64)(i), old, new)
}

// Store atomically stores the passed value.
func (i *AtomUint64) Store(n uint64) {
	atomic.StoreUint64((*uint64)(i), n)
}

// Swap atomically swaps the wrapped uint64 and returns the old value.
func (i *AtomUint64) Swap(n uint64) uint64 {
	return atomic.SwapUint64((*uint64)(i), n)
}

// AtomBool is an atomic Boolean.
type AtomBool uint32


// Load atomically loads the Boolean.
func (b *AtomBool) Load() bool {
	return truthy(atomic.LoadUint32((*uint32)(b)))
}

// CAS is an atomic compare-and-swap.
func (b *AtomBool) CAS(old, new bool) bool {
	return atomic.CompareAndSwapUint32((*uint32)(b), boolToInt(old), boolToInt(new))
}

// Store atomically stores the passed value.
func (b *AtomBool) Store(new bool) {
	atomic.StoreUint32((*uint32)(b), boolToInt(new))
}

// Swap sets the given value and returns the previous value.
func (b *AtomBool) Swap(new bool) bool {
	return truthy(atomic.SwapUint32((*uint32)(b), boolToInt(new)))
}

// Toggle atomically negates the Boolean and returns the previous value.
func (b *AtomBool) Toggle() bool {
	return truthy(atomic.AddUint32((*uint32)(b), 1) - 1)
}

func truthy(n uint32) bool {
	return n&1 == 1
}

func boolToInt(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// AtomFloat64 is an atomic wrapper around float64.
type AtomFloat64 uint64


// Load atomically loads the wrapped value.
func (f *AtomFloat64) Load() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(f)))
}

// Store atomically stores the passed value.
func (f *AtomFloat64) Store(s float64) {
	atomic.StoreUint64((*uint64)(f), math.Float64bits(s))
}

// Add atomically adds to the wrapped float64 and returns the new value.
func (f *AtomFloat64) Add(s float64) float64 {
	for {
		old := f.Load()
		cur := old + s
		if f.CAS(old, cur) {
			return cur
		}
	}
}

// Sub atomically subtracts from the wrapped float64 and returns the new value.
func (f *AtomFloat64) Sub(s float64) float64 {
	return f.Add(-s)
}

// CAS is an atomic compare-and-swap.
func (f *AtomFloat64) CAS(old, new float64) bool {
	return atomic.CompareAndSwapUint64((*uint64)(f), math.Float64bits(old), math.Float64bits(new))
}

// Duration is an atomic wrapper around time.Duration
// https://godoc.org/time#Duration
type Duration struct {
	v AtomInt64
}

// NewDuration creates a Duration.
func NewDuration(d time.Duration) *Duration {
	return &Duration{v: AtomInt64(d)}
}

// Load atomically loads the wrapped value.
func (d *Duration) Load() time.Duration {
	return time.Duration(d.v.Load())
}

// Store atomically stores the passed value.
func (d *Duration) Store(n time.Duration) {
	d.v.Store(int64(n))
}

// Add atomically adds to the wrapped time.Duration and returns the new value.
func (d *Duration) Add(n time.Duration) time.Duration {
	return time.Duration(d.v.Add(int64(n)))
}

// Sub atomically subtracts from the wrapped time.Duration and returns the new value.
func (d *Duration) Sub(n time.Duration) time.Duration {
	return time.Duration(d.v.Sub(int64(n)))
}

// Swap atomically swaps the wrapped time.Duration and returns the old value.
func (d *Duration) Swap(n time.Duration) time.Duration {
	return time.Duration(d.v.Swap(int64(n)))
}

// CAS is an atomic compare-and-swap.
func (d *Duration) CAS(old, new time.Duration) bool {
	return d.v.CAS(int64(old), int64(new))
}

// Value shadows the type of the same name from sync/atomic
// https://godoc.org/sync/atomic#Value
type Value struct{ atomic.Value }



type AtomAny struct {
	_d unsafe.Pointer
}

func (a *AtomAny)Get() interface{} {
	ptr := atomic.LoadPointer(&a._d)
	if ptr == nil{
		return nil
	}
	return *(*interface{})(ptr)
}
func (a *AtomAny)Set(v interface{})  {
	//vi := v
	atomic.StorePointer(&a._d, unsafe.Pointer(&v))
}

func (a *AtomAny)CAS(old *AtomAny, new interface{}) bool {
	return atomic.CompareAndSwapPointer(&a._d, old._d, unsafe.Pointer(&new))
}


//this once not used to wait for initialization, it is use to close something once
type Once struct {
	done AtomInt32
}
func (o *Once) Do(f func()) {
	if o.done.Inc() == 1{
		f()
	}
}
func (o *Once) IsDone()bool{
	return 0!=o.done.Load()
}

func (o *Once) Reset(){
	o.done.Store(0)
}
