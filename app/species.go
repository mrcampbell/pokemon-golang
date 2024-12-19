package app

import (
	"fmt"
	"os"
)

type Species struct {
	Name string
}

func SpeciesFromFile(id int) Species {
	data, err := readfile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return Species{Name: ""}
}

func readfile(path string) (string, error) {
	// read file
	const path_dir = "data/pokemon/"
	data, err := os.ReadFile(path_dir + path + ".json")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
