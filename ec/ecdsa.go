package ec

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io"
	"math/big"
	"strings"
)

//var EcdsaPb = `-----BEGIN EC PUBLIC KEY-----
//MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8GMZ3Ua6uKTUD+CLOZBrkdwPvS9k
//xZdo9VBXQRY6n/11vHtlK1TeJYrnAT4zlQXpKmV9EoZXxHIolHuwKJt7YA==
//-----END EC PUBLIC KEY-----
//`
//var EcdsaPr = `-----BEGIN EC PRIVATE KEY-----
//MHcCAQEEIIhEtsEVq3vURwHn6oFu7fh4kZWx8ZrFeRuTU448qyhdoAoGCCqGSM49
//AwEHoUQDQgAE8GMZ3Ua6uKTUD+CLOZBrkdwPvS9kxZdo9VBXQRY6n/11vHtlK1Te
//JYrnAT4zlQXpKmV9EoZXxHIolHuwKJt7YA==
//-----END EC PRIVATE KEY-----
//`
//
//const (
//	EcdsaSalt   = "2ol8l1o9b@tsalt"
//	EcdsaRander = "2ol8l1o9b@tEcdsaRander"
//)
//var(
//	srander = strings.NewReader(EcdsaRander)
//)

//随机生成一对ECDSA的公私钥
func DeriveEcdsaKeyToString() (pubKeyStr []byte, prKeyStr []byte, err error) {
	rander0 := Rand256Int()
	rander1 := Rand64Int()
	rander := make([]byte, 40)
	copy(rander[:32], rander0[:])
	copy(rander[32:], rander1[:])
	prikey, err := ecdsa.GenerateKey(elliptic.P256(), strings.NewReader(string(rander)))
	if err != nil{
		return nil, nil, err
	}

	pubkey := prikey.Public()

	privbytes, err := x509.MarshalECPrivateKey(prikey)
	if err != nil{
		return nil, nil, err
	}
	block := pem.Block{}
	block.Bytes = privbytes
	block.Type = "EC PRIVATE KEY"
	pubKeyStr = pem.EncodeToMemory(&block)


	pubbytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil{
		return nil, nil, err
	}

	block = pem.Block{}
	block.Type = "EC PUBLIC KEY"

	block.Bytes = pubbytes
	prKeyStr = pem.EncodeToMemory(&block)

	return pubKeyStr, prKeyStr, nil
}

//通过pbkey的string反序列化出ecdsa的pubkey
func EcdsaPubKeyStruct(pbKeyStr string) *ecdsa.PublicKey {
	pbBytes := []byte(pbKeyStr)
	pbBlock, _ := pem.Decode(pbBytes)

	pbKey, _ := x509.ParsePKIXPublicKey(pbBlock.Bytes)
	pb, _ := pbKey.(*ecdsa.PublicKey)
	return pb
}

//通过prkey的string反序列化出ecdsa的prikey
func EcdsaPriKeyStruct(prKeyStr string) *ecdsa.PrivateKey {
	prBytes := []byte(prKeyStr)
	prBlock, _ := pem.Decode(prBytes)
	prKey, _ := x509.ParseECPrivateKey(prBlock.Bytes)
	return prKey
}

//签名
func SignEcdsa(data, salt []byte, priv *ecdsa.PrivateKey, rand io.Reader) (sig string, err error) {
	hashed := hashtext(data, salt)
	r, s, err := ecdsa.Sign(rand, priv, hashed)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

/**
证书分解
通过hex解码，分割成数字证书r，s
*/
func GetEcdsaSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("GetEcdsaSign DecodeString error, " + err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("GetEcdsaSign NewReader error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		err = errors.New("GetEcdsaSign Read read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("GetEcdsaSign rs split fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("GetEcdsaSign r UnmarshalText  fail")
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("GetEcdsaSign s UnmarshalText  fail")
		return
	}
	return
}

func VerifyEcdsa(text, salt []byte, signature string, key *ecdsa.PublicKey) (bool, error) {
	rint, sint, err := GetEcdsaSign(signature)
	if err != nil {
		return false, err
	}
	hashed := hashtext(text, salt)
	result := ecdsa.Verify(key, hashed, &rint, &sint)
	return result, nil
}

func hashtext(data, salt []byte) []byte {
	Sha1Inst := sha1.New()
	Sha1Inst.Write(data)
	result := Sha1Inst.Sum(salt)
	return result
}
