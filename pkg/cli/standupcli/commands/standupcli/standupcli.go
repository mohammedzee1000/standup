package standupcli

import (
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/config"
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/report"
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/task"
	"github.com/spf13/cobra"
)

const StandUpRecommendedCommandName = "standup"

func rootStandupCommand(name, fullname string) *cobra.Command {
	var rootCmd *cobra.Command
	rootCmd = &cobra.Command{
		Use:   name,
		Short: "standup",
		Long:  "standup",
	}

	rootCmd.AddCommand(
		task.NewCmdTask(task.RecommendedCommandNameTask, common.GetFullName(fullname, task.RecommendedCommandNameTask)),
		report.NewCmdReport(report.RecommendedCommandNameReport, common.GetFullName(fullname, report.RecommendedCommandNameReport)),
		config.NewCmdConfig(config.RecommendedCommandNameConfig, common.GetFullName(fullname, config.RecommendedCommandNameConfig)),
	)

	return rootCmd
}

func NewCmdStandUp(name, fullname string) *cobra.Command {
	var rootcmd *cobra.Command
	rootcmd = rootStandupCommand(name, fullname)
	return rootcmd
}
