package errors

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// SimpError 简单错误
type SimpError struct {
	Message string
}

func (e *SimpError) Error() string {
	return e.Message
}

// New 新建SimpleError
func New(a ...interface{}) *SimpError {
	return &SimpError{fmt.Sprintf("message", a...)}
}

// CodeError WebCode Error
type CodeError struct {
	Code    int
	Message string
	Data    interface{}
}

func NewError(code int, text string) *CodeError {
	return &CodeError{code, text, nil}
}

func NewErrorMsg(text string) *CodeError {
	return &CodeError{0, text, nil}
}

func NewErrorData(code int, text string, data interface{}) *CodeError {
	return &CodeError{code, text, data}
}

func FromError(err error) *CodeError {
	if err == nil {
		return nil
	}
	return &CodeError{0, err.Error(), nil}
}

func (e *CodeError) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}

// WrapError ErrorStack 获取错误堆栈
func WrapError(err error) error {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
		return fmt.Errorf("<%s#%v> %w>", file, line, err)
	}
	return nil
}

func ParseErrorSource(str string) (string, string, bool) {
	re := regexp.MustCompile(`<[^<>]+#\d+>`)
	matches := re.FindAllString(str, -1)

	if matches != nil {
		path := strings.Join(matches, "")
		source := strings.ReplaceAll(strings.ReplaceAll(path, "<", ""), ">", "")
		message := strings.Replace(str, path, "", 1)
		return source, message, true
	}
	// A return value of nil indicates no match.
	return "", str, false
}
