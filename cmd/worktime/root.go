package worktime

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "worktime",
	Version: version,
	Short:   "worktime - a simple CLI to calculate remaining working hours",
	Long:    "worktime is ...long description",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops, There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
