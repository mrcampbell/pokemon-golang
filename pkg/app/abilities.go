package app

import (
	"fmt"
	"strings"

	"github.com/mrcampbell/pokemon-golang/pokeapi"
)

type Ability struct {
	ID   int
	Name string
}

func AbilitiesFromSource(source pokeapi.Species) []Ability {
	var abilities []Ability
	for _, ability := range source.Abilities {
		urlParts := strings.Split(ability.Ability.URL, "/")
		abilityIDStr := urlParts[len(urlParts)-2]
		abilityID := -1
		if _, err := fmt.Sscanf(abilityIDStr, "%d", &abilityID); err != nil {
			// todo: handle error
			panic(err)
		}
		abilities = append(abilities, Ability{
			ID:   abilityID,
			Name: ability.Ability.Name,
		})
	}
	return abilities
}
