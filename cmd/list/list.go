package list

import (
	"fmt"

	"github.com/spf13/cobra"

	"eagle/internal/model"
	"eagle/util"
)

type ListOptions struct {
}

func NewList() *cobra.Command {
	o := ListOptions{}
	cmd := &cobra.Command{
		GroupID: "Q",
		Use:     "list",
		Short:   "list resource",
		Long:    `list resource from local storage`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())
		},
	}

	return cmd
}

func (o *ListOptions) Run() error {
	p := model.NewPerson()
	persons, err := p.List()
	if err != nil {
		return err
	}

	for _, v := range persons {
		fmt.Printf("Person {name: %s age: %d sex: %s} \n", v.Name, v.Age, v.Sex)

	}
	return nil
}
