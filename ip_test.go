package gzu

import (
	"log"
	"testing"
)

func TestTestIpV4(t *testing.T) {
	sn := SubNetFromString("10.10.1.99")
	if sn == nil{
		log.Println("no ip/mask")
		return
	}
	mmin, mmax := sn.Range()
	log.Println("min:", IpV4FromUint32(mmin))
	log.Println("max:", IpV4FromUint32(mmax))

	i, err := sn.IsSub("192.168.0.1")
	log.Println(i, err)
	i, err = sn.IsSub("10.10.1.99")
	log.Println(i, err)



	log.Println(GetLocalAddrOfSubNet("10.10.1.1/24"))
	log.Println(GetLocalAddrOfSubNet("127.0.0.1/24"))
	log.Println(GetLocalAddrOfSubNet("192.168.1.1/24"))

	log.Println(GetPublicAddr())
}


