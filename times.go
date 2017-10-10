/**********************************************************
 * Author        :
 * Email         :
 * Last modified : 2016-09-15 16:22:18
 * Filename      : times.go
 * Description   : 常用的时间工具方法
 * *******************************************************/
package utils

import (
	"fmt"
	"strings"
	"time"
)

// Format 跟 PHP 中 date 类似的使用方式，如果 ts 没传递，则使用当前时间
func Format(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

func StrToLocalTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return StrToTime(value)
}

func StrToTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}
	panic(err)
}

const FORMAT string = "2006-01-02 15:04:05"
const FORMATDATA string = "2006-01-02 "
const FORMATSTR string = "2006/01/02/15"

/**
 * 获取当前纳秒时间截
 * @return int64
 */
func TimestampNano() int64 {
	return now().UnixNano()
}

/**
 * 获取当前时间截
 * @return int64
 */
func Timestamp() int64 {
	return now().Unix()
}

/**
 * 获取本地当天零点时间截
 * @return int64
 */
func TimestampToday() int64 {
	return TimestampTodayTime().Unix()
}

/**
 * 获取本地当天零点时间截
 * @return time.Time
 */
func TimestampTodayTime() time.Time {
	now := now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

/**
 * 返回时间截数值
 * @return int64
 */
func Time2Stamp(t time.Time) int64 {
	return t.Unix()
}

/**
 * 返回时间
 * @return int64
 */
func Stamp2Time(t int64) time.Time {
	return unix(t, 0)
}

/**
 * 获取本地当天时间20170402
 * @return int
 */
func DayDate() int {
	now := now()
	return now.Year()*10000 + int(now.Month())*100 + now.Day()
}

/**
 * 获取本地当月时间201704
 * @return int
 */
func MonthDate() int {
	now := now()
	return now.Year()*10000 + int(now.Month())*100
}

/**
 * 获取本地昨天零点时间截
 * @return int64
 */
func TimestampYesterday() int64 {
	return TimestampToday() - 86400
}

/**
 * 获取本地明天零点时间截
 * @return int64
 */
func TimestampTomorrow() int64 {
	return TimestampToday() + 86400
}

/**
 * 获取当前年,月,日
 * @return int
 */
func DateTime() (int, time.Month, int) {
	year, month, day := now().Date()
	return year, month, day
}

/**
 * 获取相对年,月,日
 * @return int
 * Some examples: AddDateTime(0, -1, 0) 上月今天时间
 */
func AddDateTime(year, month, day int) (int, time.Month, int) {
	return now().AddDate(year, month, day).Date()
}

/**
 * 获取当前年
 * @return int
 */
func Year() int {
	return now().Year()
}

/**
 * 获取当前月
 * @return time.Month
 */
func Month() time.Month {
	return now().Month()
}

/**
 * 获取当前天
 * @return int
 */
func Day() int {
	return now().Day()
}

/**
 * 获取当前时间小时
 * @return int
 */
func Hour() int {
	return now().Hour()
}

/**
 * 获取当前时间分钟
 * @return int
 */
func Minute() int {
	return now().Minute()
}

/**
 * 获取当前时间秒
 * @return int
 */
func Second() int {
	return now().Second()
}

/**
 * 获取当前时间纳秒
 * @return int
 */
func Nanosecond() int {
	return now().Nanosecond()
}

/**
 * 获取当前周
 * @return time.Weekday
 */
func Weekday() time.Weekday {
	return now().Weekday()
}

/**
 * 是否当前时间之前
 * @return bool
 */
func Before() bool {
	then := date(2009, 11, 17, 20, 34, 58, 651387237)
	return then.Before(now())
}

/**
 * 是否当前时间之后
 * @return bool
 */
func After() bool {
	then := date(2009, 11, 17, 20, 34, 58, 651387237)
	return then.After(now())
}

/**
 * 是否等于当前时间
 * @return bool
 */
func Equal() bool {
	then := date(2009, 11, 17, 20, 34, 58, 651387237)
	return then.Equal(now())
}

/**
 * 两个时间点的时间间隔
 * @return time.Duration
 */
func Diff() time.Duration {
	then := date(2009, 11, 17, 20, 34, 58, 651387237)
	return now().Sub(then)
	// diff := now.Sub(then)
	// diff.Hours()
	// diff.Minutes()
	// diff.Seconds()
	// diff.Nanoseconds()
	// then.Add(diff)
	// then.Add(-diff)
}

/**
 * 获取指定时间截的年
 * @param t int64
 * @return int
 */
func Unix2Year(t int64) int {
	return unix(t, 0).Year()
}

/**
 * 获取指定时间截的月
 * @param t int64
 * @return time.Month
 */
func Unix2Month(t int64) time.Month {
	return unix(t, 0).Month()
}

/**
 * 获取指定时间截的天
 * @param t int64
 * @return int
 */
func Unix2Day(t int64) int {
	return unix(t, 0).Day()
}

/**
 * 时间戳转str格式化时间
 * @param t int64
 * @return string
 */
func Unix2Str(t int64) string {
	return unix(t, 0).Format(FORMAT)
}

/**
 * 时间戳转str格式化时间
 * @param t int64
 * @return string
 */
func Unix2Fstr(t int64) string {
	return unix(t, 0).Format(FORMATSTR)
}

/**
 * str格式当前日期
 * @return string
 */
func DateStr() string {
	return now().Format(FORMATDATA)
}

/**
 * str格式化时间转时间戳
 * @param t string
 * @return (int64, error)
 */
func Str2Unix(t string) (int64, error) {
	the_time, err := time.Parse(FORMAT, t)
	if err == nil {
		return the_time.Unix(), err
	}
	return 0, err
}

/**
 * str格式化时间转本地时间戳
 * @param t string
 * @return (int64, error)
 */
func Str2Local(t string) (int64, error) {
	the_time, err := time.ParseInLocation(FORMAT, t, time.Local)
	if err == nil {
		return the_time.Unix(), err
	}
	return 0, err
}

/**
 * str格式化时间转时间戳
 * @param t string
 * @return time.Time
 */
func Str2Time(t string) time.Time {
	time, _ := Str2Local(t)
	return Stamp2Time(time)
}

/**
 * 获取指定年月的天数
 * @param year int
 * @param month int
 * @return days int
 */
func MonthDays(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

/**
 * Now returns the current local time
 * @return time.Time struct{}
 */
func LocalTime() time.Time {
	return now().Local()
}

/**
 * Now returns the current local time
 * @return time.Time struct{}
 */
func now() time.Time {
	return time.Now()
}

// Now returns the current time with millisecond precision. MongoDB stores
// timestamps with the same precision, so a Time returned from this method
// will not change after a roundtrip to the database. That's the only reason
// why this function exists. Using the time.Now function also works fine
// otherwise.
func BsonNow() time.Time {
	return time.Unix(0, time.Now().UnixNano()/1e6*1e6)
}

/**
 * Date returns the current local time
 * @return time.Time struct{}
 */
func date(year, month, day, hour, min, sec, nsec int) time.Time {
	Month := time.Month(month)
	return time.Date(year, Month, day, hour, min, sec, nsec, time.UTC)
}

/**
 * Unix returns the local Time corresponding to the given Unix time
 * @param sec int64
 * @param nsec int64
 * @return time.Time struct{}
 */
func unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

/**
 * 延迟second
 * @param second int
 */
func Sleep(second int) {
	<-time.After(time.Duration(second) * time.Second)
}

/**
 * 延迟1~second
 * @param second int
 */
func SleepRand(second int) {
	<-time.After(time.Duration(RandIntN(second)+1) * time.Second)
}

/**
 * 延迟second
 * @param second int64
 */
func Sleep64(second int64) {
	<-time.After(time.Duration(second) * time.Second)
}

/**
 * 延迟1~second
 * @param second int64
 */
func SleepRand64(second int64) {
	<-time.After(time.Duration(RandInt64N(second)+1) * time.Second)
}

/**
 * 指定second时间内每隔per执行一次
 * @param per int
 * @param second int
 * @param f func()
 */
func Kickers(per, second int, f func()) {
	ticker := time.NewTicker(time.Duration(per) * time.Second)
	go func() {
		for _ = range ticker.C {
			f()
		}
	}()
	time.Sleep(time.Duration(second) * time.Second)
	ticker.Stop()
}

/**
 * 时钟
 * @param second int
 * @return *time.Ticker
 */
func NewTicker(second int) *time.Ticker {
	return time.NewTicker(time.Duration(second) * time.Second)
}

/**
 * 时钟
 * @param millisecond int
 * @return *time.Ticker
 */
func NewTickerMilli(millisecond int) *time.Ticker {
	return time.NewTicker(time.Duration(millisecond) * time.Millisecond)
}

/**
 * 定时器
 * @param second int
 * @param f func()
 * @return *time.Timer
 */
func TimerRun(second int, f func()) *time.Timer {
	timer := time.NewTimer(time.Duration(second) * time.Second)
	go func() {
		<-timer.C
		f()
	}()
	return timer
}

/**
 * 取消定时器
 * @param *time.Timer
 * @return bool
 */
func TimerStop(timer *time.Timer) bool {
	if timer.Stop() {
		return true
	}
	return false
}

// func TimeOut(f func()) bool {
// 	c := make(chan string, 1)
// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		c <- "result"
// 	}()
// 	select {
// 	case res := <-c:
// 		f()
// 		return true
// 	case <-time.After(time.Second * 3):
// 		return false
// 	}
// }
