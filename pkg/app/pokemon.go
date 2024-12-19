package app

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

type PokemonService interface {
	CreatePokemon(speciesID int, level int) Pokemon
}
