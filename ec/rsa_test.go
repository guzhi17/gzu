package ec

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"testing"
)


var(
	PrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCdTSKKqoyRcZRhBIMmYUClDUqcsbQGC+WeL5rAbto2/IBxsD+y
9WydzQaJ+omVdHTBrBj9is6Kj7T4r0xlRtbxjg5a+oNs/PxJEVT717rMQ4vJLwhJ
T96046VNXEzDpY3wogDCnw9JLgtvx7x5jnqfFDW6Wgm24t4+NqTgGMUCRwIDAQAB
AoGALzqCnXm5fM3KTBrLudFHVIcaGNPuBka6KXWHlDF0SUAk3H2bkoLHmtV9Gh5k
AsCVcbTXSADOJKIjJuuTF3FehW2fi4FPnAz2B4h0jxNQReMUgHaE6Mkky8bcuaee
4dWOFNaDMRvfNATBjwp4ELadyEMqZ0T1rUT0Pltb4QnbiMECQQDO4wQQkMZaO4kU
zRC1Eiz1jNnOAYJHuN3T5QwGfzt4SeNeQCGCLCrhF+CDS0J2JLeTy7w17WX91w3S
6T+gHJihAkEAwqSyOIBdiZp0atsB1vYYS0w1ROF5N7fd23TyKgRPxbTYKaQJvUWU
kCQhCETjK8QnUSBqNhqzI4j1+rZTmRWp5wJAf4gDhm6oRxEyJGdwqB3nJwrHbK0T
cUDtRWSJMCwYLcNmbEAeJ88wM4dzd5vaAVgK7gmGILwRxhNeSyhLd1iJYQJALa2g
9YmKagSJVZpX8C6IvQMBbUzMubq4ogvr2NhyMB+kqwEIGBcAKmOQLPSdq2O5JlzJ
EDFr4Ob/cvre24ot+QJAINJDV43HJAQtOXHW2QbDkEnnEdTCAreKs5I93zURSqop
1b8bOymrYQ85YVcnQ6dFmOAcHks+z3kvq5uXEE+bZA==
-----END RSA PRIVATE KEY-----
`)

	PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCdTSKKqoyRcZRhBIMmYUClDUqc
sbQGC+WeL5rAbto2/IBxsD+y9WydzQaJ+omVdHTBrBj9is6Kj7T4r0xlRtbxjg5a
+oNs/PxJEVT717rMQ4vJLwhJT96046VNXEzDpY3wogDCnw9JLgtvx7x5jnqfFDW6
Wgm24t4+NqTgGMUCRwIDAQAB
-----END PUBLIC KEY-----
`)
)

func TestRsaEncrypt(t *testing.T) {
	log.SetFlags(11)


	origData := []byte("hello, this is a rsa encrypt test.")

	////////////////////////////////////////////
	block, _ := pem.Decode(PublicKey)
	if block == nil {
		log.Println(ErrKey)
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	////////////////////////////////////////////

	ct, err := RsaEncryptWithKey(origData, pub)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(ct, err)


	////////////////////////////////////////////

	block, _ = pem.Decode(PrivateKey)
	if block == nil {
		log.Println(ErrKey)
		return
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return
	}

	rt, err := RsaDecryptWithKey(ct, priv)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(rt))

}

func TestRsaSign(t *testing.T) {
	////////////////////////////////////////////
	block, _ := pem.Decode(PublicKey)
	if block == nil {
		log.Println(ErrKey)
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return
	}
	pub := pubInterface.(*rsa.PublicKey)



	block, _ = pem.Decode(PrivateKey)
	if block == nil {
		log.Println(ErrKey)
		return
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return
	}

	data := []byte("Hello this is my signature")
	digest := sha256.Sum256(data)
	signature, signErr := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest[:])
	if signErr != nil {
		t.Errorf("Could not sign message:%s", signErr.Error())
	}

	//pub := privateKey.PublicKey
	verifyErr := rsa.VerifyPKCS1v15(pub, crypto.SHA256, digest[:], signature)
	if verifyErr != nil {
		t.Errorf("Verification failed: %s", verifyErr)
	}
	log.Println("ok")
}
