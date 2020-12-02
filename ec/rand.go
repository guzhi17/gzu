package ec

import (
	"bytes"
	"crypto/rand"
	"io"
	"strings"
	"time"

	mrand "math/rand"
)

var(
	gr = mrand.New(mrand.NewSource(time.Now().UnixNano()))
	gspool = []byte("0123456789abcdefchijklmnopqrstuvwxyz")
	gnpool = []byte("0123456789")
)

type Nonce []byte
func (n Nonce)Hex() string {
	return ToHex(n)
}

type Nonce8 [8]byte

func (n Nonce8)Hex() string {
	return ToHex(n[:])
}
type Nonce16 [16]byte
func (n Nonce16)Hex() string {
	return ToHex(n[:])
}
type Nonce32 [32]byte
func (n Nonce32)Hex() string {
	return ToHex(n[:])
}

//Rand64Int 随机获取64位整型
func Rand64Int() Nonce8 {
	var nonce [8]byte
	io.ReadFull(rand.Reader, nonce[:])
	return nonce
}

//Rand128Int 随机获取128位整型
func Rand128Int() Nonce16 {
	var nonce [16]byte
	io.ReadFull(rand.Reader, nonce[:])
	return nonce
}

//Rand256Int 随机获取256位整型
func Rand256Int() Nonce32 {
	var nonce [32]byte
	io.ReadFull(rand.Reader, nonce[:])
	return nonce
}

func NonceN(n int) Nonce {
	var nonce = make([]byte, n)
	io.ReadFull(rand.Reader, nonce[:])
	return nonce
}



func RandomBytes(pool []byte, length int) []byte {
	strLen := len(pool)
	buf := bytes.NewBuffer(nil)
	for i := 0; i < length; i++ {
		buf.WriteByte(pool[gr.Intn(strLen)])
	}
	return buf.Bytes()
}

func RandomString(pool []byte, length int) string {
	strLen := len(pool)
	buf := strings.Builder{}
	for i := 0; i < length; i++ {
		buf.WriteByte(pool[gr.Intn(strLen)])
	}
	return buf.String()
}

func RandomStr(length int) string {
	return RandomString(gspool, length)
}

func RandomNumStr(length int) string {
	return RandomString(gnpool, length)
}
