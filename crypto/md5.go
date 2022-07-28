package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMd5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
