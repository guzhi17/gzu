// ------------------
// User: pei
// DateTime: 2019/12/27 11:20
// Description: 
// ------------------

package ec

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)



func TestNewRsaSigner(t *testing.T) {


	//pem := []byte("-----BEGIN RSA PRIVATE KEY-----MIICWwIBAAKBgQCdTSKKqoyRcZRhBIMmYUClDUqcsbQGC+WeL5rAbto2/IBxsD+y9WydzQaJ+omVdHTBrBj9is6Kj7T4r0xlRtbxjg5a+oNs/PxJEVT717rMQ4vJLwhJT96046VNXEzDpY3wogDCnw9JLgtvx7x5jnqfFDW6Wgm24t4+NqTgGMUCRwIDAQABAoGALzqCnXm5fM3KTBrLudFHVIcaGNPuBka6KXWHlDF0SUAk3H2bkoLHmtV9Gh5kAsCVcbTXSADOJKIjJuuTF3FehW2fi4FPnAz2B4h0jxNQReMUgHaE6Mkky8bcuaee4dWOFNaDMRvfNATBjwp4ELadyEMqZ0T1rUT0Pltb4QnbiMECQQDO4wQQkMZaO4kUzRC1Eiz1jNnOAYJHuN3T5QwGfzt4SeNeQCGCLCrhF+CDS0J2JLeTy7w17WX91w3S6T+gHJihAkEAwqSyOIBdiZp0atsB1vYYS0w1ROF5N7fd23TyKgRPxbTYKaQJvUWUkCQhCETjK8QnUSBqNhqzI4j1+rZTmRWp5wJAf4gDhm6oRxEyJGdwqB3nJwrHbK0TcUDtRWSJMCwYLcNmbEAeJ88wM4dzd5vaAVgK7gmGILwRxhNeSyhLd1iJYQJALa2g9YmKagSJVZpX8C6IvQMBbUzMubq4ogvr2NhyMB+kqwEIGBcAKmOQLPSdq2O5JlzJEDFr4Ob/cvre24ot+QJAINJDV43HJAQtOXHW2QbDkEnnEdTCAreKs5I93zURSqop1b8bOymrYQ85YVcnQ6dFmOAcHks+z3kvq5uXEE+bZA==-----END RSA PRIVATE KEY-----")
	//
	//log.Println(NewRsaSignerFromPem(pem))

	signer, err := NewRsaSignerFromPemFile("sign_pri.pem")
	if err != nil{
		t.Fatal(err)
	}
	unsigner, err := NewRsaSignerFromPemFile("sign_pub.pem")
	if err != nil{
		t.Fatal(err)
	}
	data := []byte("hello, this is a rsa sign test")
	sign, err := signer.Sign(data)
	if err != nil{
		t.Fatal(err)
	}

	err = unsigner.UnSign(data, sign)
	if err != nil{
		t.Fatal(err)
	}

	t.Log(base64.StdEncoding.EncodeToString(sign))
	t.Log(hex.EncodeToString(sign))
	t.Log(BytesIntToString(sign, 62))
	t.Log("ok")

}