package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mrcampbell/pokemon-golang/pokeapi"
	"github.com/mrcampbell/pokemon-golang/pokeapi/language"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

const defaultLanguage = language.English

type MoveLearnMethod string

const (
	MoveLearnMethodLevelUp MoveLearnMethod = "level-up"
	MoveLearnMethodMachine MoveLearnMethod = "machine"
	MoveLearnMethodTutor   MoveLearnMethod = "tutor"
	MoveLearnMethodEgg     MoveLearnMethod = "egg"
)

type MoveDamageClass string

const (
	MoveDamageClassStatus   MoveDamageClass = "status"
	MoveDamageClassPhysical MoveDamageClass = "physical"
	MoveDamageClassSpecial  MoveDamageClass = "special"
)

type MoveCategory string

const (
	MoveCategoryDamage           MoveCategory = "damage"
	MoveCategoryAilment          MoveCategory = "ailment"
	MoveCategoryNetGoodStats     MoveCategory = "net-good-stats"
	MoveCategoryHeal             MoveCategory = "heal"
	MoveCategoryDamageAndAilment MoveCategory = "damage+ailment"
	MoveCategorySwagger          MoveCategory = "swagger"
	MoveCategoryDamageAndLower   MoveCategory = "damage+lower"
	MoveCategoryDamageAndRaise   MoveCategory = "damage+raise"
	MoveCategoryDamageAndHeal    MoveCategory = "damage+heal"
	MoveCategoryOhko             MoveCategory = "ohko"
	MoveCategoryWholeFieldEffect MoveCategory = "whole-field-effect"
	MoveCategoryFieldEffect      MoveCategory = "field-effect"
	MoveCategoryForceSwitch      MoveCategory = "force-switch"
	MoveCategoryUnique           MoveCategory = "unique"
)

type Stat string

const (
	StatHP         Stat = "hp"
	StatAttack     Stat = "attack"
	StatDefense    Stat = "defense"
	StatSpecialAtk Stat = "special-attack"
	StatSpecialDef Stat = "special-defense"
	StatSpeed      Stat = "speed"
	StatAccuracy   Stat = "accuracy"
	StatEvasion    Stat = "evasion"
)

type StatChange struct {
	Change int
	Stat   Stat
}

type Move struct {
	ID            int
	Name          string
	Accuracy      int
	DamageClass   MoveDamageClass
	EffectChance  int
	EffectText    string
	FlavorText    string
	Ailment       Ailment
	AilmentChance int
	Category      MoveCategory
	CriticalRate  int
	Drain         int
	FlinchChance  int
	Healing       int
	MaxHits       int
	MaxTurns      int
	MinHits       int
	MinTurns      int
	StatChance    int
	Power         int
	PP            int
	Priority      int
	StatChanges   []StatChange
}

type LearnableMove struct {
	MoveID int
	Level  int
	Method MoveLearnMethod
}

// todo: move to service, just storing here for simplicity and speed of development
func LoadMove(id int) Move {
	source, err := readFile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}

	result := Move{
		ID:            source.ID,
		Name:          source.Name,
		DamageClass:   MoveDamageClass(source.DamageClass.Name),
		EffectChance:  source.GetEffectChance(),
		EffectText:    source.GetEffectText(defaultLanguage),
		FlavorText:    source.GetFlavorText(defaultLanguage),
		Ailment:       Ailment(source.Meta.Ailment.Name),
		AilmentChance: source.Meta.AilmentChance,
		Category:      MoveCategory(source.Meta.Category.Name),
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

func LearnableMovesFromSource(source pokeapi.Species, version version.Version) []LearnableMove {
	var moves []LearnableMove
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
				moves = append(moves, LearnableMove{
					MoveID: moveID,
					Method: MoveLearnMethod(versionGroupDetail.MoveLearnMethod.Name),
					Level:  versionGroupDetail.LevelLearnedAt,
				})
			}
		}
	}
	return moves
}

func StatChangesFromSource(source pokeapi.Move) []StatChange {
	result := make([]StatChange, len(source.StatChanges))
	for i, change := range source.StatChanges {
		result[i] = StatChange{
			Change: change.Change,
			Stat:   Stat(change.Stat.Name),
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
