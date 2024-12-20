package app

import "fmt"

type Pokemon struct {
	SpeciesID int    `json:"species_id"`
	Name      string `json:"name"`
	Level     int    `json:"level"`
	Moves     []Move `json:"moves"`
	BaseStats Stats  `json:"base_stats"`
	IVs       Stats  `json:"ivs"`
	EVs       Stats  `json:"evs"`
	Stats     Stats  `json:"stats"` // calculated stats, from ivs, evs, base stats, and level
}

func (p Pokemon) PrintableSummary() string {
	result := fmt.Sprintf("Name: %s\n", p.Name)
	result += fmt.Sprintf("Level: %d\n", p.Level)
	result += fmt.Sprintf("Stats: %v\n", p.Stats)
	result += "Moves: \n"
	for _, move := range p.Moves {
		result += move.PrintableSummary()
	}
	return result
}

type PokemonService interface {
	CreatePokemon(speciesID int, level int) Pokemon
}
