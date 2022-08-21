/***
* 工具类 - 文件读取、写入、解析、拼接、删除、拷贝、移动等
**/

package files

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// LoadFile 读取文件内容
func LoadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// LoadJsonToObject 读取json文件
func LoadJsonToObject(filename string, obj interface{}) error {
	buf, err := LoadFile(filename)

	if buf == nil {
		return err
	}
	if err != nil {
		return err
	}

	e := json.Unmarshal(buf, &obj)
	if e != nil {
		return e
	}
	return nil
}

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

// PathExists check if the directory or file exits
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

func ListFiles(path string) []string {
	res := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			res = append(res, file.Name())
		}
	}

	return res
}

// CopyDir copy directory from src to dst
func CopyDir(src string, dst string) error {
	var (
		err     error
		dir     []os.FileInfo
		srcInfo os.FileInfo
	)
	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}
	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}
	if dir, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range dir {
		srcPath := path.Join(src, fd.Name())
		dstPath := path.Join(dst, fd.Name())

		if fd.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyFile(src, dist string) error {
	var (
		err     error
		srcFile *os.File
		dstFile *os.File
		srcInfo os.FileInfo
	)

	if srcInfo, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()

	if dstFile, err = os.Create(dist); err != nil {
		return err
	}

	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	return os.Chmod(dist, srcInfo.Mode())
}
