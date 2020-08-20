package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mohammedzee1000/standup/pkg/cli/standupcli/commands/standupcli"
)

func main() {
	root := standupcli.NewCmdStandUp(standupcli.StandUpRecommendedCommandName, standupcli.StandUpRecommendedCommandName)
	// parse the flags but hack around to avoid exiting with error code 2 on help
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError)
	args := os.Args[1:]
	flag.Usage = func() {
		_ = root.Help()
	}
	if err := flag.CommandLine.Parse(args); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Printf("Failed to parse commandline %w", err)
			os.Exit(1)
		}
	}
	er := root.Execute()
	if er != nil {
		fmt.Printf("Failed to execute root cmd %w", er)
	}
}
