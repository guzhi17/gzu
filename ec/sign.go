// ------------------
// User: pei
// DateTime: 2019/12/27 11:00
// Description: 
// ------------------

package ec

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)


var(
	ErrNoKey = errors.New("ssh: no key found")
	ErrSign = errors.New("error sign")
	ErrPublicKey = errors.New("error public key for unsign only")
)


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Signer interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
	UnSign(data[]byte, sig []byte) error
}

type UnSigner interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	UnSign(data[]byte, sig []byte) error
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewRsaSignerFromPemFile(file string) (Signer, error) {
	pemBytes, err := ioutil.ReadFile(file)
	if err != nil{
		return nil, err
	}
	return NewRsaSignerFromPem(pemBytes)
}
//func NewRsaUnSignerFromPemFile(file string) (UnSigner, error) {
//	pemBytes, err := ioutil.ReadFile(file)
//	if err != nil{
//		return nil, err
//	}
//	return NewRsaUnSignerFromPem(pemBytes)
//}

func NewRsaSignerFromPem(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, ErrNoKey
	}
	var rawkey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsaKey
	case "PUBLIC KEY", "RSA PUBLIC KEY":
		rsaKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsaKey
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return NewRsaSignerFromKey(rawkey)
}

//
//func NewRsaUnSignerFromPem(pemBytes []byte) (UnSigner, error) {
//	block, _ := pem.Decode(pemBytes)
//	if block == nil {
//		return nil, ErrNoKey
//	}
//
//	var rawkey interface{}
//	switch block.Type {
//	case "PUBLIC KEY":
//		rsaKey, err := x509.ParsePKIXPublicKey(block.Bytes)
//		if err != nil {
//			return nil, err
//		}
//		rawkey = rsaKey
//	default:
//		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
//	}
//	return NewRsaUnSignerFromKey(rawkey)
//}


func NewRsaSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &RsaSigner{t}
	case *rsa.PublicKey:
		sshKey = &RsaUnSigner{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

//func NewRsaUnSignerFromKey(k interface{}) (UnSigner, error) {
//	var sshKey UnSigner
//	switch t := k.(type) {
//	case *rsa.PublicKey:
//		sshKey = &RsaUnSigner{t}
//	default:
//		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
//	}
//	return sshKey, nil
//}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type RsaUnSigner struct {
	*rsa.PublicKey
}

type RsaSigner struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha256
func (r *RsaSigner) Sign(data []byte) ([]byte, error) {
	d:=sha256.Sum256(data)
	//d := md5.Sum(data)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d[:])
}

func (r *RsaSigner) UnSign(data []byte, sig []byte) error {
	d:=sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(&r.PublicKey, crypto.SHA256, d[:], sig)
}


func (r *RsaUnSigner) Sign(data []byte) ([]byte, error) {
	return nil, ErrPublicKey
}
// Unsign verifies the message using a rsa-sha256 signature
func (r *RsaUnSigner) UnSign(message []byte, sig []byte) error {
	d:=sha256.Sum256(message)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, d[:], sig)
}


/////////////////////////////////////////
type SignerMd5 struct {

	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Salt []byte
}

func NewSignerMd5(salt string) *SignerMd5 {
	return &SignerMd5{
		Salt: []byte(salt),
	}
}

func (s *SignerMd5)Sign(data []byte) ([]byte, error){
	return Md5Bytes(data, s.Salt), nil
}
func (s *SignerMd5)UnSign(data[]byte, sig []byte) error{
	if bytes.Compare(Md5Bytes(data, s.Salt), sig) == 0{
		return nil
	}
	return ErrSign
}