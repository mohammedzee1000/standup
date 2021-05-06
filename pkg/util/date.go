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
		return wkday, fmt.Errorf("invalid weekday string, try `Monday` or `Tuesday` and so on %s", value)
	}
	return wkday, nil
}

func StringToMonth(value string) (mnt time.Month, err error) {
	var mntmap map[string]time.Month
	mntmap = make(map[string]time.Month)
	mntmap[time.January.String()] = time.January
	mntmap[time.February.String()] = time.February
	mntmap[time.March.String()] = time.March
	mntmap[time.April.String()] = time.April
	mntmap[time.May.String()] = time.May
	mntmap[time.June.String()] = time.June
	mntmap[time.July.String()] = time.July
	mntmap[time.August.String()] = time.August
	mntmap[time.September.String()] = time.September
	mntmap[time.October.String()] = time.October
	mntmap[time.November.String()] = time.November
	mntmap[time.December.String()] = time.December
	mnt, ok := mntmap[value]
	if !ok {
		return mnt, fmt.Errorf("invalid month")
	}
	return mnt, nil
}
