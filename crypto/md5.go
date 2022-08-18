package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMd5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(str string) string {
	return ToMd5([]byte(str))
}

/**
* PasswordMD5 密码加密
* @Description: create Account password MD5
* @param passwd 密码
* @return: string
 */
func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	// 加盐： 后面增加一个无意义的字符串，防止破解
	h.Write([]byte(passwd + salt + "x@.YuO-]"))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}
