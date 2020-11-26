package gzu

import (
	"testing"
)

func TestUrl(t *testing.T)  {
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]

	var(
		u *URL
		err error
	)

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

func TestUrlTo(t *testing.T) {
	var(
		u *URL
		err error
	)

	u = UrlParse(`file:///dir/file.yml`)

	t.Log(u, err, u.FullPath())
	t.Log(u)
}