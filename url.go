// ------------------
// User: pei
// DateTime: 2020/2/21 15:53
// Description: 
// ------------------

package gzu

import (
	"fmt"
	"strconv"
	"strings"
)


const(
	FlagNone = 0
	FlagScheme = 1
	FlagUser = 2
	FlagHost = 4
	FlagPort = 8
	FlagPath = 0x10
	FlagRawQuery = 0x20
	FlagFragment = 0x40
)

// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
type Host struct {
	Addr string
	Port int
}

func HostFromString(s string) (h Host) {
	for i := len(s) -1; i >= 0; i--{
		if s[i] == ':'{
			ps := s[i+1:]
			if len(ps) < 1{
				h.Addr = s
				return
			}
			v, err := strconv.Atoi(s[i+1:])
			if err != nil{
				h.Addr = s
				return
			}
			h.Addr = s[:i]
			h.Port = v
			return
		}
	}
	h.Addr = s
	return
}

type URL struct {
	Scheme     string // in lower case?
	//Opaque     string    // encoded opaque data
	User       *UserInfo // username and password information
	Host       Host    // host or host:port
	Path       string    // path (relative paths may omit leading slash)
	//RawPath    string    // encoded path hint (see EscapedPath method)
	//ForceQuery bool      // append a query ('?') even if RawQuery is empty
	RawQuery   string    // encoded query values, without '?'
	Fragment   string    // fragment for references, without '#'
}
func (u *URL)ParseQuery() map[string]string  {
	//todo cache it?, you'd make it private first
	kvs:=strings.Split(u.RawQuery, "&")
	var m = map[string]string{}
	for _, kvi := range kvs{
		kv := strings.SplitN(kvi, "=", 2)
		var v = ""
		if len(kv) > 1{
			v = kv[1]
		}
		m[kv[0]] = v
	}
	return m
}

func (u *URL)UnmarshalQuery(o interface{}) error {
	return MapMapByTag(o, u.ParseQuery(), "json")
}

func (u *URL)GetQuery(k string) (has bool, v string)  {
	rq := u.RawQuery
	var (
		ib = -1
		ie = -1
	)
	i:=0
	for ;i<len(rq);i++{
		switch rq[i] {
		case '=': ie = i
		case '&':
			var ip = i
			if ie > ib{
				ip = ie
			}
			if strings.EqualFold(k, rq[ib+1:ip]){
				if ip < i{
					return true, rq[ip+1:i]
				}
				return true, ""
			}
			ib = i
		}
	}
	var ip = i
	if ie > ib{
		ip = ie
	}
	if strings.EqualFold(k, rq[ib+1:ip]){
		if ip < i{
			return true, rq[ip+1:i]
		}
		return true, ""
	}
	return false, ""
}

func (u *URL)HasValidUserAndPassword() bool  {
	return u.User!=nil && len(u.User.Username)>0 && len(u.User.Password)>0
}

func (u *URL)GetHostPort(def int) (string, int) {
	if u.Host.Port < 1{
		return u.Host.Addr, def
	}
	return u.Host.Addr, u.Host.Port
	//hp := strings.SplitN(u.Host, ":", 2)
	//if len(hp) < 2{
	//	return hp[0], def
	//}
	//port, err := strconv.Atoi(hp[1])
	//if err != nil{
	//	return hp[0], def
	//}
	//return hp[0], port
}

type UserInfo struct {
	Username    string
	Password    string
	//passwordSet bool
}

func (u *UserInfo)String() string {
	if len(u.Username)>0 || len(u.Password) >0{
		return fmt.Sprintf("%s:%s", u.Username, u.Password)
	}
	return ""
}

func (u *URL)FullPath() string {
	var buf strings.Builder
	if u.Host.Addr != ""{
		buf.WriteString(u.Host.Addr)
	}
	if u.Host.Port > 0{
		buf.WriteByte(':')
		buf.WriteString(strconv.FormatInt(int64(u.Host.Port), 10))
	}
	if u.Path != ""{
		if u.Path[0] != '/'{
			buf.WriteByte('/')
		}
		buf.WriteString(u.Path)
	}else{
		buf.WriteByte('/')
	}
	return buf.String()
}

func (u *URL)FullPathName() string {
	var buf strings.Builder
	if u.Host.Addr != ""{
		buf.WriteString(u.Host.Addr)
	}
	if u.Host.Port > 0{
		buf.WriteByte(':')
		buf.WriteString(strconv.FormatInt(int64(u.Host.Port), 10))
	}
	if u.Path != ""{
		buf.WriteByte('/')
		buf.WriteString(u.Path)
	}
	return buf.String()
}

func (u *URL)Clone()(r *URL){
	r = &URL{
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     u.Path,
		RawQuery: u.RawQuery,
		Fragment: u.Fragment,
	}
	if u.User != nil{
		r.User = &UserInfo{
			Username: u.User.Username,
			Password: u.User.Password,
		}
	}
	return
}

func (u *URL) String() string {
	var buf strings.Builder
	if u.Scheme != "" {
		buf.WriteString(u.Scheme)
		buf.WriteString("://")
	}
	if u.User != nil{
		var us = u.User.String()
		if us != ""{
			buf.WriteString(us)
			buf.WriteByte('@')
		}
	}

	if u.Host.Addr != ""{
		buf.WriteString(u.Host.Addr)
	}
	if u.Host.Port > 0{
		buf.WriteByte(':')
		buf.WriteString(strconv.FormatInt(int64(u.Host.Port), 10))
	}

	if u.Path != ""{
		if u.Path[0] != '/'{
			buf.WriteByte('/')
		}
		buf.WriteString(u.Path)
	}

	if u.RawQuery != ""{
		buf.WriteByte('?')
		buf.WriteString(u.RawQuery)
	}

	if u.Fragment != ""{
		buf.WriteByte('#')
		buf.WriteString(u.Fragment)
	}

	return buf.String()
}


// [scheme:][//[userinfo@]host][/]path[?query][#fragment]

func UrlParse(raw string) *URL {
	var (
		flags uint
		u URL
	)
	//schema
	{
		vs:=strings.SplitN(raw, "://", 2)
		if len(vs) == 2{
			flags |= FlagScheme
			raw = vs[1]
			u.Scheme = vs[0]
		}else{
			//no scheme
			raw = vs[0]
		}
	}
	//then userinfo
	{
		vs:=strings.SplitN(raw, "@", 2)
		if len(vs) == 2{
			flags |= FlagUser
			raw = vs[1]
			//has user info?
			u.User = &UserInfo{}
			vs = strings.SplitN(vs[0], ":", 2)
			u.User.Username = vs[0]
			if len(vs) == 2{
				u.User.Password = vs[1]
			}
		}else{
			//no user info
			raw = vs[0]
		}
	}
	//host
	var defaultPath = ""
	{
		vs:=strings.SplitN(raw, "/", 2)
		if len(vs) == 2{
			defaultPath = "/"
			u.Host = HostFromString(vs[0])
			flags |= FlagHost
			raw = vs[1]
		}
		//else{
		//	//no host here, we make it before ?#
		//	return &u
		//}
	}
	//path
	{
		vs:=strings.SplitN(raw, "?", 2)
		if len(vs) == 2{
			if flags & FlagHost == 0{
				u.Host = HostFromString(vs[0])
				flags |= FlagHost

				flags |= FlagPath
				u.Path = defaultPath
			}else{
				flags |= FlagPath
				u.Path = defaultPath + vs[0]
			}
			raw = vs[1]
		}
	}
	//query and fragment
	{
		vs:=strings.SplitN(raw, "#", 2)
		if flags & FlagHost == 0{
			u.Host = HostFromString(vs[0])
			flags |= FlagHost
			flags |= FlagPath
			u.Path = defaultPath
			flags |= FlagRawQuery
			u.RawQuery = ""
		}else{
			if flags & FlagPath > 0{
				flags |= FlagRawQuery
				u.RawQuery = vs[0]
			}else{
				flags |= FlagPath
				u.Path = defaultPath + vs[0]
			}
		}
		if len(vs) == 2{
			flags |= FlagFragment
			u.Fragment = vs[1]
		}
	}
	return &u
}


func (u *URL) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	ux := UrlParse(s)
	if ux == nil{
		return fmt.Errorf("not url %s", s)
	}
	*u = *ux
	return nil
}



func (u *URL) WithPath(path string) string {
	var buf strings.Builder
	if u.Scheme != "" {
		buf.WriteString(u.Scheme)
		buf.WriteString("://")
	}
	if u.User != nil{
		var us = u.User.String()
		if us != ""{
			buf.WriteString(us)
			buf.WriteByte('@')
		}
	}

	if u.Host.Addr != ""{
		buf.WriteString(u.Host.Addr)
	}
	if u.Host.Port > 0{
		buf.WriteByte(':')
		buf.WriteString(strconv.FormatInt(int64(u.Host.Port), 10))
	}

	if path != ""{
		if path[0] != '/'{
			buf.WriteByte('/')
		}
		buf.WriteString(path)
	}else{
		buf.WriteByte('/')
	}

	//if u.RawQuery != ""{
	//	buf.WriteByte('?')
	//	buf.WriteString(u.RawQuery)
	//}
	//
	//if u.Fragment != ""{
	//	buf.WriteByte('#')
	//	buf.WriteString(u.Fragment)
	//}

	return buf.String()
}





type UrlInfoBuilder struct {
	Url *URL
	builder strings.Builder
	flags int64
}

func (b *UrlInfoBuilder)Scheme() *UrlInfoBuilder {
	if b.flags < FlagScheme {
		if b.Url.Scheme != "" {
			b.builder.WriteString(b.Url.Scheme)
			b.builder.WriteString("://")
		}
		b.flags = FlagScheme
	}
	return b
}

func (b *UrlInfoBuilder)User() *UrlInfoBuilder {
	if b.flags < FlagUser {
		if b.Url.User != nil{
			var us = b.Url.User.String()
			if us != ""{
				b.builder.WriteString(us)
				b.builder.WriteByte('@')
			}
		}
		b.flags = FlagUser
	}
	return b
}
func (b *UrlInfoBuilder)Addr() *UrlInfoBuilder {
	if b.flags < FlagHost {
		if b.Url.Host.Addr != ""{
			b.builder.WriteString(b.Url.Host.Addr)
		}
		b.flags = FlagHost
	}
	return b
}
func (b *UrlInfoBuilder)Port() *UrlInfoBuilder {
	if b.flags < FlagPort {
		if b.Url.Host.Port > 0{
			b.builder.WriteByte(':')
			b.builder.WriteString(strconv.FormatInt(int64(b.Url.Host.Port), 10))
		}
		b.flags = FlagPort
	}
	return b
}
func (b *UrlInfoBuilder)Path() *UrlInfoBuilder {
	if b.flags < FlagPath {
		if b.Url.Path != ""{
			if b.Url.Path[0] != '/'{
				b.builder.WriteByte('/')
			}
			b.builder.WriteString(b.Url.Path)
		}
		b.flags = FlagPath
	}
	return b
}
func (b *UrlInfoBuilder)RawQuery() *UrlInfoBuilder {
	if b.flags < FlagRawQuery {
		if b.Url.RawQuery != ""{
			b.builder.WriteByte('?')
			b.builder.WriteString(b.Url.RawQuery)
		}
		b.flags = FlagRawQuery
	}
	return b
}
func (b *UrlInfoBuilder)Fragment() *UrlInfoBuilder {
	if b.flags < FlagFragment {
		if b.Url.Fragment != ""{
			b.builder.WriteByte('#')
			b.builder.WriteString(b.Url.Fragment)
		}
		b.flags = FlagFragment
	}
	return b
}

func (b *UrlInfoBuilder)ToString()string{
	return b.builder.String()
}