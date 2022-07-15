package errors

import (
	"fmt"
	"regexp"
	"runtime"
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
func New(message string, a ...interface{}) *SimpError {
	return &SimpError{fmt.Sprintf(message, a...)}
}

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
