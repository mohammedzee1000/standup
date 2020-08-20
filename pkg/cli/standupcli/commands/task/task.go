package task

import (
	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameTask = "task"

func NewCmdTask(name string, fullname string) *cobra.Command {
	add := NewTaskAddCommand(RecommendedCommandNameAdd, common.GetFullName(fullname, RecommendedCommandNameAdd))

	var taskCommand *cobra.Command
	taskCommand = &cobra.Command{
		Use:   name,
		Short: "Manipulate tasks",
		Long:  "Add, remove, update or delete tasks for specified date",
	}

	taskCommand.AddCommand(
		add,
	)
	return taskCommand
}
