package app

import "github.com/mrcampbell/pokemon-golang/pokeapi"

type Ability struct {
	ID   int
	Name string
}

func AbilitiesFromSource(source []pokeapi.Ability) []Ability {
	result := make([]Ability, len(source))
	for i, ability := range source {
		result[i] = Ability{
			ID:   ability.ID,
			Name: ability.Name,
		}
	}
	return result
}
