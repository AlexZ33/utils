package time

import (
	"github.com/AlexZ33/utils/errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

//timestamp be used to MySql timestamp converting.
type timestamp int64

// Scan scan time.
func (ts *timestamp) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case time.Time:
		*ts = timestamp(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*ts = timestamp(i)
	}
	return
}

// NowUnix 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

//GetTimeStr ...
func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FromUnix 秒时间戳转时间
func FromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// NowTimestamp 当前毫秒时间戳
func NowTimestamp() int64 {
	return Timestamp(time.Now())
}

// Timestamp 毫秒时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// FromTimestamp 毫秒时间戳转时间
func FromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// Format 时间格式化
func Format(time time.Time, layout string) string {
	return time.Format(layout)
}

// Parse 字符串时间转时间类型
func Parse(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

// GetDay return yyyyMMdd
func GetDay(time time.Time) int {
	ret, _ := strconv.Atoi(time.Format("20060102"))
	return ret
}

// WithTimeAsStartOfDay
// 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
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
//func (t Time) MarshalJSON() ([]byte, error) {
//	if t.IsZero() {
//		return []byte(""), nil
//	}
//}

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
