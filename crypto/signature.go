package crypto

import (
	"crypto/md5"
	"fmt"
)

func HexMd5Sign(p, t, k interface{}) string {
	ptk := fmt.Sprintf("p=%v&t=%v&key=%v", p, t, k)
	hash := md5.Sum([]byte(ptk))
	return fmt.Sprintf("%x", string(hash[:]))
}

func HexMd5String(str string) string {
	hash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", string(hash[:]))
}
