package update

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"eagle/internal/model"
	"eagle/util"
)

var Male = "Male"
var FeMale = "Female"

type UpdateOptions struct {
	Name string
	Age  int
	Sex  string
}

func NewUpdate() *cobra.Command {
	o := UpdateOptions{}
	cmd := &cobra.Command{
		GroupID: "CRS",
		Use:     "update",
		Short:   "update resource",
		Long:    `update resource by name ....`,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())
		},
		SuggestionsMinimumDistance: 1,
	}

	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "name")
	// 设置必填写
	cmd.MarkFlagRequired("name")
	cmd.Flags().IntVarP(&o.Age, "age", "a", 0, "age")
	cmd.Flags().StringVarP(&o.Sex, "sex", "s", "", "sex")
	// 组合标志
	cmd.MarkFlagsRequiredTogether("name", "age")

	return cmd
}

func (o *UpdateOptions) Validate() error {
	if !strings.EqualFold(o.Sex, Male) && !strings.EqualFold(o.Sex, FeMale) {
		return errors.New("sex must be Male or Female")

	}
	return nil
}

func (o *UpdateOptions) Run() error {
	person := model.NewPerson()
	person.Name = o.Name
	person.Sex = o.Sex
	person.Age = o.Age
	return person.Update()
}
