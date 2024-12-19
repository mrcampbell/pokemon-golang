package file

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pokeapi"
	"github.com/mrcampbell/pokemon-golang/pokeapi/language"
)

type MoveService struct {
	language language.Language
}

func NewMoveService(language language.Language) app.MoveService {
	return MoveService{
		language: language,
	}
}

func (m MoveService) GetMove(id int) app.Move {
	return loadMove(id, m.language)
}

func loadMove(id int, language language.Language) app.Move {
	source, err := readFile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}

	result := app.Move{
		ID:            source.ID,
		Name:          source.Name,
		DamageClass:   app.MoveDamageClass(source.DamageClass.Name),
		EffectChance:  source.GetEffectChance(),
		EffectText:    source.GetEffectText(language),
		FlavorText:    source.GetFlavorText(language),
		Ailment:       app.Ailment(source.Meta.Ailment.Name),
		AilmentChance: source.Meta.AilmentChance,
		Category:      app.MoveCategory(source.Meta.Category.Name),
		CriticalRate:  source.Meta.CritRate,
		Drain:         source.Meta.Drain,
		FlinchChance:  source.Meta.FlinchChance,
		Healing:       source.Meta.Healing,
		MaxHits:       nilableToInt(source.Meta.MaxHits),
		MaxTurns:      nilableToInt(source.Meta.MaxTurns),
		MinHits:       nilableToInt(source.Meta.MinHits),
		MinTurns:      nilableToInt(source.Meta.MinTurns),
		StatChance:    source.Meta.StatChance,
		Power:         source.Power,
		PP:            source.Pp,
		Priority:      source.Priority,
		StatChanges:   StatChangesFromSource(source),
	}

	return result
}

func StatChangesFromSource(source pokeapi.Move) []app.StatChange {
	result := make([]app.StatChange, len(source.StatChanges))
	for i, change := range source.StatChanges {
		result[i] = app.StatChange{
			Change: change.Change,
			Stat:   app.Stat(change.Stat.Name),
		}
	}
	return result
}

func readFile(path string) (pokeapi.Move, error) {
	const path_dir = "data/moves/"
	data, err := os.ReadFile(path_dir + path + ".json")
	if err != nil {
		return pokeapi.Move{}, err
	}

	pa_move := pokeapi.Move{}
	err = json.Unmarshal(data, &pa_move)
	if err != nil {
		return pokeapi.Move{}, err
	}

	return pa_move, nil
}

func nilableToInt(source interface{}) int {
	if source == nil {
		return 0
	}
	value := source.(float64)
	return int(value)
}
