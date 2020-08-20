package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/standup/standup"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameGet = "get"

type ReportGetOptions struct {
	*common.CommonOptions
	dt   time.Time
	week bool
}

func newReportGetOptions() *ReportGetOptions {
	var sgo *ReportGetOptions
	sgo = &ReportGetOptions{
		CommonOptions: common.NewCommonOptions(),
	}
	return sgo
}

func (sgo *ReportGetOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := sgo.InitContext()
	if err != nil {
		return err
	}
	sgo.dt = time.Now()
	return nil
	return nil
}

func (sgo *ReportGetOptions) Validate() error {
	return nil
}

func (sgo *ReportGetOptions) printWeekStandUp() error {
	//get first day of week
	firstDayWeek, err := sgo.Context.GetStartOfWeekDay()
	if err != nil {
		return err
	}
	dt := sgo.dt
	for dt.Weekday() != firstDayWeek {
		dt = dt.AddDate(0, 0, -1)
	}
	for dt.Weekday() != firstDayWeek-1 {
		var isHoliday bool
		if dt.Weekday() == sgo.dt.AddDate(0, 0, 1).Weekday() {
			fmt.Printf("----week still in progress/exceeded today----\n")
			break
		}

		for _, h := range sgo.Context.GetHolidays() {
			if dt.Weekday().String() == h {
				tz, _ := dt.Zone()
				fmt.Printf("Holiday on Date %d %s %d %s %s: \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
				isHoliday = true
			}
		}
		if !isHoliday {
			sgo.printStandUp(dt)
		}
		fmt.Println("")
		dt = dt.AddDate(0, 0, 1)
	}
	return nil
}

func (sgo *ReportGetOptions) printStandUp(dt time.Time) error {
	stc := standup.NewStandUpConfig(dt)
	e, err := stc.ConfigFileExists(sgo.Context)
	if err != nil {
		return fmt.Errorf("unable to check standup config exists %w", err)
	}
	if e {
		err = stc.FromFile(sgo.Context)
		if err != nil {
			return err
		}
		stup := stc.GetStandUp()
		tz, _ := dt.Zone()
		fmt.Printf("Standup for Date %d %s %d %s %s: \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
		for s, ts := range stup.Sections {
			fmt.Printf("%s:\n", strings.Title(s))
			desc := sgo.Context.GetSectionDescription(s)
			if desc != "" {
				fmt.Printf("Description: %s\n", desc)
			}

			for _, t := range ts {
				fmt.Printf("  - %s\n", t.Description)
			}
			fmt.Println("")
		}
	} else {
		tz, _ := dt.Zone()
		fmt.Printf("No Standup recorded for Date %d %s %d %s %s, skipping \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
	}
	return nil
}

func (sgo *ReportGetOptions) Run() error {
	if sgo.week {
		fmt.Printf("----Weekly Report----\n\n")
		//calculate first day of week
		errx := sgo.printWeekStandUp()
		fmt.Println("----end----")
		return errx
	} else {
		fmt.Printf("----Day Report----\n\n")
		errx := sgo.printStandUp(sgo.dt)
		fmt.Println("----end----")
		return errx
	}
}

func NewReportGetCommand(name, fullname string) *cobra.Command {
	o := newReportGetOptions()
	var reportGetCmd *cobra.Command
	reportGetCmd = &cobra.Command{
		Use:   name,
		Short: "Get report",
		Long:  "Get report for standup",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	reportGetCmd.Flags().BoolVarP(&o.week, "week", "w", false, "Standup for the week")
	return reportGetCmd
}
