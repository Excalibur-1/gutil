// Package gutil 封装常用go工具类
package gutil

import (
	"context"
	"time"
)

const (
	YearMonthDay     = "2006-01-02"
	HourMinuteSecond = "15:04:05"
	DefaultLayout    = YearMonthDay + " " + HourMinuteSecond
)

const (
	Zero = 0
)

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration time.Duration

// NewDuration create a Duration
func NewDuration(str string) (dur Duration) {
	tmp, err := time.ParseDuration(str)
	if err == nil {
		dur = Duration(tmp)
	}
	return
}

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if timeout := time.Until(deadline); timeout < time.Duration(d) {
			// deliver small timeout
			return Duration(timeout), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, time.Duration(d))
	return d, ctx, cancel
}

// TimeToDefaultStr time转为默认格式的日期字符串
func TimeToDefaultStr(t time.Time) string {
	return TimeToStr(t, DefaultLayout)
}

// TimeToStr time转换为指定格式的日期字符串
func TimeToStr(t time.Time, layout string) string {
	return t.Format(layout)
}

// TimeToUnixSecond time转换为秒级时间戳
func TimeToUnixSecond(t time.Time) int64 {
	return t.Unix()
}

// TimeToUnixMilli time转换为毫秒级别时间戳
func TimeToUnixMilli(t time.Time) int64 {
	return TimeToUnixNano(t) / 1000
}

// TimeToUnixNano time转换为纳秒级别时间戳
func TimeToUnixNano(t time.Time) int64 {
	return t.UnixNano()
}

// UnixSecondToTime 秒级时间戳转time
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// UnixMilliToTime 毫秒级时间戳转time
func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(milli/1000, (milli%1000)*(1000*1000))
}

// UnixNanoToTime 纳秒级时间戳转time
func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(nano/(1000*1000*1000), nano%(1000*1000*1000))
}

// UnixSecondToDefaultTimeStr 秒级时间戳转默认格式日期字符串
func UnixSecondToDefaultTimeStr(second int64) string {
	return UnixSecondToTimeStr(second, DefaultLayout)
}

// UnixMilliToDefaultTimeStr 毫秒级时间戳转默认格式日期字符串
func UnixMilliToDefaultTimeStr(milli int64) string {
	return UnixMilliToTimeStr(milli, DefaultLayout)
}

// UnixNanoToDefaultTimeStr 纳秒级时间戳转默认格式日期字符串
func UnixNanoToDefaultTimeStr(nano int64) string {
	return UnixNanoToTimeStr(nano, DefaultLayout)
}

// UnixSecondToTimeStr 秒级时间戳转指定格式日期字符串
func UnixSecondToTimeStr(second int64, layout string) string {
	return TimeToStr(UnixSecondToTime(second), layout)
}

// UnixMilliToTimeStr 毫秒级时间戳转指定格式日期字符串
func UnixMilliToTimeStr(second int64, layout string) string {
	return TimeToStr(UnixMilliToTime(second), layout)
}

// UnixNanoToTimeStr 纳秒级时间戳转指定格式日期字符串
func UnixNanoToTimeStr(nano int64, layout string) string {
	return TimeToStr(UnixNanoToTime(nano), layout)
}

// TimeStrToTimeDefault 默认格式日期字符串转time
func TimeStrToTimeDefault(str string) time.Time {
	parseTime, _ := time.ParseInLocation(DefaultLayout, str, time.Local)
	return parseTime
}

// TimeStrToTime 指定格式日期字符串转time
func TimeStrToTime(str, layout string) time.Time {
	parseTime, _ := time.ParseInLocation(layout, str, time.Local)
	return parseTime
}

// TimeStrToUnixSecondDefault 默认格式日期字符串转秒时间戳
func TimeStrToUnixSecondDefault(str string) int64 {
	return TimeStrToTimeDefault(str).Unix()
}

// TimeStrToUnixMilliDefault 默认格式日期字符串转毫秒时间戳
func TimeStrToUnixMilliDefault(str string) int64 {
	return TimeStrToUnixNanoDefault(str) / 1000
}

// TimeStrToUnixNanoDefault 默认格式日期字符串转纳秒时间戳
func TimeStrToUnixNanoDefault(str string) int64 {
	return TimeStrToTimeDefault(str).UnixNano()
}

// TimeStrToUnixSecond 指定格式日期字符串转秒时间戳
func TimeStrToUnixSecond(str, layout string) int64 {
	return TimeStrToTime(str, layout).Unix()
}

// TimeStrToUnixMilli 指定格式日期字符串转毫秒时间戳
func TimeStrToUnixMilli(str, layout string) int64 {
	return TimeStrToUnixNano(str, layout) / (1000 * 1000)
}

// TimeStrToUnixNano 指定格式日期字符串转纳秒时间戳
func TimeStrToUnixNano(str, layout string) int64 {
	return TimeStrToTime(str, layout).UnixNano()
}

// NextDayStartedAtUnix 获取下一个零点时间戳
func NextDayStartedAtUnix() int64 {
	return NextDayStartedAt().Unix()
}

// NextDayStartedAt 获取下一个零点（明天的零点）
func NextDayStartedAt() time.Time {
	return TodayStartedAt().Add(time.Hour * 24)
}

// TodayStartedAtUnix 获取今天零点时间戳
func TodayStartedAtUnix() int64 {
	return TodayStartedAt().Unix()
}

// TodayStartedAt 获取今天零点
func TodayStartedAt() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), Zero, Zero, Zero, Zero, now.Location())
}
