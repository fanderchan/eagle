package create

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"eagle/internal/model"
	"eagle/util"
)

type CreateOptions struct {
	Data string
}

var createExample = `
# Create a Person

create -d "{'name':'jack','age':10,'sex':'man'}"

`

func NewCreate() *cobra.Command {
	o := CreateOptions{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create resource",
		Long: `crate resource from a file or form stdin.
					JSON and YAML formats are accepted`,
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Validate())
			util.CheckErr(o.Run())

		},
		GroupID: "CRS",
	}

	cmd.Flags().StringVarP(&o.Data, "data", "d", "", "createdata")

	return cmd
}

func (o *CreateOptions) Validate() error {
	if o.Data == "" || len(o.Data) == 0 {
		return errors.New("data can`t be null")
	}

	strings.ReplaceAll(o.Data, "'", "\"")
	return nil
}

func (o *CreateOptions) Run() error {
	return o.createPerson()
}

func (o *CreateOptions) createPerson() error {

	p := model.Person{}
	// d, err := strconv.Unquote(o.Data)
	// if err != nil {
	// 	return err
	// }
	err := json.Unmarshal([]byte(o.Data), &p)
	if err != nil {
		return err
	}

	if err = p.Save(); err != nil {
		return err
	}

	return nil
}
