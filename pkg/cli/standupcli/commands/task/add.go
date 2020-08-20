package task

import (
	"fmt"
	"time"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/common"
	"github.com/mohammedzee1000/standup/pkg/standup/standup"
	"github.com/mohammedzee1000/standup/pkg/standup/task"
	"github.com/spf13/cobra"
)

const RecommendedCommandNameAdd = "add"

type TaskAddOptions struct {
	*common.CommonOptions
	dt          time.Time
	section     string
	description string
}

func newTaskAddOptions() *TaskAddOptions {
	return &TaskAddOptions{
		CommonOptions: common.NewCommonOptions(),
	}
}

func (tao *TaskAddOptions) Complete(name string, cmd *cobra.Command, args []string) error {
	err := tao.InitContext()
	if err != nil {
		return err
	}
	tao.dt = time.Now()
	return nil
}

func (tao *TaskAddOptions) Validate() error {
	if tao.section == "" {
		return fmt.Errorf("section cannot be empty")
	}
	if tao.description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	return nil
}

func (tao *TaskAddOptions) Run() error {
	stc := standup.NewStandUpConfig(tao.dt)
	err := stc.FromFile(tao.Context)
	if err != nil {
		return err
	}
	t := task.New()
	t.Description = tao.description
	std := stc.GetStandUp()
	std.Sections[tao.section] = append(std.Sections[tao.section], t)
	err = stc.ToFile(tao.Context)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added task with id %s to standup\n", t.ID)
	return nil
}

func NewTaskAddCommand(name string, fullname string) *cobra.Command {
	o := newTaskAddOptions()

	var taskAddCommand *cobra.Command
	taskAddCommand = &cobra.Command{
		Use:   name,
		Short: "New task",
		Long:  "Add new task to specified days standup",
		Run: func(cmd *cobra.Command, args []string) {
			common.GenericRun(o, cmd, args)
		},
	}
	taskAddCommand.Flags().StringVarP(&o.section, "section", "s", "Tasks", "section name")
	taskAddCommand.Flags().StringVarP(&o.description, "description", "d", "", "description of task")
	return taskAddCommand
}
