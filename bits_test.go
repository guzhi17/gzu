// ------------------
// User: pei
// DateTime: 2020/5/16 8:42
// Description: 
// ------------------

package gzu

import "testing"

func TestHi(t *testing.T)  {
	t.Log(BitHighestInt64(0))
	t.Log(BitHighestInt64(1))
	t.Log(BitHighestInt64(8))
	t.Log(BitHighestInt64(14))
	t.Log(BitHighestInt64(15))
	t.Log(BitHighestInt64(16))
	t.Log(BitHighestInt64(17))
}

func TestNormalTo2N(t *testing.T) {
	t.Log(NormalTo2N(0))
	t.Log(NormalTo2N(1))
	t.Log(NormalTo2N(2))
	t.Log(NormalTo2N(4))
	t.Log(NormalTo2N(7))
	t.Log(NormalTo2N(8))
	t.Log(NormalTo2N(9))
}