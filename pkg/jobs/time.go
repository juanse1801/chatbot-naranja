package jobs

import (
	"fmt"
	"time"
)

var months = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

func FormatDate(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprintf("%d", num)
}

type CronTime struct {
	Month  string
	Day    string
	Hour   string
	Minute string
}

func NewFirstTime() CronTime {
	nextTime := time.Now().Local().Add(time.Minute * 10)
	return CronTime{
		Month:  months[nextTime.Month().String()],
		Day:    FormatDate(nextTime.Day()),
		Hour:   FormatDate(nextTime.Hour()),
		Minute: FormatDate(nextTime.Minute()),
	}
}

func NewSecondTime() CronTime {
	nextTime := time.Now().Local().Add(time.Minute * 20)
	return CronTime{
		Month:  months[nextTime.Month().String()],
		Day:    FormatDate(nextTime.Day()),
		Hour:   FormatDate(nextTime.Hour()),
		Minute: FormatDate(nextTime.Minute()),
	}
}

func GetNextTime() string {
	nextTime := NewFirstTime()
	return fmt.Sprintf("%s %s %s %s *", nextTime.Minute, nextTime.Hour, nextTime.Day, nextTime.Month)
}

func GetSecondTime() string {
	nextTime := NewSecondTime()
	return fmt.Sprintf("%s %s %s %s *", nextTime.Minute, nextTime.Hour, nextTime.Day, nextTime.Month)
}
