package common

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
)

type Runnable interface {
	Complete(name string, cmd *cobra.Command, args []string) error
	Validate() error
	Run() error
}

func logErrorAndExit(err error, context string, a ...interface{}) {
	if err != nil {
		if context == "" {
			pterm.Error.Printfln("%s\n", err)
		} else {
			pterm.Error.Printfln("%s %s", context, err)
		}
		os.Exit(1)
	}

}

func GenericRun(o Runnable, cmd *cobra.Command, args []string) {
	logErrorAndExit(o.Complete(cmd.Name(), cmd, args), "")
	logErrorAndExit(o.Validate(), "")
	logErrorAndExit(o.Run(), "")
}

func GetFullName(parentName, name string) string {
	return parentName + " " + name
}
