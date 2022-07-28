package crypto

import "encoding/base64"

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, string) {
	return base64.StdEncoding.DecodeString(data)
}
