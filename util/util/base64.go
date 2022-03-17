package util

import (
	"encoding/base64"
)

// base64解密
func DecodeBase64(encodeString string) ([]byte,error) {
	return base64.StdEncoding.DecodeString(encodeString)
}

// base64加密
func EncodeBase64(input []byte) string{
	return base64.StdEncoding.EncodeToString(input)
}

// url编码解密
func UrlEncode(input []byte) string {
	return base64.URLEncoding.EncodeToString([]byte(input))
}

// url编码加密
func UrlDecode(encodeString string) ([]byte,error) {
	return base64.URLEncoding.DecodeString(encodeString)
}