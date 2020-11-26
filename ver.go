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
	Ver20191201 = 568298029518422017
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

//0-32767, 0-65535, 0-2147483647
func VerToInt64(ver string)(int64, error){
	vers := strings.Split(ver, ".")
	if len(vers) != 3{
		return 0, ErrorVersion{Val: ver}
	}

	//v0
	v0, err := strconv.Atoi(vers[0])
	if err != nil {
		return 0, err
	}
	if v0 < 0 || v0 > 32767{
		return 0, ErrorVersion{Val: ver}
	}

	//v1
	v1, err := strconv.Atoi(vers[1])
	if err != nil {
		return 0, err
	}
	if v1 < 0 || v1 > 65535{
		return 0, ErrorVersion{Val: ver}
	}


	//v2
	v2, err := strconv.ParseInt(vers[2], 10, 0)
	if err != nil {
		return 0, err
	}
	if v2 < 0 || v2 > 2147483647{
		return 0, ErrorVersion{Val: ver}
	}

	return int64(v0) << 48 | int64(v1) << 32 | v2, nil
}

func VerFromInt64(v int64) string {
	return fmt.Sprintf("%d.%d.%d", (v >> 48)&0xffff, (v >> 32)&0xffff, v &0xffffffff)
}

func VerIsBiggerThan(ver string, v int64) (bool, error) {
	vc, err := VerToInt64(ver)
	if err != nil{
		return false, err
	}
	return vc > v, nil
}