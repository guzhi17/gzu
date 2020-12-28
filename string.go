// ------------------
// User: pei
// DateTime: 2020/2/24 10:32
// Description: 
// ------------------

package gzu

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func StringArrayZero(v []string) string {
	if len(v) < 1{
		return ""
	}
	return v[0]
}

func AnyToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
func IntPtrOrDefault(s *int, d int) int {
	if s == nil{return d}
	return *s
}
func Int64PtrOrDefault(s *int64, d int64) int64 {
	if s == nil{return d}
	return *s
}
func IntOrDefault(s, d int) int {
	if s == 0{return d}
	return s
}
func StringOrDefault(s, d string) string {
	if s == ""{return d}
	return s
}

func StringToInt64(s string, d int64) int64 {
	i64, err := strconv.ParseInt(s, 10, 0)
	if err != nil{
		return d
	}
	return i64
}

func BytesOrDefault(s, d []byte) []byte {
	if s == nil{
		return d
	}
	return s
}

func StringByPtrOrDefault(p *string, d string) string {
	if p == nil{
		return d
	}
	return *p
}
func StringByPtr(p *string) string {
	if p == nil{
		return ""
	}
	return *p
}
func StringsByPtr(p *[]string) []string {
	if p == nil{
		return nil
	}
	return *p
}

func StringPtr(s string)(*string)  {
	return &s
}


func StringToBool(s string) bool {
	switch strings.ToLower(s) {
	default:
		return false
	case "1", "true": return true
	}
}


///======================================================================================================================
func GetReg(reg string) (*regexp.Regexp, error) {
	r, err := regexp.Compile(reg)
	if err != nil{
		return nil, err
	}
	return r, nil
}
///======================================================================================================================

func ReplaceEnvironmentFunc() func([]byte) []byte {
	reg, err := GetReg(`\$\{[0-9a-zA-Z_]+\}`)
	if err != nil{
		return nil
	}
	return func(v []byte) []byte {
		return reg.ReplaceAllFunc(v , func(bytes []byte) []byte {
			//n := string(bytes[2:len(bytes)-1])
			env := os.Getenv(string(bytes[2:len(bytes)-1]))
			return []byte(env)
		})
	}
}


func ReplaceEnvironmentMapFunc(src map[string]string) func([]byte) []byte {
	reg, err := GetReg(`\$\{[0-9a-zA-Z_]+\}`)
	if err != nil{
		return nil
	}
	return func(v []byte) []byte {
		return reg.ReplaceAllFunc(v , func(bytes []byte) []byte {
			return []byte(src[string(bytes[2:len(bytes)-1])])
		})
	}
}


func StringSplit(s string, chrs string) (left, right string) {
	var n = strings.IndexAny(s, chrs)
	if n < 0{
		return s, ""
	}
	return s[0:n], s[n+1:]
}

func StringSplitAny(s string, chrs string) (r []string) {
	for {
		var n = strings.IndexAny(s, chrs)
		if n < 0{
			r = append(r, s)
			break
		}else{
			r = append(r, s[:n])
			s = s[n+1:]
		}
	}
	return
}

func StringSplitAnyIgnoreEmpty(s string, chrs string) (r []string) {
	for {
		var n = strings.IndexAny(s, chrs)
		if n < 0{
			if len(s) > 0{
				r = append(r, s)
			}
			break
		}else{
			if n > 0{
				r = append(r, s[:n])
			}
			s = s[n+1:]
		}
	}
	return
}

func StringRemoveAll(s string, chrs string) (r string) {
	for {
		var n = strings.IndexAny(s, chrs)
		if n < 0{
			if len(s) > 0{
				r += s
			}
			break
		}else{
			if n > 0{
				r += s[:n]
			}
			s = s[n+1:]
		}
	}
	return
}

func BytesRemoveAll(s []byte, chrs string) (r []byte) {
	for {
		var n = bytes.IndexAny(s, chrs)
		if n < 0{
			if len(s) > 0{
				r = append(r, s...)
			}
			break
		}else{
			if n > 0{
				r = append(r, s[:n]...)
			}
			s = s[n+1:]
		}
	}
	return
}

func PGStringRemove00(s string) string {
	return StringRemoveAll(s, string([]byte{0}))
}

func PGBytesRemove00(s []byte) []byte {
	return bytes.ReplaceAll(s, []byte{0}, nil)
}

func StringSplitBy(s string, sub string) (a, b string) {
	var n = strings.Index(s, sub)
	if n < 0{
		return s, ""
	}
	return s[:n], s[n+len(sub):]
}

func StringTailAny(s string, chars string) (n int, b string) {
	jl := len(chars)
	for i:=len(s)-1; i >= 0; i--{
		for j:=0;j<jl;j++{
			if s[i] == chars[j]{
				return i, s[i+1:]
			}
		}
	}
	return -1, s
}

func GetRune(s string, n int, as...string) string {
	if n < 0{
		return s
	}
	var body = []rune(s)
	if n >= len(body){
		return s
	}
	if len(as) > 0{
		return string(body[:n]) + as[0]
	}
	return string(body[:n])
}

func StartWith(s string, prefixes ...string) bool  {
	for _, prefix := range prefixes  {
		if strings.HasPrefix(s, prefix){
			return true
		}
	}
	return false
}

func NormalUtf8(s string) string {
	return strings.Replace(s, string([]byte{0}), "", -1)
}

func ShuffleStrings(ss []string) []string {
	cnt := len(ss)
	if cnt < 1{ return nil}
	nss := make([]string, cnt)
	copy(nss, ss)
	if cnt < 2{ return nss}
	for i:=0;i<cnt;i++{
		j := RandInt(0, cnt)
		nss[i], nss[j] = nss[j], nss[i]
	}
	return nss
}

func RepeatN(r string, n int, sep string) string {
	var sb strings.Builder
	for i:=0;i<n;i++{
		if i!= 0{
			sb.WriteString(sep)
		}
		sb.WriteString(r)
	}
	return sb.String()
}

///======================================================================================================================
//const MaxRepeat = 3
var SpaceRunes = []rune("\r\t \n")
func StringContentNormal(s string, maxRepeat int, maxLine int) string {
	//x, _ := regexp.Compile(``)
	var v = []rune(s)
	var l = len(v)
	if l < 2{return s}
	var(
		out strings.Builder
		old rune = v[0]
		n int
	)
	oldspci :=  RuneIndex(SpaceRunes, old)
	for i:=1; i<l; i++{
		if old == v[i]{ continue } //comment this line to let it not do
		if oldspci >= 0{
			vspci := RuneIndex(SpaceRunes, v[i])
			if vspci >=0 {
				if oldspci < vspci{
					old = v[i]
					oldspci = vspci
				}
				continue
			}
		}
		//check is modified?
		if old == '\n'{
			maxLine--
			if maxLine < 0{
				out.WriteByte(' ')
			}else {
				out.WriteRune(old)
			}
		} else{
			out.WriteRune(old)
		}
		if oldspci < 0{
			for j:=0; j< Min2Int(i-n-1, maxRepeat);j++{
				out.WriteRune(old)
			}
		}
		old = v[i]
		oldspci =  RuneIndex([]rune{'\r', '\t', ' ', '\n'}, old)
		n=i
	}
	out.WriteRune(old)
	if oldspci < 0{
		for j:=0; j< Min2Int(l-n-1, maxRepeat);j++{
			out.WriteRune(old)
		}
	}
	return out.String()
}
func RuneIndex(s []rune, v rune) int {
	for i, si := range s{
		if si == v{
			return i
		}
	}
	return -1
}
///======================================================================================================================
///======================================================================================================================
//word, min-word-len, min-count
func RuneCountRepeat(s []rune, wn, mn int) (has bool, start, cnt int, w []rune) {
	l := len(s)
	if wn < 1{ wn = 1 }
	if mn < 2{ mn = 2}
	for wz := l/mn; wz >= wn; wz--{
		for i:=0; i <= l-mn*wz; i++{
			w0 := s[i:i+wz]
			j := 1
			for ; i+(j+1)*wz <= l; j++{
				if !RuneCompare(w0, s[i+j*wz:i+(j+1)*wz]){
					break
				}
			}
			if j >= mn{
				return true, i, j, w0
			}
		}
	}
	return
}

//split to tow set with same order
//除了u4e00-u9fa5 (中文)之外，还有(0x3400, 0x4DB5)也是
func RuneChineseOnly(s []rune) ( r []rune)  {
	r, _ = RuneSplit(s, func(v rune) bool {
		return (v >= 0x4e00 && v<= 0x9fa5) || (v >= 0x3400 && v <= 0x4DB5)
	})
	return r
}
func RuneSplit(s []rune, f func(rune)bool) (a, b []rune) {
	for i:=0; i<len(s); i++{
		if f(s[i]){
			a = append(a, s[i])
		}else{
			b = append(b, s[i])
		}
	}
	return
}

func RuneCompare(a, b []rune) bool {
	if len(a) != len(b){return false}
	for i:=0; i < len(a); i++{
		if a[i] != b[i]{return false}
	}
	return true
}