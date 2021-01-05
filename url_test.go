package gzu

import (
	"log"
	"testing"
)

func TestUrl(t *testing.T)  {
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]

	var(
		u *URL
		err error
	)

	u = UrlParse(`10.10.1.2:9090#hello=nihao`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`1:`)
	t.Log(u, err, u.FullPath())
	u = UrlParse(`10.10.1.2:9090`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`./application.yml`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`/root/config`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`./`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`c:/dir/file.yml`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`file://c:/dir/file.yml`)
	t.Log(u, err, u.FullPath())

	u = UrlParse(`redis://root:bat2019!#$@118.31.126.111:27017/?authSource=admin`)
	t.Log(u, err, u.FullPath())
	//
	u = UrlParse(`https://me:pass@example.com/foo/bar?x=1&y=2#anchor`)
	t.Log(u, err, u.FullPath())
}



type XT struct {
	X string `json:"x"`
	A int `json:"a"`
}

func TestURL_GetQuery(t *testing.T) {
	u := UrlParse(`https://me:pass@example.com/foo/bar?x=q&a=8&=b&c=&x=1&y=2#anchor`)
	t.Log(u)
	log.Println(u.GetQuery("x"))
	log.Println(u.GetQuery("y"))
	log.Println(u.GetQuery("a"))
	log.Println(u.GetQuery("b"))
	log.Println(u.GetQuery("c"))
	log.Println(u.GetQuery("d"))

	var xt XT
	err := u.UnmarshalQuery(&xt)
	log.Println(err, xt)
}

func TestUrlTo(t *testing.T) {
	var(
		u *URL
		err error
	)

	u = UrlParse(`file:///dir/file.yml`)

	t.Log(u, err, u.FullPath())
	t.Log(u)
}

func TestUrlInfoBuilder_Addr(t *testing.T) {
	log.SetFlags(11)
	s := &UrlInfoBuilder{Url: UrlParse(`https://me:pass@example.com/foo/bar?x=q&a=8&=b&c=&x=1&y=2#anchor`)}
	log.Println(s.Scheme().User().Addr().Port().Path().RawQuery().Fragment().ToString())
}