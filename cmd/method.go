package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var methodCmd = &cobra.Command{
	Use:     "method",
	Aliases: []string{"m", "M"},
	Short:   "set request Method",
	Long:    "set a request method is support post、get",
	Example: "-M GET|POST",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre run method command")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run method command")
	},
}
