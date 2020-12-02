package ec

import (
	"crypto/md5"
	"crypto/sha1"
	"io"
	"math/big"
)


const (
	//BASE64字符表,不要有重复
	//"<>:;',./?~!@#$CDVWX%^&*ABYZabcghijklmnopqrstuvwxyz01EFGHIJKLMNOPQRSTU2345678(def)_+|{}[]9/"
	base64Table 	= "<>:;',./?~!@#$CDVWX%^&*ABYZbghijklmnopqt1EFGHQRSTU(def)_+|{}[]95"
	encodeStd 	= "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	encodeURL 	= "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)


type EncodeMessage []byte

func (e EncodeMessage)Hex() string {
	return ToHex(e)
}
func (e EncodeMessage)Integer(n int) string {
	return (&big.Int{}).SetBytes(e).Text(n)
}
/**
 * 对一个字符串进行MD5加密,不可解密
 */
func Md5String(ss... string) EncodeMessage {
	h := md5.New()
	for _, s := range ss{
		h.Write([]byte(s)) //
	}
	return h.Sum(nil)
}

func Md5Bytes(ss... []byte) EncodeMessage {
	h := md5.New()
	for _, s := range ss{
		h.Write(s) //
	}
	return h.Sum(nil)
}

func Md5File(reader io.Reader) (EncodeMessage, error)  {
	h := md5.New()
	_, err := io.Copy(h, reader)
	if err != nil{
		return nil, err
	}
	return h.Sum(nil), nil
}

/*获取 SHA1 字符串*/
func SHA1String(ss... string) EncodeMessage {
	t := sha1.New()
	for _, s := range ss{
		t.Write([]byte(s))
	}
	return t.Sum(nil)
}

func SHA1Bytes(ss... []byte) EncodeMessage {
	t := sha1.New()
	for _, s := range ss{
		t.Write(s)
	}
	return t.Sum(nil)
}

func BytesIntToString(s []byte, base int) string {
	b := big.NewInt(0)
	b.SetBytes(s)
	return b.Text(base)
}