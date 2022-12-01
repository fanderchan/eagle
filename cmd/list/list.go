package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewList() *cobra.Command {
	cmd := &cobra.Command{
		GroupID: "Q",
		Use:     "list",
		Short:   "list resource",
		Long:    `list resource by page and query`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("list resouces, %v\n", args)
		},
	}

	return cmd
}
