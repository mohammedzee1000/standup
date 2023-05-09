package config

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/cli/ptermutils"
	"github.com/mohammedzee1000/standup/pkg/config"
	"github.com/pterm/pterm"
	"strings"
	"time"
)

type Viewer interface {
	View(startOfWeekDay time.Weekday, holidays []string, sections []*config.ConfigSection, defaultSection, name string, sectionsPerRow int, configViewMode, reportViewMode uint) error
}

type PanelViewer struct{}

func (pv *PanelViewer) View(swd time.Weekday, holi []string, secs []*config.ConfigSection, dsec, nm string, spp int, configViewMode, reportViewMode uint) error {
	var sectionPanels = make(pterm.Panels, 0)
	var panelRow = make([]pterm.Panel, 0)
	var panels string

	pterm.DefaultSection.Printfln("Standup Configuration")

	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Name").Sprintf(nm)})
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Default Section").Sprintf(dsec)})
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Sections Per Row").Sprintf("%d", spp)})
	sectionPanels = append(sectionPanels, panelRow)

	panelRow = make([]pterm.Panel, 0)
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Start of Week Day").Sprintf(swd.String())})
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Holidays Every Week").Sprintf(strings.Join(holi, ","))})
	sectionPanels = append(sectionPanels, panelRow)

	panelRow = make([]pterm.Panel, 0)
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Config View Mode").Sprintf(config.ViewModeToString(configViewMode))})
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Report View Mode").Sprintf(config.ViewModeToString(reportViewMode))})
	sectionPanels = append(sectionPanels, panelRow)

	panels, _ = pterm.DefaultPanel.WithPanels(sectionPanels).Srender()
	pterm.DefaultBox.WithTitle("General Configuration").Println(panels)

	if len(secs) > 0 {
		var sectionTable = make(pterm.TableData, 0)
		sectionTable = append(sectionTable, []string{"Full Section Name", "Short", "Description"})
		for _, cs := range secs {
			sectionTable = append(sectionTable, []string{cs.Name, cs.Short, cs.Description})
		}
		ts, _ := pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(sectionTable).Srender()
		ts = fmt.Sprintf(
			"%s%s",
			pterm.Warning.Sprintfln("First name/short name will be used in case of repetition"),
			ts,
		)
		pterm.DefaultBox.WithTitle("Standup Sections").Println(ts)
	}
	return nil
}

type SimpleViewer struct {
}

func (pv *SimpleViewer) View(swd time.Weekday, holi []string, secs []*config.ConfigSection, dsec, nm string, spp int, configViewMode, reportViewMode uint) error {
	pterm.DefaultSection.Printfln("Standup Configuration")
	pterm.DefaultSection.WithLevel(2).Printfln("General Configuration")
	prefixSize := 19
	nmText := ptermutils.NewCustumInfoPrinter("Name", prefixSize).Sprintfln(nm)
	dsecText := ptermutils.NewCustumInfoPrinter("Default Section", prefixSize).Sprintfln(dsec)
	sppText := ptermutils.NewCustumInfoPrinter("Sections Per Row", prefixSize).Sprintfln("%d", spp)
	swdText := ptermutils.NewCustumInfoPrinter("Start of Week Day", prefixSize).Sprintfln("%s", swd.String())
	holiText := ptermutils.NewCustumInfoPrinter("Holidays Every Week", prefixSize).Sprintfln("%s", strings.Join(holi, ","))
	configViewModeText := ptermutils.NewCustumInfoPrinter("Config View Mode", prefixSize).Sprintfln(config.ViewModeToString(configViewMode))
	reportViewModeText := ptermutils.NewCustumInfoPrinter("Report View Mode", prefixSize).Sprintfln(config.ViewModeToString(reportViewMode))
	pterm.DefaultPanel.WithPanels(pterm.Panels{{{Data: pterm.DefaultBasicText.Sprintf("%s%s%s%s%s%s%s", nmText, dsecText, sppText, swdText, holiText, configViewModeText, reportViewModeText)}}}).Render()

	pterm.DefaultSection.WithLevel(2).Printfln("Standup Sections")
	prefixSize = 11
	if len(secs) > 0 {
		pterm.Warning.Printfln("First name/short name will be used in case of repetition")
		for _, cs := range secs {
			pterm.DefaultSection.WithLevel(3).Printfln("%s (SHORT: %s)", cs.Name, cs.Short)
			pterm.DefaultParagraph.Println(cs.Description)
		}
		fmt.Println()
	}

	return nil
}
