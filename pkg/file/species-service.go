package file

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pokeapi"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

type SpeciesService struct {
	gameVersion version.Version
}

func NewSpeciesService(gameVersion version.Version) app.SpeciesService {
	return SpeciesService{
		gameVersion: gameVersion,
	}
}

func (s SpeciesService) GetSpecies(id int) app.Species {
	return speciesFromFile(id, s.gameVersion)
}

func speciesFromFile(id int, version version.Version) app.Species {
	source, err := readfile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}
	result := app.Species{Name: source.Name}

	// stats
	result.Stats = app.Stats{
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
	result.MovesLearned = LearnableMovesFromSource(source, version)
	result.Abilities = AbilitiesFromSource(source)

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

func LearnableMovesFromSource(source pokeapi.Species, version version.Version) []app.LearnableMove {
	var moves []app.LearnableMove
	for _, move := range source.Moves {
		for _, versionGroupDetail := range move.VersionGroupDetails {
			if versionGroupDetail.VersionGroup.Name == string(version) {
				urlParts := strings.Split(move.Move.URL, "/")
				moveIDStr := urlParts[len(urlParts)-2]
				moveID := -1
				if _, err := fmt.Sscanf(moveIDStr, "%d", &moveID); err != nil {
					// todo: handle error
					panic(err)
				}
				moves = append(moves, app.LearnableMove{
					MoveID: moveID,
					Method: app.MoveLearnMethod(versionGroupDetail.MoveLearnMethod.Name),
					Level:  versionGroupDetail.LevelLearnedAt,
				})
			}
		}
	}
	return moves
}

func AbilitiesFromSource(source pokeapi.Species) []app.Ability {
	var abilities []app.Ability
	for _, ability := range source.Abilities {
		urlParts := strings.Split(ability.Ability.URL, "/")
		abilityIDStr := urlParts[len(urlParts)-2]
		abilityID := -1
		if _, err := fmt.Sscanf(abilityIDStr, "%d", &abilityID); err != nil {
			// todo: handle error
			panic(err)
		}
		abilities = append(abilities, app.Ability{
			ID:   abilityID,
			Name: ability.Ability.Name,
		})
	}
	return abilities
}
