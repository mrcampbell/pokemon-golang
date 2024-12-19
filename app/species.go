package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrcampbell/pokemon-golang/app/element"
	"github.com/mrcampbell/pokemon-golang/pokeapi"
)

type Species struct {
	Name     string
	Stats    Stats
	Elements [2]element.Element
}

func SpeciesFromFile(id int) Species {
	s, err := readfile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}
	result := Species{Name: s.Name}

	// stats
	result.Stats = Stats{
		HP:      s.GetStat("hp"),
		Attack:  s.GetStat("attack"),
		Defense: s.GetStat("defense"),
		SpAtk:   s.GetStat("special-attack"),
		SpDef:   s.GetStat("special-defense"),
		Speed:   s.GetStat("speed"),
	}

	// elements
	result.Elements[0] = s.GetElement(0)
	result.Elements[1] = s.GetElement(1)

	return result
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
