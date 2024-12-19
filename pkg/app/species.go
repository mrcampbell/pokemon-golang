package app

import (
	"github.com/mrcampbell/pokemon-golang/pkg/app/element"
)

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
