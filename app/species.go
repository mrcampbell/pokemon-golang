package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrcampbell/pokemon-golang/pokeapi"
)

type Species struct {
	Name string
}

func SpeciesFromFile(id int) Species {
	s, err := readfile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}
	return Species{Name: s.Name}
}

func readfile(path string) (pokeapi.Species, error) {
	// read file
	const path_dir = "data/pokemon/"
	data, err := os.ReadFile(path_dir + path + ".json")
	if err != nil {
		return pokeapi.Species{}, err
	}

	pa_species := pokeapi.Species{}
	err = json.Unmarshal(data, &pa_species)
	if err != nil {
		return pokeapi.Species{}, err
	}

	return pa_species, nil
}
