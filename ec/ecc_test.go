package ec

import (
	"crypto/subtle"
	"log"
	"testing"
)

func TestGenEccKeys(t *testing.T) {
	log.SetFlags(11)

	priv1, pub1, err := GenEccKeys()
	if err != nil{
		log.Println(err)
		return
	}

	log.Println(priv1)
	log.Println(pub1)

	priv2, pub2, err := GenEccKeys()
	if err != nil{
		log.Println(err)
		return
	}

	log.Println(priv2)
	log.Println(pub2)

	S1, err := GenEccShareKey(pub1, priv2)
	if err != nil{
		log.Println(err)
		return
	}

	S2, err := GenEccShareKey(pub2, priv1)
	if err != nil{
		log.Println(err)
		return
	}

	log.Println(S1)
	log.Println(S2)

	log.Println(subtle.ConstantTimeCompare(S1, S2))
}