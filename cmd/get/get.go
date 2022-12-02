package get

import (
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
		Short:   "get resource",
		Long:    `get resource from local storage`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())
		},
		SuggestionsMinimumDistance: 1,
		SuggestFor:                 []string{"find"},
	}

	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "get resource by name")
	cmd.Flags().IntVarP(&o.ID, "ID", "i", 0, "get resource by id")

	return cmd
}

func (o *GetOptions) Run() error {
	p := model.Person{}
	p.Name = o.Name

	person, err := p.Get()
	if err != nil {
		return err
	}
	fmt.Printf("Person {name: %s age: %d sex: %s} \n", person.Name, person.Age, person.Sex)
	return nil

}
