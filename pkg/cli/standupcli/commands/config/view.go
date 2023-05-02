package config

import (
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameGet = "view"

type ViewOptions struct {
	viewer Viewer
	*common.CommonOptions
}

func newViewOptions() *ViewOptions {
	return &ViewOptions{
		CommonOptions: common.NewCommonOptions(),
	}
}

func (goo *ViewOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := goo.InitContext()
	if err != nil {
		return err
	}
	goo.viewer = &PanelViewer{}
	return nil
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

	if goo.viewer != nil {
		return goo.viewer.View(swd, holi, secs, dsec, nm, spp)
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
