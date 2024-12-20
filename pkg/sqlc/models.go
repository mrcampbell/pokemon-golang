// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/google/uuid"
)

type Pokemon struct {
	ID        uuid.UUID
	SpeciesID int32
	Level     int32
	IvKey     uuid.UUID
	EvKey     uuid.UUID
	StatsKey  uuid.UUID
}

type PokemonStat struct {
	ID             uuid.UUID
	Hp             int32
	Attack         int32
	Defense        int32
	SpecialAttack  int32
	SpecialDefense int32
	Speed          int32
}

type SchemaMigration struct {
	Version int64
	Dirty   bool
}
