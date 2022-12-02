package model

import (
	"errors"

	"eagle/internal/db"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func NewPerson() *Person {
	return &Person{}
}

func (p *Person) Save() error {
	d, err := db.GetDb()
	if err != nil {
		return err
	}
	return d.Save(p.Name, p)
}

func (p *Person) Get() (Person, error) {
	person := Person{}
	d, err := db.GetDb()
	if err != nil {
		return person, err
	}

	m, err := d.Get(p.Name)
	if err != nil {
		return person, err
	}
	result, ok := m.(map[string]interface{})
	if !ok {
		return person, errors.New("transform failed")
	}
	person.Name = result["name"].(string)
	person.Age = int(result["age"].(float64))
	person.Sex = result["sex"].(string)

	return person, nil
}

func (p *Person) Delete() error {
	d, err := db.GetDb()
	if err != nil {
		return err
	}
	return d.Delete(p.Name)

}

func (p *Person) List() ([]Person, error) {
	d, err := db.GetDb()
	if err != nil {
		return nil, err
	}

	list, err := d.List()
	if err != nil {
		return nil, err
	}

	m, ok := list.(map[string]interface{})
	if !ok {
		return nil, errors.New("transform failed")
	}

	persons := make([]Person, 0)
	for _, v := range m {
		m2 := v.(map[string]interface{})
		name := m2["name"].(string)
		age := m2["age"].(float64)
		sex := m2["sex"].(string)
		person := Person{
			Name: name,
			Age:  int(age),
			Sex:  sex,
		}

		persons = append(persons, person)
	}
	return persons, nil

}

func (p *Person) Update() error {
	d, err := db.GetDb()
	if err != nil {
		return err
	}
	return d.Update(p.Name, p)
}
