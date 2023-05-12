package report

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/cli/ptermutils"
	"github.com/mohammedzee1000/standup/pkg/util"
	"github.com/pterm/pterm"
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
	for i := 0; i < 7; i++ {
		dt := datesOfWeek[i]
		if dt.After(time.Now()) {
			pterm.Info.Println("week still in progress/exceeded today")
			fmt.Println("")
			break
		}
		ro.printStandup(dt)
		fmt.Println("")
	}
	return nil
}

func (ro *ReportOptions) printStandup(dt time.Time) error {
	holidayMessage := fmt.Sprintf("Holiday on Date, skipping")
	noStandupMessage := fmt.Sprintf("No Standup recorded, skipping")
	stc := standup.NewStandUpConfig(dt)
	e, err := stc.ConfigFileExists(ro.Context)
	panel := pterm.DefaultBox.WithTitle(fmt.Sprintf("Standup Date %s", util.DateToString(dt)))

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
		for _, h := range ro.Context.GetHolidays() {
			if dt.Weekday().String() == h {
				isHoliday = true
				break
			}
		}
		if isHoliday || stup.IsHoliday {
			panel.Println(pterm.Info.Sprintfln(holidayMessage))
			return nil
		}
		if len(stup.Sections) == 0 {
			panel.Println(pterm.Info.Sprintfln(noStandupMessage))
			return nil
		}
		var sectionPanels = make(pterm.Panels, 0)
		var panelRow = make([]pterm.Panel, 0)
		for s, section := range stup.Sections {
			if len(panelRow) == 0 {
				panelRow = make([]pterm.Panel, 0)
			}
			var bli []pterm.BulletListItem
			for _, task := range section {
				bli = append(bli, pterm.BulletListItem{
					Level: 0,
					Text:  task.Description,
				})
			}
			blstr, _ := pterm.DefaultBulletList.WithItems(bli).Srender()
			blstr = strings.ReplaceAll(blstr, "%!(EXTRA <nil>)", "")
			if strings.HasSuffix(blstr, "\n") {
				blstr = strings.TrimRight(blstr, "\n")
			}
			descStr := ro.Context.GetSectionDescription(s)
			desc := ""
			if len(descStr) > 0 {
				desc = ptermutils.NewCustumInfoPrinter("Desc", 4).Sprintfln(ro.Context.GetSectionDescription(s))
			}
			blstr = fmt.Sprintf("%s\n%s", desc, blstr)
			panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle(s).Sprintf(blstr)})
			if len(panelRow) == ro.Context.GetSectionsPerRow() {
				sectionPanels = append(sectionPanels, panelRow)
				panelRow = make([]pterm.Panel, 0)
			}
		}
		if len(panelRow) != 0 {
			sectionPanels = append(sectionPanels, panelRow)
		}
		panels, _ := pterm.DefaultPanel.WithPanels(sectionPanels).Srender()
		panel.Println(panels)
	} else {
		panel.Println(pterm.Info.Sprintfln(noStandupMessage))
	}
	return nil
}

func (ro *ReportOptions) Run() error {
	pterm.DefaultSection.Println("Standup information")
	namePrinter := ptermutils.NewCustumInfoPrinter("Name", 4)
	standupTypePrinter := ptermutils.NewCustumInfoPrinter("Type", 4)
	namePrinter.Println(ro.Context.GetName())
	if ro.week {
		standupTypePrinter.Println("Weekly")
		pterm.DefaultSection.Println("Reports")
		errx := ro.printWeekStandUp()
		return errx
	} else {
		standupTypePrinter.Println("Specific Day")
		pterm.DefaultSection.Println("Report")
		errx := ro.printStandup(ro.GetDate())
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
