package config

import (
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameConfig = "config"

func NewCmdConfig(name string, fullname string) *cobra.Command {
	var configCmd *cobra.Command
	view := NewCmdConfigView(RecommendedCommandNameGet, common.GetFullName(fullname, RecommendedCommandNameGet))
	set := NewCmdConfigSet(RecommendedCommandNameSet, common.GetFullName(fullname, RecommendedCommandNameSet))

	configCmd = &cobra.Command{
		Use:   name,
		Short: "Manipulate config",
		Long:  "Get config or update its parameters",
	}
	configCmd.AddCommand(
		view,
		set,
	)
	return configCmd
}
