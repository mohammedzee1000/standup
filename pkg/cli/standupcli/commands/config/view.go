package config

import (
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"strings"
)

const RecommendedCommandNameGet = "view"

type ViewOptions struct {
	*common.CommonOptions
}

func newViewOptions() *ViewOptions {
	return &ViewOptions{
		CommonOptions: common.NewCommonOptions(),
	}
}

func (goo *ViewOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	return goo.InitContext()
}

func (goo *ViewOptions) Validate() error {
	return nil
}

func (goo *ViewOptions) Run() error {
	//collect
	swd, err := goo.Context.GetStartOfWeekDay()
	if err != nil {
		return err
	}
	holi := goo.Context.GetHolidays()
	secs := goo.Context.GetSections()
	dsec := goo.Context.GetDefaultSection()
	nm := goo.Context.GetName()
	spp := goo.Context.GetSectionsPerRow()

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

	panels, _ = pterm.DefaultPanel.WithPanels(sectionPanels).Srender()
	pterm.DefaultBox.WithTitle("General Configuration").Println(panels)

	if len(secs) > 0 {
		var sectionTable = make(pterm.TableData, 0)
		sectionTable = append(sectionTable, []string{"Full Section Name", "Short", "Description"})
		for _, cs := range secs {
			sectionTable = append(sectionTable, []string{cs.Name, cs.Short, cs.Description})
		}
		ts, _ := pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(sectionTable).Srender()
		pterm.DefaultBox.WithTitle("Standup Sections").Println(ts)
	}
	return nil
}

func NewCmdConfigView(name, fullname string) *cobra.Command {
	o := newViewOptions()
	var configViewCmd *cobra.Command
	configViewCmd = &cobra.Command{
		Use:   name,
		Short: "View config",
		Long:  "View and view config values",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	return configViewCmd
}
