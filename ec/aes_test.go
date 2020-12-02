package ec

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestAESCBCEncrypt(t *testing.T) {
	log.SetFlags(11)

	pwd := []byte("hello this is my password, mustbe 32 length or 16 so you can make me a favor")


	ct, err := AESCBCEncrypt(nil, []byte("google is the best search engine all over the world"), pwd[:16], pwd)
	if err != nil{
		log.Println(err)
		return
	}

	rt, err := AESCBCDecrypt(nil, ct, pwd[:16], pwd)
	if err != nil{
		log.Println(err)
		return
	}
	log.Println(string(rt))



	crypter, err := NewCBCCrypterDes7(pwd[:8], pwd)
	if err != nil{
		log.Println(err)
		return
	}
	ct, err = crypter.Encrypt(nil, []byte("google is the best search engine all over the world"))
	if err != nil{
		log.Println(err)
		return
	}
	rt, err = crypter.Decrypt(nil, ct)
	if err != nil{
		log.Println(err)
		return
	}
	log.Println(string(rt))

}

func TestAESCBCDecrypt(t *testing.T) {
	k := []byte("0123456789abcdef")
	a, err := NewCBCCrypterAes5(k, k)
	if err != nil{
		t.Fatal(err)
	}
	v, err := a.Encrypt(nil, []byte("hello, this is a test for"))
	if err != nil{
		t.Fatal(err)
	}
	t.Log(v)
	bv := base64.StdEncoding.EncodeToString(v)
	t.Log(bv)


	v, err = a.Decrypt(nil, v)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(v)
	t.Log(string(v))
}