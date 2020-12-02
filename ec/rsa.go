package ec

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"math/big"
	"os"
)

//openssl genrsa -out private.pem 1024
//openssl rsa -in private.pem -pubout -out public.pem

var (


	ErrKey = errors.New("key")
)


func RsaEncryptWithKey(origData []byte, publicKey *rsa.PublicKey) (rt []byte, err error) {
	l := len(origData)
	if l < 1{
		return origData, nil
	}
	//分片
	k := (publicKey.N.BitLen() + 7) / 8 - 11
	sp, ep := 0,k
	if ep>l{
		ep = l
	}
	var buffer = bytes.Buffer{}
	for {
		b, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, origData[sp:ep])
		if err != nil{
			return nil, err
		}
		buffer.Write(b)
		sp, ep = sp + k, ep +k
		if ep>l{
			ep = l
		}
		if sp >= l{
			break
		}
	}
	return buffer.Bytes(), err
	//return rsa.EncryptPKCS1v15(rand.Reader, publicKey, origData)
}

func RsaDecryptWithKey(ciphertext []byte, privateKey *rsa.PrivateKey) (rt []byte, err error) {

	l := len(ciphertext)
	if l < 1{
		return ciphertext, nil
	}

	pub := privateKey.Public().(*rsa.PublicKey)

	k := (pub.N.BitLen() + 7) / 8
	sp, ep := 0,k
	if ep>l{
		ep = l
	}
	var buffer = bytes.Buffer{}
	for {
		b, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext[sp:ep])
		if err != nil{
			return nil, err
		}
		buffer.Write(b)

		sp, ep = sp + k, ep +k
		if ep>l{
			ep = l
		}
		if sp >= l{
			break
		}
	}
	return buffer.Bytes(), err
	//return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

func RsaPublicKeyToXmlString(pub rsa.PublicKey)string{

	return fmt.Sprintf(`<RSAKeyValue><Modulus>%s</Modulus><Exponent>%s</Exponent></RSAKeyValue>`,
		base64.StdEncoding.EncodeToString(pub.N.Bytes()),
		base64.StdEncoding.EncodeToString(big.NewInt(int64(pub.E)).Bytes()))
}

func RsaPublicKeyFromPemBytes(publicKey []byte)*rsa.PublicKey{
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	return pubInterface.(*rsa.PublicKey)
}

func RsaPublicKeyFromString(publicKey string)*rsa.PublicKey{
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	return pubInterface.(*rsa.PublicKey)
}

func RsaPublicKeyFromDer(der []byte)*rsa.PublicKey{
	pubInterface, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil
	}
	return pubInterface.(*rsa.PublicKey)
}

//n Modules,  e Exponent
func RsaPublicKeyFromModulesAndExponent(n, e []byte)*rsa.PublicKey{
	E := big.Int{}
	E.SetBytes(e)
	ie := int(E.Int64())
	if ie<0{
		return nil
	}

	N := big.Int{}
	N.SetBytes(n)
	if N.Sign() < 0{
		return nil
	}

	return &rsa.PublicKey{
		N:&N,
		E:ie,
	}
}


type XmlRSAKeyValue struct {
	Modulus string
	Exponent string
}

func RsaPublicKeyFromXml(nexml string) *rsa.PublicKey  {
	//<RSAKeyValue><Modulus>M</Modulus><Exponent>C</Exponent></RSAKeyValue>
	x := XmlRSAKeyValue{}
	err := xml.Unmarshal([]byte(nexml), &x)
	if err != nil{
		return nil
	}
	e, err := base64.StdEncoding.DecodeString(x.Exponent)
	if err != nil{
		return nil
	}
	n, err := base64.StdEncoding.DecodeString(x.Modulus)
	if err != nil{
		return nil
	}
	return RsaPublicKeyFromModulesAndExponent(n, e)
}


func RsaPublicKeyToBuffer(publicKey *rsa.PublicKey, bw *bytes.Buffer)(error){
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return ErrKey
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(bw, block)
	if err != nil {
		return ErrKey
	}
	return nil
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}



type ByteWriter struct {
	_d []byte
}

func (w *ByteWriter)Write(p []byte) (n int, err error){
	w._d = append(w._d, p...)
	return len(p), nil
}

func (w *ByteWriter)String()string{
	if w._d == nil{
		return ""
	}
	return string(w._d)
}