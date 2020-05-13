package utils

import (
	"testing"
	"time"
)

type test struct {
	time    time.Time
	format  string
	strTime string
}

var testCases = []test{
	{
		time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
		"Y-m-d H:i:s",
		"2012-11-22 21:28:10",
	},
	{
		time.Date(2012, 11, 22, 0, 0, 0, 0, time.Local),
		"Y-m-d",
		"2012-11-22",
	},
	{
		time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
		"Y-m-d H:i:s",
		"2012-11-22 21:28:10",
	},
}

func TestFormat(t *testing.T) {
	for _, testCase := range testCases {
		strTime := Format(testCase.format, testCase.time)
		if strTime != testCase.strTime {
			t.Errorf("(expected) %v != %v (actual)", testCase.time, strTime)
		}
	}
}

func TestStrToLocalTime(t *testing.T) {
	for _, testCase := range testCases {
		time := StrToLocalTime(testCase.strTime)
		if time != testCase.time {
			t.Errorf("(expected) %v != %v (actual)", time, testCase.time)
		}
	}
}

func TestStrToTime(t *testing.T) {
	// zoneName, err := time.LoadLocation("CST")
	// if err != nil {
	//     t.Error(err)
	// }

	var testCases = []test{
		{
			time.Date(2012, 11, 22, 21, 28, 10, 0, time.Local),
			"",
			"2012-11-22 21:28:10 +0800 +0800",
		},
		{
			time.Date(2012, 11, 22, 0, 0, 0, 0, time.Local),
			"",
			"2012-11-22 +0800 +0800",
		},
		{
			time.Date(2012, 11, 22, 21, 28, 10, 0, time.FixedZone("CST", 28800)),
			"",
			"2012-11-22 21:28:10 +0800 CST",
		},
	}
	for _, testCase := range testCases {
		time := StrToTime(testCase.strTime)
		// if time != testCase.time {
		if !time.Equal(testCase.time) {
			t.Errorf("(expected) %v != %v (actual)", time, testCase.time)
		}
	}
}

func Test_Day(t *testing.T) {
	d := Day()
	i := Unix2Day(Timestamp())
	t.Log(d, i)
	if i != d {
		t.FailNow()
	}
}

func Test_TimeStr(t *testing.T) {
	Bdate := "2016-11-01 22:49:45" //时间字符串
	u, err := Str2Unix(Bdate)
	t.Log(u, err)
	l, err := Str2Local(Bdate)
	t.Log(l, err)

	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	t.Logf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)

	year, mon, day = now.Date()
	hour, min, sec = now.Clock()
	zone, _ = now.Zone()
	t.Logf("本地时间是 %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)
}

func Test_Time2Str(t *testing.T) {
	now := time.Now()
	t.Log(Time2Str(now))
}

func Test_Week(t *testing.T) {
	start, end := ThisWeek()
	t.Log(Time2Str(start), Time2Str(end))
	start, end = LastWeek()
	t.Log(Time2Str(start), Time2Str(end))
}
