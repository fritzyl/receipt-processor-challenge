package utilities

import (
	"time"
)

const DateFormat string = "2006-01-02"
const TimeFormat string = "15:04"
const DayFormat string = "02"

func CompareTime(ts time.Time, oper string, comp time.Time) bool {
	if oper == "<" {
		return ts.Before(comp)
	} else if oper == ">" {
		return ts.After(comp)
	}
	return false
}

func GetTime(timeStr string) time.Time {
	ts, _ := time.Parse(TimeFormat, timeStr)
	return ts
}

func GetDate(dateStr string) time.Time {
	ts, _ := time.Parse(DateFormat, dateStr)
	return ts
}

func GetDay(dateStr string) time.Time {
	ts, _ := time.Parse(DayFormat, dateStr)
	return ts
}
