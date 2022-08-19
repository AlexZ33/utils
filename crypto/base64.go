package crypto

import "encoding/base64"

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, string) {
	return base64.StdEncoding.DecodeString(data)
}

// EncryptAesBase64 先做AES加密，再做base64加密
func EncryptAesBase64(d string) (string, error) {
	aesEnc, e := auth
}
