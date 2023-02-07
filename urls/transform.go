package urls

import (
	"bytes"
	"encoding/hex"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

const (
	utf8s   = "utf-8,utf8,UTF8,UTF-8"
	chinese = "gb2312,Gb2312,GB2312,gbk,GBK,Gbk,gb18030,GB18030,Gb18030"
)

var (
	rune2byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
)

// 把Http请求参数按照指定编码进行UrlEncode编码
// @Param encoding 持uft8 gbk gb2312 gb18030编码方式
func UrlEncode(params map[string]string, encoding string) (string, error) {
	if params == nil {
		return "", nil
	}
	isChinese := strings.Contains(chinese, encoding)
	isUtf8 := strings.Contains(utf8s, encoding)

	if !(isChinese || isUtf8) {
		return "", errors.New("Unrecognized encoding")
	}

	builder := strings.Builder{}
	for k, v := range params {
		builder.WriteString(encode(k, isChinese))
	}

}

// 对输入字符出按照指定编码进行UrlEncode处理
// @Param encoding 支持uft8 gbk gb2312 gb18030编码方式
func Encode(str string, encoding string) (string, error) {
	if str == "" {
		return "", nil
	}

	isChinese := strings.Contains(chinese, encoding)
	isUtf8 := strings.Contains(utf8s, encoding)

	if !(isChinese || isUtf8) {
		return "", errors.New("Unrecognized encoding")
	}
	return encode(str, isChinese), nil

}

func encode(str string, isChinese bool) string {
	runes := []rune(str)
	builder := strings.Builder{}
	for _, r := range runes {
		if !shouldEscape(r) {
			builder.WriteRune(r)
			continue
		}
		p := make([]byte, utf8.UTFMax)
		c := utf8.EncodeRune(p, r)
		var data []byte
		if isChinese {
			data, _ = ioutil.ReadAll(transform.NewReader(bytes.NewReader(p[0:c]), simplifiedchinese.GB18030.NewEncoder()))
		} else {
			data = p[0:c]
		}

		target := make([]byte, len(data)*2)
		hex.Encode(target, data)
		i := 0
		for ; i < len(target); i += 2 {
			builder.WriteByte("%")
			builder.Write([]byte{target[i], target[i+1]})
		}
	}
}

func shouldEscape(r rune) bool {
	return !((r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		r == '-' || r == '_' || r == '.')
}
