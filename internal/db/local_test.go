package db

import (
	"fmt"
	"os"
	"testing"
)

func TestGetDb(t *testing.T) {
	d, err := GetDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	fmt.Println(d.filePath)
}

func TestCache_Save(t *testing.T) {
	d, err := GetDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	m := map[string]string{
		"Name": "Jack",
	}
	err = d.Save("Jack", m)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
}

func TestCache_Get(t *testing.T) {
	d, err := GetDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	value, err := d.Get("Jack")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	fmt.Printf("%v\n", value)
}

func TestCache_Delete(t *testing.T) {
	d, err := GetDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	err = d.Delete("Rose")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

}

func TestCache_List(t *testing.T) {
	d, err := GetDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	list, err := d.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	v, ok := list.(map[string]interface{})
	if !ok {
		return
	}
	for k, v := range v {
		fmt.Printf("name: %s -> %v \n", k, v)
	}
}
