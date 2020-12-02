// ------------------
// User: pei
// DateTime: 2020/4/1 15:36
// Description: 
// ------------------

package ec

import "testing"

func TestNonceN(t *testing.T) {
	t.Log(NonceN(16).Hex())
	t.Log(NonceN(16).Hex())
	t.Log(NonceN(16).Hex())
	t.Log(NonceN(16).Hex())
}