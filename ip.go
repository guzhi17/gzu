package gzu

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const(
 	MAX_UINT = 0xffffffff
)


type ErrorIP struct {
	Val string
}
var _ error = ErrorIP{}
func (e ErrorIP)Error() string {
	var b strings.Builder
	b.WriteString("not ip format: ")
	b.WriteString(e.Val)
	return b.String()
}
//10.10.1.1/24
func GetLocalAddrOfSubNet(subnet string) (r []string, err error){
	sn := SubNetFromString(subnet)
	if sn == nil{
		return nil, ErrorIP{Val: subnet}
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ipm := strings.SplitN(addr.String(), "/", 2)
		ok, err := sn.IsSub(ipm[0])
		if err != nil || !ok{
			continue
		}
		r = append(r, ipm[0])
	}
	return r, nil
}

func GetPublicAddr() (r []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ipm := strings.SplitN(addr.String(), "/", 2)
		ip, err := IpV4ToUint32(ipm[0])
		if err != nil{
			continue
		}
		if !IsPublicIpV4(ip){
			continue
		}
		r = append(r, ipm[0])
	}
	return r, nil
}


func GetAddrs() (r []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ipm := strings.SplitN(addr.String(), "/", 2)
		ip, err := IpV4ToUint32(ipm[0])
		if err != nil{
			continue
		}
		if IsLocalIpV4(ip){
			continue
		}
		r = append(r, ipm[0])
	}
	return r, nil
}


type SubNet struct {
	IP uint32
	Msk uint32
}

//10.10.1.1/24 or range ip1-ip2
func SubNetFromString(subnet string) *SubNet {
	ipsub := strings.Split(subnet, "/")
	vi, err := IpV4ToUint32(ipsub[0])
	if err != nil{
		return nil
	}
	vs := 32
	if len(ipsub) > 1{
		vs, err = strconv.Atoi(ipsub[1])
		if err != nil{
			return nil
		}
		if vs < 0 || vs > 32{
			return nil
		}
	}
	return &SubNet{
		IP: vi,
		Msk: uint32(vs),
	}
}
func (sn SubNet) Range() (uint32, uint32) {
	if sn.IP == 0 && sn.Msk == 0{
		//"0.0.0.0/0"
		return 0, MAX_UINT
	}
	l := 32 - sn.Msk
	min := MAX_UINT & sn.IP & (MAX_UINT << l)
	return min, min + (1 << l) - 1
}
func (sn SubNet) IsSub(ip string) (bool, error) {
	vip, err := IpV4ToUint32(ip)
	if err != nil{
		return false, err
	}
	min, max := sn.Range()
	return vip >= min && vip <= max, nil
}

func (sn SubNet)String()string {
	return fmt.Sprintf("%d.%d.%d.%d/%d", (sn.IP >> 24)&0xff, (sn.IP >> 16)&0xff, (sn.IP >> 8)&0xff, sn.IP & 0xff, sn.Msk)
}



func IpV4ToUint32(ips string) (r uint32, err error) {
	ipfs := strings.Split(ips, ".")
	if len(ipfs) != 4 {
		return 0, ErrorIP{Val: ips}
	}
	var ipis []uint32
	for _, iipf := range ipfs{
		iv,err := strconv.Atoi(iipf)
		if err != nil{
			return 0, err
		}
		ipis = append(ipis, uint32(iv))
	}
	r = (ipis[0]&0xff) << 24 | (ipis[1]&0xff) << 16 | (ipis[2]&0xff) << 8 | (ipis[3]&0xff)
	return r, nil
}

//
func IsLocalIpV4(ip uint32) bool {
	f0 := (ip >> 24)&0xff
	if f0 == 127{
		return true
	}
	return false
}


func IsPublicIpV4(ip uint32) bool {
	f0 := (ip >> 24)&0xff
	if f0 == 10 || f0 == 127{
		return false
	}
	f1 := (ip >> 16)&0xff
	if f0 == 172 && f1 > 15 && f1 < 32{
		return false
	}

	if f0 == 192 && f1 == 168{
		return false
	}
	return true
}

func IpV4FromUint32(ip uint32)string {
	return fmt.Sprintf("%d.%d.%d.%d", (ip >> 24)&0xff, (ip >> 16)&0xff, (ip >> 8)&0xff, ip & 0xff)
}
