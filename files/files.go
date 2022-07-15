package files

import (
	"os"
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

func WriteStringToFile(fileName string, content string) error {
	return WriteBytesToFile(fileName, []byte(content))
}

func WriteBytesToFile(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	return err
}

func WriteString(path string, content string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	}
	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.WriteString(content)
	return err
}

func AppendLine(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	content = strings.Join([]string{content, "\n"}, "")
	_, err = file.WriteString(content)
	return err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
