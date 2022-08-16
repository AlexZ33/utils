package json

import (
	"encoding/json"
	"fmt"
	"github.com/AlexZ33/utils/errors"
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
