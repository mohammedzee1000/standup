package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/standup/standup"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameReport = "report"

type ReportOptions struct {
	*common.DatedOptions
	week bool
}

func newReportOptions() *ReportOptions {
	var sgo *ReportOptions
	sgo = &ReportOptions{
		DatedOptions: common.NewDatedOptions(),
	}
	return sgo
}

func (ro *ReportOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := ro.InitContext()
	if err != nil {
		return err
	}
	return ro.CompleteDate()
}

func (ro *ReportOptions) Validate() error {
	return nil
}

func (ro *ReportOptions) printWeekStandUp() error {
	//get first day of week
	firstDayWeek, err := ro.Context.GetStartOfWeekDay()
	if err != nil {
		return err
	}
	dt := ro.GetDate()
	for dt.Weekday() != firstDayWeek {
		t := dt.AddDate(0, 0, -1)
		dt = t
	}
	for dt.Weekday() != firstDayWeek-1 {
		var isHoliday bool
		if dt.Weekday() == ro.GetDate().AddDate(0, 0, 1).Weekday() {
			fmt.Printf("----week still in progress/exceeded today----\n")
			break
		}

		for _, h := range ro.Context.GetHolidays() {
			if dt.Weekday().String() == h {
				tz, _ := dt.Zone()
				fmt.Printf("Holiday on Date %d %s %d %s %s: \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
				isHoliday = true
			}
		}
		if !isHoliday {
			ro.printStandUp(dt)
		}
		fmt.Println("")
		dt = dt.AddDate(0, 0, 1)
	}
	return nil
}

func (ro *ReportOptions) printStandUp(dt time.Time) error {
	stc := standup.NewStandUpConfig(dt)
	e, err := stc.ConfigFileExists(ro.Context)
	if err != nil {
		return fmt.Errorf("unable to check standup config exists %w", err)
	}
	if e {
		err = stc.FromFile(ro.Context)
		if err != nil {
			return err
		}
		stup := stc.GetStandUp()
		tz, _ := dt.Zone()
		fmt.Printf("Standup for Date %d %s %d %s %s: \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
		for s, ts := range stup.Sections {
			fmt.Printf("%s:\n", strings.Title(s))
			desc := ro.Context.GetSectionDescription(s)
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

func (ro *ReportOptions) Run() error {
	if ro.week {
		fmt.Printf("----Weekly Report----\n\n")
		errx := ro.printWeekStandUp()
		fmt.Println("----end----")
		return errx
	} else {
		fmt.Printf("----Day Report----\n\n")
		errx := ro.printStandUp(ro.GetDate())
		fmt.Println("----end----")
		return errx
	}
}

func NewCmdReport(name, fullname string) *cobra.Command {
	o := newReportOptions()
	var reportCommand *cobra.Command
	reportCommand = &cobra.Command{
		Use:   name,
		Short: "get reports",
		Long:  "get reports for standup",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	o.AddDateFlags(reportCommand)
	reportCommand.Flags().BoolVarP(&o.week, "week", "w", false, "use to get week report")
	return reportCommand
}
