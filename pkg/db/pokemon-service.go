package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pkg/rand"
	"github.com/mrcampbell/pokemon-golang/pkg/sqlc"
)

type PokemonService struct {
	db             *pgx.Conn
	queries        *sqlc.Queries
	moveService    app.MoveService
	speciesService app.SpeciesService
}

func NewPokemonService(db *pgx.Conn, queries *sqlc.Queries, moveService app.MoveService, speciesService app.SpeciesService) app.PokemonService {
	return &PokemonService{
		db:             db,
		queries:        queries,
		moveService:    moveService,
		speciesService: speciesService,
	}
}

func (s PokemonService) SavePokemon(ctx context.Context, pokemon app.Pokemon) (uuid.UUID, error) {
	// TODO: Isolate each step into an individual function
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := s.queries.WithTx(tx)

	ivID, err := saveStats(ctx, qtx, pokemon.IVs)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error saving IVs: %w", err)
	}

	evID, err := saveStats(ctx, qtx, pokemon.EVs)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error saving EVs: %w", err)
	}

	pID, err := qtx.CreatePokemon(ctx, sqlc.CreatePokemonParams{
		ID:        uuid.New(),
		SpeciesID: int32(pokemon.SpeciesID),
		Level:     int32(pokemon.Level),
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("error saving pokemon: %w", err)
	}

	_, err = qtx.CreatePokemonStats(ctx, sqlc.CreatePokemonStatsParams{
		PokemonID: pID,
		StatsID:   ivID,
		StatType:  "iv",
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("error saving pokemon IV stats: %w", err)
	}

	_, err = qtx.CreatePokemonStats(ctx, sqlc.CreatePokemonStatsParams{
		PokemonID: pID,
		StatsID:   evID,
		StatType:  "ev",
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("error saving pokemon EV stats: %w", err)
	}

	tx.Commit(ctx)
	return pID, nil
}

func (s PokemonService) GetPokemon(ctx context.Context, id uuid.UUID) (app.Pokemon, error) {
	return app.Pokemon{}, nil
}

func (s PokemonService) ListPokemon(ctx context.Context) ([]app.Pokemon, error) {
	return []app.Pokemon{}, nil
}

func (s PokemonService) CreatePokemon(speciesID int, level int) app.Pokemon {
	species := s.speciesService.GetSpecies(speciesID)
	learnedMoves := getRandomLevelUpMoveSet(species, level)
	moveSet := []app.Move{}
	for _, learnedMove := range learnedMoves {
		move := s.moveService.GetMove(learnedMove.MoveID)
		moveSet = append(moveSet, move)
	}

	result := app.Pokemon{
		SpeciesID: speciesID,
		Name:      species.Name,
		Level:     level,
		Moves:     moveSet,
		BaseStats: species.Stats,
		IVs:       randomStats(),
		EVs:       app.Stats{},
	}

	result.Stats = CalculateStats(result.BaseStats, result.IVs, result.EVs, level)
	return result
}

func saveStats(ctx context.Context, queries *sqlc.Queries, stats app.Stats) (uuid.UUID, error) {
	ivID, err := queries.CreateStats(ctx, sqlc.CreateStatsParams{
		ID:             uuid.New(),
		Hp:             int32(stats.HP),
		Attack:         int32(stats.Attack),
		Defense:        int32(stats.Defense),
		SpecialAttack:  int32(stats.SpAtk),
		SpecialDefense: int32(stats.SpDef),
		Speed:          int32(stats.Speed),
	})
	if err != nil {
		return uuid.Nil, err
	}
	return ivID, nil
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

func getRandomLevelUpMoveSet(species app.Species, level int) []app.LearnableMove {
	rand.ShuffleLearnableMoves(species.MovesLearned)
	moves := []app.LearnableMove{}
	for _, move := range species.MovesLearned {
		if move.Level <= level && move.Method == app.MoveLearnMethodLevelUp {
			moves = append(moves, move)
		}
	}
	movesLength := len(moves)
	movesLengthOr4 := min(movesLength, 4)
	return moves[0:movesLengthOr4]
}

func calcHP(base int, iv int, ev int, level int) int {
	return ((2*base + iv + ev/4) * level / 100) + level + 10
}

func calcStat(base int, iv int, ev int, level int) int {
	return ((2*base + iv + ev/4) * level / 100) + 5
}
