package delete

import (
	"github.com/spf13/cobra"

	"eagle/internal/model"
	"eagle/util"
)

type DeleteOptions struct {
	Name string
}

func NewDelete() *cobra.Command {
	o := DeleteOptions{}

	cmd := &cobra.Command{
		GroupID: "CRS",
		Use:     "delete",
		Short:   "delete resource",
		Long:    `delete resource by name`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())

		},
		SuggestionsMinimumDistance: 1,
		SuggestFor:                 []string{"remove", "truncate"},
	}

	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "delete by name")
	cmd.MarkFlagRequired("name")

	return cmd
}

func (o *DeleteOptions) Run() error {
	p := model.NewPerson()
	p.Name = o.Name
	return p.Delete()
}
