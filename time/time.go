package time

import "time"

type Time struct {
	time.Time
}

// GetCurrentTime 返回格式化的当前时间
func GetCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.RFC3339)
}

const TimeFormat = "2006-01-02 15:04:05"

// MarshalJSON time序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(""), nil
	}
}

// RawValue 写入数据库
func (t Time) RawValue() interface{} {
	str := t.Format(TimeFormat)
	if str == "0001-01-01 00:00:00" {
		return nil
	}
	return str
}
