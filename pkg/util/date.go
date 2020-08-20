package util

import (
	"fmt"
	"time"
)

func StringToWeekDay(value string) (wkday time.Weekday, err error) {
	var wkdaymap map[string]time.Weekday
	wkdaymap = make(map[string]time.Weekday)
	wkdaymap[time.Monday.String()] = time.Monday
	wkdaymap[time.Tuesday.String()] = time.Tuesday
	wkdaymap[time.Wednesday.String()] = time.Wednesday
	wkdaymap[time.Thursday.String()] = time.Thursday
	wkdaymap[time.Friday.String()] = time.Friday
	wkdaymap[time.Saturday.String()] = time.Saturday
	wkdaymap[time.Sunday.String()] = time.Sunday
	wkday, ok := wkdaymap[value]
	if !ok {
		return wkday, fmt.Errorf("invalid workday string %s", value)
	}
	return wkday, nil
}
