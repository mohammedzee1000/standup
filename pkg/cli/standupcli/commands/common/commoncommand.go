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
	tnow  *time.Time
}

func NewDatedOptions() *DatedOptions {
	t := time.Now()
	return &DatedOptions{
		CommonOptions: NewCommonOptions(),
		tnow:          &t,
	}
}

func (do *DatedOptions) GetDate() time.Time {
	return *do.dt
}

func (do *DatedOptions) AddDateFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&do.day, "day", "", do.tnow.Day(), "Day of the month")
	cmd.Flags().StringVarP(&do.month, "month", "", do.tnow.Month().String(), "Name of Month")
	cmd.Flags().IntVarP(&do.year, "year", "", do.tnow.Year(), "year")
}

func (do *DatedOptions) CompleteDate() error {
	m, err := util.StringToMonth(do.month)
	if err != nil {
		return err
	}
	dt := time.Date(do.year, m, do.day, do.tnow.Hour(), do.tnow.Minute(), do.tnow.Second(), do.tnow.Nanosecond(), do.tnow.Location())
	if do.tnow.Before(dt) {
		return fmt.Errorf("cannot manipulate after today")
	}
	do.dt = &dt
	return nil
}
