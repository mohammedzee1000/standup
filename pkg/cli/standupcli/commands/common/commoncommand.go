package common

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/util"
	"time"

	"github.com/mohammedzee1000/standup/pkg/config"
	"github.com/spf13/cobra"
)

type CommonOptions struct {
	Context *config.Context
}

func NewCommonOptions() *CommonOptions {
	return &CommonOptions{}
}

func (cc *CommonOptions) InitContext() (err error) {
	c, err := config.New()
	if err != nil {
		return err
	}
	cc.Context = c
	return nil
}

type DatedOptions struct {
	*CommonOptions
	day         int
	month       int
	monthActual time.Month
	year        int
	dt          *time.Time
}

func NewDatedOptions() *DatedOptions {
	return &DatedOptions{
		CommonOptions: NewCommonOptions(),
	}
}

func (do *DatedOptions) GetDate() time.Time {
	return *do.dt
}

func (do *DatedOptions) AddDateFlags(cmd *cobra.Command) {
	y, m, d := time.Now().Date()
	mi, _ := util.MonthToInt(m)
	cmd.Flags().IntVarP(&do.day, "day", "", d, "Day of the month")
	cmd.Flags().IntVarP(&do.month, "month", "", mi, "Number of Month")
	cmd.Flags().IntVarP(&do.year, "year", "", y, "the year")
}

func (do *DatedOptions) CompleteDate() error {
	var err error
	//now := time.Now()
	//if do.year <= 0 {
	//	do.year = now.Year() + do.year
	//}
	//if do.day <= 0 {
	//	minusDay := now.Day() + do.day
	//	if minusDay <= 0 {
	//		return fmt.Errorf("you can only use -ve day to go till the start of the month")
	//	}
	//	do.day = now.Day() + do.day
	//}
	//do.monthActual, err = util.StringToMonth(do.month)
	//if err != nil {
	//	return err
	//}
	//dt := time.Date(do.year, do.monthActual, do.day, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	//if time.Now().Before(dt) {
	//	return fmt.Errorf("cannot manipulate after today")
	//}
	//do.dt = &dt
	var dt, now time.Time
	now = time.Now()
	dtStr := fmt.Sprintf("%s-%s-%s", util.DateNumberToString(do.year), util.DateNumberToString(do.month), util.DateNumberToString(do.day))
	dt, err = time.Parse(time.DateOnly, dtStr)
	dt = time.Date(dt.Year(), dt.Month(), dt.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	if err != nil {
		return err
	}
	if dt.After(now) {
		return fmt.Errorf("cannot manipulete dates after today")
	}
	do.dt = &dt
	return nil
}
