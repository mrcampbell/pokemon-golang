-- name: PokemonByID :one
select 
p.*
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
where p.id = $1;

-- name: CreateStats :one
insert into stats(id, hp, attack, defense, special_attack, special_defense, speed)
values ($1, $2, $3, $4, $5, $6, $7)
returning id;

-- name: CreatePokemon :one
insert into pokemon(id, species_id, "level")
values ($1, $2, $3)
returning id;

-- name: CreatePokemonStats :one
insert into pokemon_stats(pokemon_id, stats_id, stat_type)
values ($1, $2, $3)
returning *;

-- name: ListPokemon :many
SELECT * FROM pokemon;