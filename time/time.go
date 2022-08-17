package time

import (
	"github.com/AlexZ33/utils/errors"
	"log"
	"math"
	"strings"
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

//ChangeStrTimeToUTC 把字符串时间转换为UTC时间
func ChangeStrTimeToUTC(strTime string) time.Time {
	if strTime == "" {
		return time.Unix(0, 0)
	}
	// https://gosamples.dev/date-time-format-cheatsheet/
	// 软件包time: const TimeFormat string = "2006-01-02 15:04:05"
	timeResult, err := time.ParseInLocation(TimeFormat, strTime, time.Local)
	if err != nil {
		errors.New(err)
		return time.Now()
	}
	return timeResult

}

// Timestamp 时间戳
func Timestamp(milliseconds float64) time.Time {
	seconds := math.Floor(milliseconds / 1000)
	nanoseconds := (milliseconds - seconds*1000) * 1000000
	return time.Unix(int64(seconds), int64(nanoseconds))
}

// ParseTimestamp 解析时间戳
func ParseTimestamp(value string) (time.Time, bool) {
	layout := ""
	length := len(value)
	switch {
	case strings.Count(value, "/") == 2:
		if length == 10 {
			layout = "2006/01/02"
		} else if length == 19 {
			layout = "2006/01/02 15:04:05"
		} else if length < 10 {
			layout = "2006/1/2"
		} else {
			layout = "2006/1/2 15:04:05"
		}
	case length > 0 && length < 20:
		str := "2006-01-02 15:04:05"
		layout = string(str[0:length])
		value = strings.Replace(value, "T", " ", 1)
	case length == 24:
		layout = "2006-01-02T15:04:05.999Z"
	case length == 25:
		layout = "2006-01-02T15:04:05-07:00"
	case length >= 26 && length <= 35:
		layout = "2006-01-02T15:04:05." + strings.Repeat("9", length-26) + "-07:00"
		value = strings.Replace(value, " ", "+", 1)
	}
	if layout != "" {
		if timestamp, err := time.Parse(layout, value); err != nil {
			log.Println(err)
		} else {
			return timestamp, true
		}
	}
	return time.Unix(0, 0), false
}

// StringifyTime 时间转换为字符串
func StringifyTime(t time.Time) string {
	layout := "2006-01-02T15:04:05.999999-07:00"
	return t.Format(layout)
}
