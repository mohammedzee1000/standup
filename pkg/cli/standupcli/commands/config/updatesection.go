package config

import (
	"fmt"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameUpdateSection = "updatesection"

type UpdateSectionOptions struct {
	*common.CommonOptions
	sectionName        string
	sectionDescription string
	deleteSection      bool
}

func newUpdateSectionOptions() *UpdateSectionOptions {
	return &UpdateSectionOptions{
		CommonOptions: common.NewCommonOptions(),
	}
}

func (uso *UpdateSectionOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	return uso.InitContext()
}

func (uso *UpdateSectionOptions) Validate() error {
	if uso.sectionName == "" {
		return fmt.Errorf("please provide name of section to update")
	}
	if !uso.deleteSection && uso.sectionDescription == "" {
		return fmt.Errorf("please provide section description")
	}
	return nil
}

func (uso *UpdateSectionOptions) Run() (err error) {
	if uso.deleteSection {
		err = uso.Context.DeleteSection(uso.sectionName)
		if err != nil {
			return err
		}
		fmt.Println("deleted successfully")
		return nil
	}
	err = uso.Context.UpdateSectionDescription(uso.sectionName, uso.sectionDescription)
	if err != nil {
		return err
	}
	fmt.Println("updated successfully")
	return nil
}

func NewCmdConfigUpdateSection(name, fullname string) *cobra.Command {
	o := newUpdateSectionOptions()
	var configUpdateSectionCmd *cobra.Command
	configUpdateSectionCmd = &cobra.Command{
		Use:   name,
		Short: "update or delete sections in config",
		Long:  "update or delete sections in config",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	configUpdateSectionCmd.Flags().StringVarP(&o.sectionName, "name", "n", "", "name of sections")
	configUpdateSectionCmd.Flags().StringVarP(&o.sectionDescription, "description", "d", "", "description of section")
	configUpdateSectionCmd.Flags().BoolVar(&o.deleteSection, "delete", false, "delete specified section")
	return configUpdateSectionCmd
}
