package game

import (
	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pkg/rand"
)

type PokemonService struct {
	moveService    app.MoveService
	speciesService app.SpeciesService
}

func NewPokemonService(moveService app.MoveService, speciesService app.SpeciesService) app.PokemonService {
	return &PokemonService{
		moveService:    moveService,
		speciesService: speciesService,
	}
}

func (s *PokemonService) CreatePokemon(speciesID int, level int) app.Pokemon {
	species := s.speciesService.GetSpecies(speciesID)

	result := app.Pokemon{
		SpeciesID: speciesID,
		Name:      species.Name,
		Level:     level,
		Moves:     []app.Move{},
		BaseStats: species.Stats,
		IVs:       randomStats(),
		EVs:       app.Stats{},
	}

	result.Stats = CalculateStats(result.BaseStats, result.IVs, result.EVs, level)
	return result
}

func randomStats() app.Stats {
	return app.Stats{
		HP:      rand.RandBetween(0, 31),
		Attack:  rand.RandBetween(0, 31),
		Defense: rand.RandBetween(0, 31),
		SpAtk:   rand.RandBetween(0, 31),
		SpDef:   rand.RandBetween(0, 31),
		Speed:   rand.RandBetween(0, 31),
	}
}

func CalculateStats(base app.Stats, ivs app.Stats, evs app.Stats, level int) app.Stats {
	return app.Stats{
		HP:      calcHP(base.HP, ivs.HP, evs.HP, level),
		Attack:  calcStat(base.Attack, ivs.Attack, evs.Attack, level),
		Defense: calcStat(base.Defense, ivs.Defense, evs.Defense, level),
		SpAtk:   calcStat(base.SpAtk, ivs.SpAtk, evs.SpAtk, level),
		SpDef:   calcStat(base.SpDef, ivs.SpDef, evs.SpDef, level),
		Speed:   calcStat(base.Speed, ivs.Speed, evs.Speed, level),
	}
}

func calcHP(base int, iv int, ev int, level int) int {
	return ((2*base + iv + ev/4) * level / 100) + level + 10
}

func calcStat(base int, iv int, ev int, level int) int {
	return ((2*base + iv + ev/4) * level / 100) + 5
}
