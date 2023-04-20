package config

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/util"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameSet = "set"

type SetOptions struct {
	*common.CommonOptions
	defaultSection string
	holidays       []string
	startOfWeek    string
	name           string
}

func newSetOptions() *SetOptions {
	return &SetOptions{
		CommonOptions: common.NewCommonOptions(),
	}
}

func (so *SetOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := so.InitContext()
	if err != nil {
		return err
	}
	return nil
}

func (so *SetOptions) Validate() error {
	if so.defaultSection != "" && !so.Context.SectionExists(so.defaultSection) {
		return fmt.Errorf("default sections should exist in sections in config")
	}
	if so.startOfWeek != "" {
		_, err := util.StringToWeekDay(so.startOfWeek)
		if err != nil {
			return err
		}
	}
	if len(so.holidays) != 0 {
		for _, val := range so.holidays {
			_, err := util.StringToWeekDay(val)
			if err != nil {
				return fmt.Errorf("holdays should be valid weekdays: %w", err)
			}
		}
	}
	return nil
}

func (so *SetOptions) Run() (err error) {
	if so.defaultSection != "" {
		err = so.Context.SetDefaultSection(so.defaultSection)
		if err != nil {
			return err
		}
	}
	if so.startOfWeek != "" {
		err = so.Context.SetStartOfWeekDay(so.startOfWeek)
		if err != nil {
			return err
		}
	}
	if len(so.holidays) != 0 {
		err = so.Context.SetHolidays(so.holidays)
		if err != nil {
			return err
		}
	}
	if so.name != "" {
		err = so.Context.SetName(so.name)
		if err != nil {
			return err
		}
	}
	fmt.Println("Updated configuration")
	return nil
}

func NewCmdConfigSet(name, fullname string) *cobra.Command {
	o := newSetOptions()
	var configSetCmd *cobra.Command
	configSetCmd = &cobra.Command{
		Use:   name,
		Short: "Set basic config",
		Long:  "Set basic config values",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	configSetCmd.Flags().StringVarP(&o.defaultSection, "defaultsection", "d", "", "use to set default section")
	configSetCmd.Flags().StringVarP(&o.startOfWeek, "startofweekday", "w", "", "use to update start of week")
	configSetCmd.Flags().StringVarP(&o.name, "name", "n", "", "name of the owner of this standup")
	return configSetCmd
}
