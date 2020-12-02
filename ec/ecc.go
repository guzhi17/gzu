package ec

import (
	"crypto/rand"
	"golang.org/x/crypto/curve25519"
	"io"
)

func GenEccKeys() (privateKey, ecdhePublic []byte, err error) {
	var scalar, public [32]byte
	if _, err = io.ReadFull(rand.Reader, scalar[:]); err != nil {
		return nil, nil, err
	}

	curve25519.ScalarBaseMult(&public, &scalar)
	privateKey = scalar[:]
	ecdhePublic = public[:]
	return privateKey, ecdhePublic, nil
}

func GenEccShareKey(pubKey []byte, priKey []byte) ([]byte, error) {
	var theirPublic, scalar [32]byte
	copy(theirPublic[:], pubKey)
	copy(scalar[:], priKey)
	return curve25519.X25519(scalar[:], theirPublic[:])
	//sharedKey
	//curve25519.ScalarMult(&sharedKey, &scalar, &theirPublic)
	//return sharedKey[:], nil
}