package get

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"eagle/internal/model"
	"eagle/util"
)

type GetOptions struct {
	ID   int
	Name string
}

func NewGet() *cobra.Command {

	o := GetOptions{}

	cmd := &cobra.Command{
		GroupID: "Q",
		Use:     "get",
		Short:   "get resources",
		Long:    `get  resources from map`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())
		},
	}

	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "get resource by name")
	cmd.Flags().IntVarP(&o.ID, "ID", "i", 0, "get resource by id")

	return cmd
}

func (o *GetOptions) Run() error {
	p := model.Person{}
	p.Name = o.Name

	v, ok := p.Get()
	if !ok {
		return errors.New("record not found")
	}

	fmt.Printf("%v \n", v)

	return nil

}
