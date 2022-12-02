package create

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"eagle/internal/model"
	"eagle/util"
)

type CreateOptions struct {
	Name string
	Age  int
	Sex  bool
}

var Gender = map[bool]string{
	false: "Male",
	true:  "Female",
}

var createExample = `
# Create a Person

create -n Jack -a 10 -s 

`

func NewCreate() *cobra.Command {
	o := CreateOptions{}

	cmd := &cobra.Command{
		Use:                   "create",
		Short:                 "create resource",
		Long:                  `crate resource to local storage`,
		Example:               createExample,
		DisableFlagsInUseLine: true,
		// Args:    cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Check())
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Validate())
			util.CheckErr(o.Run())

		},
		GroupID:                    "CRS",
		SuggestionsMinimumDistance: 1,                         // 开启建议
		SuggestFor:                 []string{"save", "store"}, // 开启建议命令

	}

	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "name")
	// 设置必填写
	cmd.MarkFlagRequired("name")
	cmd.Flags().IntVarP(&o.Age, "age", "a", 0, "age")
	cmd.Flags().BoolVarP(&o.Sex, "sex", "s", false, "sex")
	// 组合标志
	cmd.MarkFlagsRequiredTogether("name", "age")

	return cmd
}

func (o *CreateOptions) Check() error {
	p := model.Person{
		Name: o.Name,
	}

	v, err := p.Get()
	if err != nil {
		return err
	}
	if v.Name != "" {
		return errors.New("duplicate name")
	}
	return nil
}

func (o *CreateOptions) Validate() error {
	if o.Age < 10 {
		return errors.New("age cannot be negative")
	}
	return nil
}

func (o *CreateOptions) Run() error {
	return o.createPerson()
}

func (o *CreateOptions) createPerson() error {

	p := model.Person{
		Name: o.Name,
		Age:  o.Age,
		Sex:  Gender[o.Sex],
	}

	if err := p.Save(); err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}
