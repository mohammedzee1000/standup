package config

import (
	"fmt"
	"github.com/pterm/pterm"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameUpdateSection = "updatesection"

type UpdateSectionOptions struct {
	*common.CommonOptions
	sectionName           string
	newSectionDescription string
	deleteSection         bool
	newSectionName        string
	newShortName          string
}

func newUpdateSectionOptions() *UpdateSectionOptions {
	return &UpdateSectionOptions{
		CommonOptions: common.NewCommonOptions(),
	}
}

func (uso *UpdateSectionOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := uso.InitContext()
	if err != nil {
		return err
	}
	if sn := uso.Context.GetSectionNameByShortName(uso.sectionName); sn != "" {
		uso.sectionName = sn
	}
	return nil
}

func (uso *UpdateSectionOptions) Validate() error {
	if uso.sectionName == "" {
		return fmt.Errorf("please provide name of section to update")
	}
	if uso.deleteSection && (uso.newSectionDescription != "" || uso.newShortName != "") {
		return fmt.Errorf("requesting for deletion, please do not provide any other parameter")
	}
	if !uso.deleteSection && uso.newSectionName == "" && uso.newShortName == "" && uso.newSectionDescription == "" {
		return fmt.Errorf("please provide a field that you wish to update")
	}
	return nil
}

func (uso *UpdateSectionOptions) Run() (err error) {
	if uso.deleteSection {
		err = uso.Context.DeleteSection(uso.sectionName)
		if err != nil {
			return err
		}
		pterm.Success.Printfln("section %s deleted successfully", uso.sectionName)
		return nil
	}

	if uso.newSectionDescription != "" {
		err = uso.Context.UpdateSectionDescription(uso.sectionName, uso.newSectionDescription)
		if err != nil {
			return err
		}
	}
	if uso.newSectionName != "" {
		err = uso.Context.UpdateSectionName(uso.sectionName, uso.newSectionName)
		if err != nil {
			return err
		}
	}
	if uso.newShortName != "" {
		err = uso.Context.UpdateSectionShortName(uso.sectionName, uso.newShortName)
		if err != nil {
			return err
		}
	}
	pterm.Success.Printfln("section %s updated successfully", uso.sectionName)
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
	configUpdateSectionCmd.Flags().StringVarP(&o.sectionName, "name", "n", "", "current name of sections to refer update")
	configUpdateSectionCmd.Flags().StringVarP(&o.newSectionDescription, "description", "d", "", "description of section")
	configUpdateSectionCmd.Flags().BoolVar(&o.deleteSection, "delete", false, "delete specified section")
	configUpdateSectionCmd.Flags().StringVarP(&o.newSectionName, "changename", "c", "", "new name of the section")
	configUpdateSectionCmd.Flags().StringVarP(&o.newShortName, "changeshortname", "s", "", "new short name of the section")
	return configUpdateSectionCmd
}
