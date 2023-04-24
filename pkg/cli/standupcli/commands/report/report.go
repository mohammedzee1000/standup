package report

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/util"
	"time"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/standup/standup"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameReport = "report"

type ReportOptions struct {
	*common.DatedOptions
	week bool
	wide bool
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

func printHolidayMessage(dt time.Time) {
	tz, _ := dt.Zone()
	fmt.Printf("Holiday on Date %d %s %d %s %s: \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
}

func printNoStandupMessage(dt time.Time) {
	tz, _ := dt.Zone()
	fmt.Printf("No Standup recorded for Date %d %s %d %s %s, skipping \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
}

func (ro *ReportOptions) printWeekStandUp() error {
	//get first day of week
	firstDayWeek, err := ro.Context.GetStartOfWeekDay()
	if err != nil {
		return err
	}
	datesOfWeek := util.GetDatesofWeek(firstDayWeek, ro.GetDate())
	fmt.Printf("Name: %s\n", ro.Context.GetName())
	for i := 0; i < 7; i++ {
		dt := datesOfWeek[i]
		ro.printStandUp(dt, true)
		fmt.Println("")

		if dt.After(time.Now()) {
			fmt.Printf("----week still in progress/exceeded today----\n")
			break
		}
	}
	return nil
}

func (ro *ReportOptions) printStandUp(dt time.Time, printName bool) error {
	stc := standup.NewStandUpConfig(dt)
	e, err := stc.ConfigFileExists(ro.Context)
	if err != nil {
		return fmt.Errorf("unable to check standup config exists %w", err)
	}
	if e {
		isHoliday := false
		err = stc.FromFile(ro.Context)
		if err != nil {
			return err
		}
		stup := stc.GetStandUp()
		tz, _ := dt.Zone()
		for _, h := range ro.Context.GetHolidays() {
			if dt.Weekday().String() == h {
				isHoliday = true
				break
			}
		}
		if isHoliday || stup.IsHoliday {
			printHolidayMessage(dt)
			return nil
		}
		if len(stup.Sections) == 0 {
			printNoStandupMessage(dt)
			return nil
		}
		fmt.Printf("Standup for Date %d %s %d %s %s: \n\n", dt.Day(), dt.Month(), dt.Year(), dt.Weekday(), tz)
		if printName {
			fmt.Printf("Name: %s\n", ro.Context.GetName())
		}
		for s, ts := range stup.Sections {
			fmt.Printf("%s:\n", s)
			desc := ro.Context.GetSectionDescription(s)
			if desc != "" {
				fmt.Printf("Description: %s\n", desc)
			}

			for _, t := range ts {
				ids := ""
				if ro.wide {
					ids = fmt.Sprintf("[%s] ", t.ID)
				}
				fmt.Printf("  - %s%s\n", ids, t.Description)
			}
			fmt.Println("")
		}
	} else {
		printNoStandupMessage(dt)
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
		errx := ro.printStandUp(ro.GetDate(), true)
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
	reportCommand.Flags().BoolVar(&o.wide, "wide", false, "show wide/full standup")
	return reportCommand
}
