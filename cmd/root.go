package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"eagle/cmd/create"
	"eagle/cmd/delete"
	"eagle/cmd/get"
	"eagle/cmd/list"
	"eagle/cmd/update"
)

var rootCmd = &cobra.Command{
	Use:     "eagle",
	Short:   "eagle is an example",                            // 简短介绍
	Long:    "eagle is an example to show how to use cobra  ", // 完整介绍
	Version: "0.0.1",                                          // 设置版本号，如果添加了可以
	Run:     runHelp,
}

var rootUsageTemplate = `
Use eagle you can store the person info to local storage

# Example
	eagle create -n Jack -a 10 -n
	eagle get -n Jack
	eagle list
	...
`

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	basicCommandQ := cobra.Group{
		Title: "Basic Command(Q)",
		ID:    "Q",
	}
	rootCmd.AddGroup(&basicCommandQ)
	rootCmd.AddCommand(get.NewGet(), list.NewList())

	basicCommandCRS := cobra.Group{
		Title: "Basic Command(CRS)",
		ID:    "CRS",
	}

	rootCmd.AddGroup(&basicCommandCRS)
	rootCmd.AddCommand(create.NewCreate(), delete.NewDelete(), update.NewUpdate())

	// 设置使用介绍模版
	// rootCmd.SetUsageTemplate(rootUsageTemplate)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
