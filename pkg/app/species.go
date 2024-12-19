package app

import (
	"github.com/mrcampbell/pokemon-golang/pkg/app/element"
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

type SpeciesService interface {
	GetSpecies(id int) Species
}
