package atom

import (
	"sync/atomic"
	"unsafe"
)

type Any struct {
	_d unsafe.Pointer
}

func (a *Any)Get() interface{} {
	ptr := atomic.LoadPointer(&a._d)
	if ptr == nil{
		return nil
	}
	return *(*interface{})(ptr)
}
func (a *Any)Set(v interface{})  {
	//vi := v
	atomic.StorePointer(&a._d, unsafe.Pointer(&v))
}

func (a *Any)CAS(old *Any, new interface{}) bool {
	return atomic.CompareAndSwapPointer(&a._d, old._d, unsafe.Pointer(&new))
}
