-- name: PokemonByID :one
SELECT * FROM pokemon WHERE id = $1;

-- name: CreatePokemon :one
INSERT INTO pokemon (
    species_id, "level"
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListPokemon :many
SELECT * FROM pokemon;