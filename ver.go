// ------------------
// User: pei
// DateTime: 2019/12/2 10:07
// Description: 
// ------------------

package gzu

import (
	"fmt"
	"strconv"
	"strings"
)

// "2019.12.1" -> 568298029518422017
const(
	Ver20191201 Version = 568298029518422017
)

type ErrorVersion struct {
	Val string
}
var _ error = ErrorVersion{}
func (e ErrorVersion)Error() string {
	var b strings.Builder
	b.WriteString("not version format: ")
	b.WriteString(e.Val)
	return b.String()
}


type Version uint64

func (v Version) Major() uint64  {
	return (uint64(v) >> 48)&0xffff //, (v >> 32)&0xffff, v &0xffffffff)
}
func (v Version) Minor() uint64  {
	return (uint64(v) >> 32)&0xffff //, v &0xffffffff)
}
func (v Version) Build() uint64  {
	return uint64(v) &0xffffffff //, v &0xffffffff)
}

//0-32767, 0-65535, 0-2147483647
func VerToInt64(ver string)(Version, error){
	ver = strings.TrimLeft(ver, "vV")
	vers := strings.Split(ver, ".")
	v, err := verToInt64(vers)
	if err != nil{
		vv, err := strconv.ParseUint(ver, 0, 0)
		if err != nil{
			return 0, err
		}
		return Version(vv), nil
	}
	return v, err
	//
	//
	//if len(vers) != 3{
	//	vv, err := strconv.ParseUint(ver, 0, 0)
	//	if err != nil{
	//		return 0, err
	//	}
	//	return Version(vv), nil
	//}
	//
	////v0
	//v0, err := strconv.Atoi(vers[0])
	//if err != nil {
	//	return 0, err
	//}
	//if v0 < 0 || v0 > 32767{
	//	return 0, ErrorVersion{Val: ver}
	//}
	//
	////v1
	//v1, err := strconv.Atoi(vers[1])
	//if err != nil {
	//	return 0, err
	//}
	//if v1 < 0 || v1 > 65535{
	//	return 0, ErrorVersion{Val: ver}
	//}
	//
	////v2
	//v2, err := strconv.ParseUint(vers[2], 10, 0)
	//if err != nil {
	//	return 0, err
	//}
	//if v2 < 0 || v2 > 2147483647{
	//	return 0, ErrorVersion{Val: ver}
	//}
	//
	//return Version(uint64(v0) << 48 | uint64(v1) << 32 | v2), nil
}


func verToInt64(vers []string)(Version, error){
	//v0
	v0, err := strconv.Atoi(vers[0])
	if err != nil {
		return 0, err
	}
	if v0 < 0 || v0 > 32767{
		return 0, ErrorVersion{}
	}

	//v1
	v1, err := strconv.Atoi(vers[1])
	if err != nil {
		return 0, err
	}
	if v1 < 0 || v1 > 65535{
		return 0, ErrorVersion{}
	}

	var v2, v3 uint64
	switch len(vers) {
	case 3:
		v2, err = strconv.ParseUint(vers[2], 10, 0)
		if err != nil {
			return 0, err
		}
		if v2 < 0 || v2 > 2147483647{
			return 0, ErrorVersion{}
		}
	case 4:
		v3, err = strconv.ParseUint(vers[2], 10, 0)
		if err != nil {
			return 0, err
		}
		if v3 < 0 || v3 > 65535{
			return 0, ErrorVersion{}
		}
		v2, err = strconv.ParseUint(vers[3], 10, 0)
		if err != nil {
			return 0, err
		}
		if v2 < 0 || v2 > 65535{
			return 0, ErrorVersion{}
		}
	}
	return Version(uint64(v0) << 48 | uint64(v1) << 32 | v2 | (v3 << 16)), nil
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", (v >> 48)&0xffff, (v >> 32)&0xffff, v &0xffffffff)
}

func (v Version) String4() string {
	return fmt.Sprintf("v%d.%d.%d.%d", (v >> 48)&0xffff, (v >> 32)&0xffff, (v >> 16)&0xffff, v&0xffff)
}

func VerFromInt64(v Version) string {
	return fmt.Sprintf("%d.%d.%d", (v >> 48)&0xffff, (v >> 32)&0xffff, v &0xffffffff)
}

func VerIsBiggerThan(ver string, v Version) (bool, error) {
	vc, err := VerToInt64(ver)
	if err != nil{
		return false, err
	}
	return vc > v, nil
}