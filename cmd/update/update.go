package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUpdate() *cobra.Command {
	cmd := &cobra.Command{
		GroupID: "CRS",
		Use:     "update",
		Short:   "update resource",
		Long:    `update resource....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("update resouces, %v\n", args)
		},
	}

	return cmd
}
