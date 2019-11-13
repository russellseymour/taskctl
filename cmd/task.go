package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/trntv/wilson/pkg/runner"
)

func init() {
	runCommand.AddCommand(taskRunCommand)
}

var taskRunCommand = &cobra.Command{
	Use:   "task [task]",
	Short: "Schedule task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// todo: OnlyValidArgs
		var tname = args[0]
		t, ok := tasks[tname]
		if !ok {
			logrus.Fatalf("unknown task %s", tname)
		}

		rr := runner.NewRunner(nil, contexts, raw, quiet)
		rr.Run(t)

		fmt.Println(aurora.Yellow("\r\nSummary:"))
		printSummary(t)

		close(done)
	},
}