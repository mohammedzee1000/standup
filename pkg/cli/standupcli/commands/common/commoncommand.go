package common

import (
	"fmt"
	"time"

	"github.com/mohammedzee1000/standup/pkg/config"
	"github.com/mohammedzee1000/standup/pkg/util"
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
	day   int
	month string
	year  int
	dt    *time.Time
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
	cmd.Flags().IntVarP(&do.day, "day", "", time.Now().Day(), "Day of the month")
	cmd.Flags().StringVarP(&do.month, "month", "", time.Now().Month().String(), "Name of Month")
	cmd.Flags().IntVarP(&do.year, "year", "", time.Now().Year(), "year")
}

func (do *DatedOptions) CompleteDate() error {
	m, err := util.StringToMonth(do.month)
	if err != nil {
		return err
	}
	dt := time.Date(do.year, m, do.day, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Now().Location())
	if time.Now().Before(dt) {
		return fmt.Errorf("cannot manipulate after today")
	}
	do.dt = &dt
	return nil
}
