// ------------------
// User: pei
// DateTime: 2019/12/4 14:45
// Description: 
// ------------------

package ec

import (
	"github.com/OneOfOne/xxhash"
	"log"
	"testing"
)

func TestGetHash(t *testing.T) {
	log.SetFlags(11)

	log.Println(xxhash.ChecksumString32("test"))
	log.Println(xxhash.ChecksumString32("12345"))
}