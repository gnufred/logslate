package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logslate",
	Short: "Translate NGINX log files",
	Long: `Translate NGINX log files into human readable and usable format, for defensive investigation.

Basic usage:
	cat access.log | logslate
`,
	Args: func(cmd *cobra.Command, args []string) error {
		return rootCmdValidateArgs()
	},
	Run: func(cmd *cobra.Command, args []string) {
		rootCmdRun()
	},
}

func rootCmdValidateArgs() error {
	fmt.Println("Arguments are valid")
	return nil
}

func rootCmdRun() {
	fmt.Println("Command is running")
}

func Execute() error {
	return rootCmd.Execute()
}
