package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDelete() *cobra.Command {
	cmd := &cobra.Command{
		GroupID: "CRS",
		Use:     "delete",
		Short:   "delete resource",
		Long:    `delete resource by key`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("delete resouces, %v\n", args)
		},
	}

	return cmd
}
