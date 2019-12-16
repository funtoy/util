package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func Md5Byte(in []byte) []byte {
	h := md5.New()
	h.Write(in)
	return h.Sum(nil)
}

func Md5String(in string) string {
	return fmt.Sprintf("%x", Md5Byte([]byte(in)))
}

func Md5Str(in ...interface{}) string {
	return fmt.Sprintf("%x", Md5Byte([]byte(fmt.Sprint(in...))))
}

func Sha1Byte(in []byte) []byte {
	h := sha1.New()
	h.Write(in)
	return h.Sum(nil)
}

func Sha1String(in string) string {
	return fmt.Sprintf("%x", Sha1Byte([]byte(in)))
}

func Sha256(in string) []byte {
	hash := sha256.New()
	hash.Write([]byte(in))
	return hash.Sum(nil)
}

func Base64StrEncode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func Base64StrDecode(in string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(in)
}

func Base64UrlSaveEncode(in string) string {
	in = base64.StdEncoding.EncodeToString([]byte(in))
	in = strings.Split(in, "=")[0]
	in = strings.Replace(in, "+", "-", -1)
	in = strings.Replace(in, "/", "_", -1)
	return in
}

func Base64UrlSaveDecode(in string) ([]byte, error) {
	s := in
	s = strings.Replace(s, "-", "+", -1)
	s = strings.Replace(s, "_", "/", -1)
	switch len(s) % 4 {
	case 0:
	case 2:
		s += "=="
	case 3:
		s += "="
	default:
		return nil, fmt.Errorf("Base64UrlSaveDecode fail")
	}
	return base64.StdEncoding.DecodeString(s)
}
