package markholiday

import (
	"fmt"
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/standup/standup"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameMarkHoliday = "markholiday"

type MarkHolidayOptions struct {
	*common.DatedOptions
	isHoliday bool
}

func newMarkHolidayOptions() *MarkHolidayOptions {
	var mho *MarkHolidayOptions
	mho = &MarkHolidayOptions{
		DatedOptions: common.NewDatedOptions(),
		isHoliday:    false,
	}
	return mho
}

func (mho *MarkHolidayOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := mho.InitContext()
	if err != nil {
		return err
	}
	return mho.CompleteDate()
}

func (mho *MarkHolidayOptions) Validate() error {
	return nil
}

func (mho *MarkHolidayOptions) Run() error {
	stc := standup.NewStandUpConfig(mho.GetDate())
	err := stc.FromFile(mho.Context)
	if err != nil {
		return err
	}
	stc.GetStandUp().IsHoliday = mho.isHoliday
	err = stc.ToFile(mho.Context)
	if err != nil {
		return err
	}
	markStr := "successfully marked date as "
	if !mho.isHoliday {
		markStr = fmt.Sprintf("%s not ", markStr)
	}
	pterm.Success.Printfln("%sa holiday", markStr)
	return nil
}

func NewCmdMarkHoliday(name, fullname string) *cobra.Command {
	o := newMarkHolidayOptions()
	var markHolidayCmd *cobra.Command
	markHolidayCmd = &cobra.Command{
		Use:   name,
		Short: "Mark specific day as holiday.",
		Long:  "Mark specific day as holiday (outside of regular holidays marked in config)",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	o.AddDateFlags(markHolidayCmd)
	markHolidayCmd.Flags().BoolVarP(&o.isHoliday, "holiday", "m", false, "is this day a holiday.")
	return markHolidayCmd
}
