package ec

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)


//128bit、192bit、256bit

var(
	ErrBlockSize = errors.New("input not multiple of block size")
	ErrIvLength = errors.New("cipher.AESCBCDecrypt: IV length must equal block size")
	ErrPaddingSize = errors.New("padding size error")
)


type ICrypter interface {
	Encrypt(dst, src[]byte) ([]byte, error)
	Decrypt(dst, src[]byte) ([]byte, error)
}


type Padding func(ciphertext []byte, blockSize int) []byte
type UnPadding func(origData []byte, blockSize int) ([]byte, error)

type CBCCrypter struct {
	cipher.Block
	Padding
	UnPadding
	iv []byte
	bs int
}

func NewCBCCrypterAes5(key, iv []byte) (*CBCCrypter, error) {
	ciph, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := ciph.BlockSize()
	if len(iv) < bs {
		return nil, ErrIvLength
	}
	//make a copy
	iv16bytes := iv[:bs]
	c := &CBCCrypter{
		Block: ciph,
		Padding:PKCS5Padding,
		UnPadding:PKCS5UnPadding,
		iv: iv16bytes,
		bs: bs,
	}
	return c, nil
}

func NewCBCCrypterAes7(key, iv []byte) (*CBCCrypter, error) {
	ciph, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := ciph.BlockSize()
	if len(iv) < bs {
		return nil, ErrIvLength
	}
	//make a copy
	iv16bytes := iv[:bs]
	c := &CBCCrypter{
		Block: ciph,
		Padding:PKCS7Padding,
		UnPadding:PKCS7UnPadding,
		iv: iv16bytes,
		bs: bs,
	}
	return c, nil
}

func NewCBCCrypterDes7(key, iv []byte) (*CBCCrypter, error) {
	ciph, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := ciph.BlockSize()
	if len(iv) < bs {
		return nil, ErrIvLength
	}
	//make a copy
	iv16bytes := iv[:bs]

	c := &CBCCrypter{
		Block: ciph,
		Padding:PKCS7Padding,
		UnPadding:PKCS7UnPadding,
		iv: iv16bytes,
		bs: bs,
	}
	return c, nil
}

func (c *CBCCrypter) Encrypt(dst, src []byte) ([]byte, error) {
	src = c.Padding(src, c.bs)
	if dst == nil {
		dst = make([]byte, len(src))
	}
	blockModel := cipher.NewCBCEncrypter(c.Block, c.iv)
	blockModel.CryptBlocks(dst, src)

	return dst, nil
}



func (c *CBCCrypter)Decrypt(dst, src []byte) ([]byte, error) {
	var(
		err error
	)
	if len(src)%c.bs != 0 {
		return nil,ErrBlockSize
	}
	if dst == nil {
		dst = make([]byte, len(src))
	}
	mode := cipher.NewCBCDecrypter(c.Block, c.iv)
	mode.CryptBlocks(dst, src)

	dst, err = c.UnPadding(dst, c.bs)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

//
//
////AES CBC
func AESCBCDecrypt(dst, src, key, iv []byte) ([]byte, error) {
	c, err := NewCBCCrypterAes5(key, iv)
	if err != nil{
		return nil, err
	}
	return c.Decrypt(dst, src)
}
func DESCBCEncrypt(dst, src, key, iv []byte) ([]byte, error) {
	c, err := NewCBCCrypterDes7(key, iv)
	if err != nil{
		return nil, err
	}
	return c.Encrypt(dst, src)
}
func DESCBCDecrypt(dst, src, key, iv []byte) ([]byte, error) {
	c, err := NewCBCCrypterDes7(key, iv)
	if err != nil{
		return nil, err
	}
	return c.Decrypt(dst, src)
}
//
//	if len(src)%ciph.BlockSize() != 0 {
//		return nil,ErrBlockSize
//	}
//
//	bs := ciph.BlockSize()
//
//	if len(iv) < bs {
//		return nil, ErrIvLength
//	}
//
//	if dst == nil {
//		dst = make([]byte, len(src))
//	}
//
//	iv16bytes := iv[:bs]
//
//	mode := cipher.NewCBCDecrypter(ciph, iv16bytes)
//	mode.CryptBlocks(dst, src)
//
//	dst, err = PKCS5UnPadding(dst, bs)
//	if err != nil {
//		return nil, err
//	}
//	return dst, nil
//}
//
//
////AES CBC
func AESCBCEncrypt(dst, src, key, iv []byte) ([]byte, error) {
	c, err := NewCBCCrypterAes5(key, iv)
	if err != nil{
		return nil, err
	}
	return c.Encrypt(dst, src)
}

func AESCBCEncrypt7(dst, src, key, iv []byte) ([]byte, error) {
	c, err := NewCBCCrypterAes7(key, iv)
	if err != nil{
		return nil, err
	}
	return c.Encrypt(dst, src)
}
//
//	bs := ciph.BlockSize()
//	if len(iv) < bs {
//		return nil, ErrIvLength
//	}
//
//	iv16bytes := iv[:bs]
//
//	src = PKCS5Padding(src, bs)
//	if dst == nil {
//		dst = make([]byte, len(src))
//	}
//
//	blockModel := cipher.NewCBCEncrypter(ciph, iv16bytes)
//	blockModel.CryptBlocks(dst, src)
//
//	return dst, nil
//}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte, blockSize int) ([]byte, error) {
	length := len(origData)
	paddingLen := int(origData[length-1])
	if paddingLen > length {
		return nil, ErrPaddingSize
	}
	return origData[:length-paddingLen], nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte, blockSize int) ([]byte, error) {
	length := len(origData)
	paddingLen := int(origData[length-1])
	if paddingLen > length {
		return nil, ErrPaddingSize
	}
	return origData[:length-paddingLen], nil
}