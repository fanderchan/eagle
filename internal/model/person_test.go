package model

import (
	"fmt"
	"os"
	"testing"
)

func TestPerson_Save(t *testing.T) {
	p := Person{
		Name: "Ethan",
		Age:  10,
		Sex:  "woman",
	}

	err := p.Save()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
