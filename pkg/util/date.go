package util

import (
	"fmt"
	"math"
	"time"
)

const (
	secondsInDayFactor   = 86400
	secondsIneWekFactor  = 604800
	secondsInMonthFactor = 2600640
	secondsInYearFactor  = 31207680
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

func DateToString(dt time.Time) string {
	tz, _ := dt.Zone()
	return fmt.Sprintf("%d %s %d %s %s", dt.Day(), dt.Month().String(), dt.Year(), dt.Weekday().String(), tz)
}

func GetDatesofWeek(firstDayWeek time.Weekday, refDate time.Time) [7]time.Time {
	dtl := [7]time.Time{}
	dt := refDate
	for dt.Weekday() != firstDayWeek {
		t := dt.AddDate(0, 0, -1)
		dt = t
	}
	for i := 0; i < 7; i++ {
		dtl[i] = dt.AddDate(0, 0, i)
	}
	return dtl
}

func RoundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

func DaysFromSeconds(s int) int {
	return RoundTime(float64(s / secondsInDayFactor))
}

func WeeksFromSeconds(s int) int {
	return RoundTime(float64(s / secondsIneWekFactor))
}

func MonthsFromSeconds(s int) int {
	return RoundTime(float64(s / secondsInMonthFactor))
}

func YearsFromSeconds(s int) int {
	return RoundTime(float64(s / secondsInYearFactor))
}
