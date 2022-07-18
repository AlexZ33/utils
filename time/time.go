package time

import "time"

type Time struct {
	time.Time
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
