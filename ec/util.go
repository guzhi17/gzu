package ec

import "encoding/hex"

func ToHex(v []byte) string {
	return hex.EncodeToString(v)
}

func FromHex(v string) ([]byte, error){
	return hex.DecodeString(v)
}

//func BytesIsEqual(a, b []byte) bool {
//	return false
//}