package model

import (
	"eagle/internal/db"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Sex  string `json:"sex"`
}

func (p *Person) Save() error {
	d, err := db.GetDb()
	if err != nil {
		return err
	}
	return d.Save(p.Name, p)
}

func (p *Person) Get() (interface{}, bool) {
	return nil, false
}
