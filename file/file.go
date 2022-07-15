package file

import (
	"path"
	"strings"
)

var stdImageExt = []string{"png", "jpg", "gif", "jpeg"}

// FileExt 获取文件后缀名
func FileExt(fileName string) string {
	// 取文件后缀名
	ext := path.Ext(fileName)
	// 去掉点 .
	return strings.ReplaceAll(ext, ".", "")
}

// FileType return image or file
// 根据后缀名区分文件是 file 还是 image
func FileType(fileName string) string {
	ext := FileExt(fileName)

	if ext == "" {
		return "file"
	}

	lowerCase := strings.ToLower(ext)

	for _, v := range stdImageExt {
		if v == lowerCase {
			return "image"
		}
	}
	return "file"
}
