package gutil

import (
	"fmt"
	"testing"
	"time"
)

var now = time.Now()

func TestTimeToDefaultStr(t *testing.T) {
	t.Log(TimeToDefaultStr(now))
}

func TestTimeToStr(t *testing.T) {
	t.Log(TimeToStr(now, YearMonthDay))
}

func TestTimeToUnixSecond(t *testing.T) {
	t.Log(TimeToUnixSecond(now))
}

func TestTimeToUnixMilli(t *testing.T) {
	t.Log(TimeToUnixMilli(now))
}

func TestTimeToUnixNano(t *testing.T) {
	t.Log(TimeToUnixNano(now))
}

func TestUnixSecondToTime(t *testing.T) {
	t.Log(UnixSecondToTime(now.Unix()))
}

func TestUnixMilliToTime(t *testing.T) {
	t.Log(UnixMilliToTime(now.Unix()*1000 + 511))
}

func TestUnixNanoToTime(t *testing.T) {
	t.Log(UnixNanoToTime(now.UnixNano()))
}

func TestUnixSecondToDefaultTimeStr(t *testing.T) {
	t.Log(UnixSecondToDefaultTimeStr(now.Unix()))
}

func TestUnixMilliToDefaultTimeStr(t *testing.T) {
	t.Log(UnixMilliToDefaultTimeStr(now.Unix() * 1000))
}

func TestUnixNanoToDefaultTimeStr(t *testing.T) {
	t.Log(UnixNanoToDefaultTimeStr(now.UnixNano()))
}

func TestUnixSecondToTimeStr(t *testing.T) {
	t.Log(UnixSecondToTimeStr(now.Unix(), YearMonthDay))
}

func TestUnixMilliToTimeStr(t *testing.T) {
	t.Log(UnixMilliToTimeStr(now.Unix()*1000, YearMonthDay))
}

func TestUnixNanoToTimeStr(t *testing.T) {
	t.Log(UnixNanoToTimeStr(now.UnixNano(), YearMonthDay))
}

func TestTimeStrToTimeDefault(t *testing.T) {
	t.Log(TimeStrToTimeDefault("2021-02-04 18:00:00"))
}

func TestTimeStrToTime(t *testing.T) {
	t.Log(TimeStrToTime("2021-02-0418:00:00", YearMonthDay+HourMinuteSecond))
}

func TestTimeStrToUnixSecondDefault(t *testing.T) {
	t.Log(TimeStrToUnixSecondDefault("2021-02-04 18:00:00"))
}

func TestTimeStrToUnixMilliDefault(t *testing.T) {
	t.Log(TimeStrToUnixMilliDefault("2021-02-04 18:00:00"))
}

func TestTimeStrToUnixNanoDefault(t *testing.T) {
	t.Log(TimeStrToUnixNanoDefault("2021-02-04 18:00:00"))
}

func TestTimeStrToUnixSecond(t *testing.T) {
	t.Log(TimeStrToUnixSecond("2021-02-0418:00:00", YearMonthDay+HourMinuteSecond))
}

func TestTimeStrToUnixMilli(t *testing.T) {
	t.Log(TimeStrToUnixMilli("2021-02-0418:00:00.521", YearMonthDay+HourMinuteSecond+".000"))
}

func TestTimeStrToUnixNano(t *testing.T) {
	t.Log(TimeStrToUnixNano("2021-02-0418:00:00.521123134", YearMonthDay+HourMinuteSecond+".000000000"))
}

func TestNextDayStartedAtUnix(t *testing.T) {
	fmt.Println(NextDayStartedAtUnix())
}

func TestNextDayStartedAt(t *testing.T) {
	fmt.Println(NextDayStartedAt())
}

func TestTodayStartedAtUnix(t *testing.T) {
	fmt.Println(TodayStartedAtUnix())
}

func TestTodayStartedAt(t *testing.T) {
	fmt.Println(TodayStartedAt())
}
