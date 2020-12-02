// ------------------
// User: pei
// DateTime: 2020/2/20 13:30
// Description: 
// ------------------

package ec

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var(
	ErrDataFormat = errors.New("error data format")
	ErrEncrypt = errors.New("error encrypt")

)


type EncryptType uint32

const (
	EncryptTypeNone EncryptType = 0
	EncryptTypeAes = 1
	EncryptTypeDes = 2
)
type FileType uint32
const(
	FileTypeNone FileType = 0
	FileTypeJson = 1
	FileTypeYml = 2
	FileTypeXml = 3
)


func Ext(fn string)string {
	ext := path.Ext(fn)
	if strings.HasPrefix(ext, "."){
		ext = ext[1:]
	}
	return strings.ToLower(ext)
}

func GetFileType(fn string) FileType  {
	switch Ext(fn){
	default: return FileTypeNone
	case "yml", "yaml": return FileTypeYml
	case "json": return FileTypeJson
	case "xml": return FileTypeXml
	}
}

func (t FileType)Ext()string{
	switch t {
	default: return ""
	case FileTypeYml: return ".yml"
	case FileTypeJson: return ".json"
	case FileTypeXml: return ".xml"
	}
}

const (
	HeaderSize = 17
	TagName = "BATEN"
)
//TYPE: 1 json, 2 yaml 3 xml
//<<"BATEN",VER:32,ENC:32,TYPE:32, DATA>>
type PackageHeader struct {
	Tag string //"BATEN"
	Ver uint32
	Enc EncryptType
}
type PackageEnc struct {
	//header
	PackageHeader
	Type FileType
	//header end
	Body []byte
}

type Package struct {
	//header
	Type FileType
	//header end
	Body []byte
}

func PackageEncParse(b []byte, tag string) (r *PackageEnc, err error){
	if len(b) < HeaderSize{
		return nil, ErrDataFormat
	}
	if string(b[:len(tag)]) != tag{
		return nil, ErrDataFormat
	}
	r = &PackageEnc{
		PackageHeader: PackageHeader{
			Tag:  tag,
		},
	}
	r.Ver = binary.BigEndian.Uint32(b[5:])
	r.Enc = EncryptType(binary.BigEndian.Uint32(b[9:]))
	r.Type = FileType(binary.BigEndian.Uint32(b[13:]))
	r.Body = b[HeaderSize:]
	return
}

func PackageEncFromFile(fn string, tag string) (r *PackageEnc, err error){
	raw, err := ioutil.ReadFile(fn)
	if err != nil{
		return nil, err
	}
	return PackageEncParse(raw, tag)
}
func (p *PackageEnc)ToFile(fn string) (err error) {
	raw, err := p.ToBytes()
	if err != nil{
		return
	}
	return ioutil.WriteFile(fn, raw, os.ModePerm)
}
func (p *PackageEnc)ToBytes() (raw []byte, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(p.Tag)
	err = binary.Write(&buffer, binary.BigEndian, p.Ver)
	if err != nil{
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, p.Enc)
	if err != nil{
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, p.Type)
	if err != nil{
		return nil, err
	}
	buffer.Write(p.Body)
	return buffer.Bytes(), nil
}

func (p *PackageEnc)Decrypt(pwds string) (*Package, error) {
	var (
		r = Package{
			Type: p.Type,
		}
		err error
		pwd = Md5String(pwds)
	)
	switch EncryptType(p.Enc) {
	default:
		return nil, ErrEncrypt
	case EncryptTypeNone:
		r.Body = p.Body
	case EncryptTypeAes:
		r.Body, err = AESCBCDecrypt(nil, p.Body, pwd, pwd)
	case EncryptTypeDes:
		r.Body, err = DESCBCDecrypt(nil, p.Body, pwd, pwd)
	}
	if err != nil{
		return nil, err
	}
	return &r, nil
}

func (p *Package)Encrypt(pwds string, header PackageHeader) (*PackageEnc, error) {
	var (
		r = PackageEnc{
			PackageHeader: header,
			Type: p.Type,
		}
		err error
		pwd = Md5String(pwds)
	)
	switch EncryptType(header.Enc) {
	default:
		return nil, ErrEncrypt
	case EncryptTypeNone:
		r.Body = p.Body
	case EncryptTypeAes:
		r.Body, err = AESCBCEncrypt(nil, p.Body, pwd, pwd)
	case EncryptTypeDes:
		r.Body, err = DESCBCEncrypt(nil, p.Body, pwd, pwd)
	}
	if err != nil{
		return nil, err
	}
	return &r, nil
}
