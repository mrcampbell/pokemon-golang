// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createPokemon = `-- name: CreatePokemon :one
insert into pokemon(id, species_id, "level")
values ($1, $2, $3)
returning id
`

type CreatePokemonParams struct {
	ID        uuid.UUID
	SpeciesID int32
	Level     int32
}

func (q *Queries) CreatePokemon(ctx context.Context, arg CreatePokemonParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createPokemon, arg.ID, arg.SpeciesID, arg.Level)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createPokemonStats = `-- name: CreatePokemonStats :one
insert into pokemon_stats(pokemon_id, stats_id, stat_type)
values ($1, $2, $3)
returning pokemon_id, stats_id, stat_type
`

type CreatePokemonStatsParams struct {
	PokemonID uuid.UUID
	StatsID   uuid.UUID
	StatType  string
}

func (q *Queries) CreatePokemonStats(ctx context.Context, arg CreatePokemonStatsParams) (PokemonStat, error) {
	row := q.db.QueryRow(ctx, createPokemonStats, arg.PokemonID, arg.StatsID, arg.StatType)
	var i PokemonStat
	err := row.Scan(&i.PokemonID, &i.StatsID, &i.StatType)
	return i, err
}

const createStats = `-- name: CreateStats :one
insert into stats(id, hp, attack, defense, special_attack, special_defense, speed)
values ($1, $2, $3, $4, $5, $6, $7)
returning id
`

type CreateStatsParams struct {
	ID             uuid.UUID
	Hp             int32
	Attack         int32
	Defense        int32
	SpecialAttack  int32
	SpecialDefense int32
	Speed          int32
}

func (q *Queries) CreateStats(ctx context.Context, arg CreateStatsParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createStats,
		arg.ID,
		arg.Hp,
		arg.Attack,
		arg.Defense,
		arg.SpecialAttack,
		arg.SpecialDefense,
		arg.Speed,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const listPokemon = `-- name: ListPokemon :many
SELECT id, species_id, level FROM pokemon
`

func (q *Queries) ListPokemon(ctx context.Context) ([]Pokemon, error) {
	rows, err := q.db.Query(ctx, listPokemon)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pokemon
	for rows.Next() {
		var i Pokemon
		if err := rows.Scan(&i.ID, &i.SpeciesID, &i.Level); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const pokemonByID = `-- name: PokemonByID :one
select 
p.id, p.species_id, p.level
, si.hp i_hp
, si.attack i_attack
, si.defense i_defense
, si.special_attack i_spec_atk
, si.special_defense i_spec_def
, si.speed i_speed
, se.hp e_hp
, se.attack e_attack
, se.defense e_defense
, se.special_attack e_spec_atk
, se.special_defense e_spec_def
, se.speed e_speed
from pokemon p 
left join pokemon_stats psi on psi.pokemon_id = p.id and psi.stat_type = 'iv'
left join pokemon_stats pse on pse.pokemon_id = p.id and pse.stat_type = 'ev'
left join stats si on si.id = psi.stats_id
left join stats se on se.id = pse.stats_id
where p.id = $1
`

type PokemonByIDRow struct {
	ID        uuid.UUID
	SpeciesID int32
	Level     int32
	IHp       pgtype.Int4
	IAttack   pgtype.Int4
	IDefense  pgtype.Int4
	ISpecAtk  pgtype.Int4
	ISpecDef  pgtype.Int4
	ISpeed    pgtype.Int4
	EHp       pgtype.Int4
	EAttack   pgtype.Int4
	EDefense  pgtype.Int4
	ESpecAtk  pgtype.Int4
	ESpecDef  pgtype.Int4
	ESpeed    pgtype.Int4
}

func (q *Queries) PokemonByID(ctx context.Context, id uuid.UUID) (PokemonByIDRow, error) {
	row := q.db.QueryRow(ctx, pokemonByID, id)
	var i PokemonByIDRow
	err := row.Scan(
		&i.ID,
		&i.SpeciesID,
		&i.Level,
		&i.IHp,
		&i.IAttack,
		&i.IDefense,
		&i.ISpecAtk,
		&i.ISpecDef,
		&i.ISpeed,
		&i.EHp,
		&i.EAttack,
		&i.EDefense,
		&i.ESpecAtk,
		&i.ESpecDef,
		&i.ESpeed,
	)
	return i, err
}
