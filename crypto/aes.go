package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func HexAesEncrypt(key, input string) (string, error) {
	hkey, _ := hex.DecodeString(key)
	return aesEncrypt(hkey, []byte(input))
}

func aesEncrypt(key, input []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := md5.Sum(key)
	blockSize := block.BlockSize()
	origData := PKCS5Padding(input, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return hex.EncodeToString(crypted), nil
}

func UrlBase64Encode(src []byte) string {
	s := base64.URLEncoding.EncodeToString(src)
	return strings.Replace(s, "=", ".", 2)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if unpadding < length {
		return origData[:(length - unpadding)]
	}
	return origData
}
