package time

import (
	"github.com/AlexZ33/utils/errors"
	"time"
)

type Time struct {
	time.Time
}

// GetCurrentTime 返回格式化的当前时间
func GetCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.RFC3339)
}

func GetTimeFromTimeStamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func Time2String(t time.Time) string {
	// https://stackoverflow.com/questions/55409774/the-result-of-time-formatting-of-rfc3339-in-go-on-linux-and-macos-are-different
	lo, _ := time.LoadLocation("Local")
	return t.In(lo).Format(time.RFC3339)
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

//
func ChangeStrTimeToUTC(strTime string) time.Time {
	if strTime == "" {
		return time.Unix(0, 0)
	}
	// https://gosamples.dev/date-time-format-cheatsheet/
	// 软件包time: const TimeFormat string = "2006-01-02 15:04:05"
	timeResult, err := time.ParseInLocation(TimeFormat, strTime, time.Local)
	if err != nil {
		errors.New("ChangeStrTimeToUTC error: %v", err)
		return time.Now()
	}
	return timeResult

}
