package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:     "data",
	Short:   "set request data",
	Long:    "this command can set request data",
	Aliases: []string{"d", "D"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("data command PersistentPreRun")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("data command preRun")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("data command run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("data command PersistentPostRun")
	},
}
