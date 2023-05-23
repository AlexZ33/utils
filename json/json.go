package json

import (
	"encoding/json"
	"fmt"
	"github.com/AlexZ33/utils/errors"
	"strings"
)

func Parse(str string, t interface{}) error {
	return json.Unmarshal([]byte(str), t)
}

func ToStr(t interface{}) (string, error) {
	if t == nil {
		return "", nil
	}
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ToJsonStr(t interface{}) string {
	str, err := ToStr(t)
	if err != nil {
		errors.New(err)
		return fmt.Sprintln("json.ToJsonStr error:", err.Error())
	}
	return str
}

type JsonMap map[string]string

// FastJsonMap 快速解析单层jsonmap, eg: {"a": "1", "b": "2"}
func FastJsonMap(data string) (JsonMap, error) {
	if data == "" || !strings.HasPrefix(data, "{") || !strings.HasSuffix(data, "}") {
		return nil, errors.New("data is not jsonmap.[" + data + "]")
	}
	data = strings.TrimLeft(data, "{")
	data = strings.TrimRight(data, "}")
	kvList := strings.Split(data, ",")

	rsp := make(map[string]string, 5)
	for _, kv := range kvList {
		kv = strings.TrimSpace(kv)
		kvArr := strings.Split(kv, ":")
		if len(kvArr) != 2 {
			return nil, errors.New("data is not jsonmap.[" + data + "]")
		}
		//去掉引号
		k := strings.Trim(kvArr[0], "\"")
		v := strings.Trim(kvArr[1], "\"")

		rsp[k] = v
	}
	return rsp, nil
}

type JsonResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func JsonErrorCode(code int, message string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}
