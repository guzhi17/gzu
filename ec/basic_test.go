// ------------------
// User: pei
// DateTime: 2019/10/16 10:45
// Description: 
// ------------------

package ec

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"
)

//52+10
const digitaltable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestHex(t *testing.T) {


	x := Md5String("hello")

	log.Println(x.Integer(62))
	log.Println(fmt.Sprintf("%s", x))

	var b big.Int
	b.SetBytes(x)
	xx := b.Text(62)
	log.Println(hex.EncodeToString(x), xx)


	log.Println(hex.EncodeToString(SHA1String("hello")))
}