package config

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/config"
	"github.com/pterm/pterm"
	"strings"
	"time"
)

type Viewer interface {
	View(startOfWeekDay time.Weekday, holidays []string, sections []*config.ConfigSection, defaultSection, name string, sectionsPerRow int) error
}

type PanelViewer struct{}

func (pv *PanelViewer) View(swd time.Weekday, holi []string, secs []*config.ConfigSection, dsec, nm string, spp int) error {
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
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Holidays").Sprintf(strings.Join(holi, ","))})
	sectionPanels = append(sectionPanels, panelRow)

	panelRow = make([]pterm.Panel, 0)
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Config View Mode").Sprintf("todo")})
	panelRow = append(panelRow, pterm.Panel{Data: pterm.DefaultBox.WithTitle("Report View Mode").Sprintf("todo")})
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
