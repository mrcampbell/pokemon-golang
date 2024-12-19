package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrcampbell/pokemon-golang/pkg/app/element"
	"github.com/mrcampbell/pokemon-golang/pokeapi"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

const defaultVersion = version.Emerald

// const defaultVersionGroup = versiongroup.Emerald
// const defaultGeneration = generation.Generation3

type Species struct {
	GameIndex    int
	Name         string
	Stats        Stats
	Elements     [2]element.Element
	MovesLearned []LearnableMove
	Abilities    []Ability
}

func Load(id int) Species {
	return SpeciesFromFile(id, defaultVersion)
}

func SpeciesFromFile(id int, version version.Version) Species {
	source, err := readfile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}
	result := Species{Name: source.Name}

	// stats
	result.Stats = Stats{
		HP:      source.GetStat("hp"),
		Attack:  source.GetStat("attack"),
		Defense: source.GetStat("defense"),
		SpAtk:   source.GetStat("special-attack"),
		SpDef:   source.GetStat("special-defense"),
		Speed:   source.GetStat("speed"),
	}

	// elements
	result.Elements[0] = source.GetElement(0)
	result.Elements[1] = source.GetElement(1)

	// version specific data
	result.GameIndex = source.GetGameIndex(version)

	learnableMoves := source.GetMovesLearned(version)
	result.MovesLearned = LearnableMovesFromSource(learnableMoves)

	sourceAbilities := source.GetAbilities()
	result.Abilities = AbilitiesFromSource(sourceAbilities)

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
