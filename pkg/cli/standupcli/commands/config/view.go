package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
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
	sec := goo.Context.GetSections()
	dsec := goo.Context.GetDefaultSection()
	nm := goo.Context.GetName()

	//out
	w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "Configuration Parameters:\n ")
	fmt.Fprintln(w, fmt.Sprintf("Name: %s", nm))
	fmt.Fprintln(w, "Sections:")
	for sn, sval := range sec {
		fmt.Fprintf(w, "\t%s: %s\n", sn, sval)
	}
	fmt.Fprintf(w, "Default Section: %s\n", dsec)
	fmt.Fprintf(w, "Start of weekday:\t%s\n", swd.String())
	fmt.Fprintln(w, "Holidays:")
	for _, val := range holi {
		fmt.Fprintf(w, "\t-%s\n", val)
	}
	w.Flush()
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
