package report

import (
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameReport = "report"

func NewCmdReport(name, fullname string) *cobra.Command {
	get := NewReportGetCommand(RecommendedCommandNameGet, common.GetFullName(fullname, RecommendedCommandNameGet))

	var reportCommand *cobra.Command
	reportCommand = &cobra.Command{
		Use:   name,
		Short: "get reports",
		Long:  "get reports for standup",
	}

	reportCommand.AddCommand(
		get,
	)
	return reportCommand
}
