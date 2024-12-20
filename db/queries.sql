-- name: PokemonByID :one
SELECT * FROM pokemon WHERE id = $1;

/* 
with new_user as (
  insert into user_account(name, email)
  values ('arthur', 'some@where.com')
  returning user_id
)
insert into other_table (user_id, some_column)
select user_id, 'some value'
from new_user;
*/

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

-- with iv_stat as (
--     insert into pokemon_stats(hp, attack, defense, special_attack, special_defense, speed)
--     values (0, 0, 0, 0, 0, 0)
--     returning id
-- ), ev_stat as (
--     insert into pokemon_stats(hp, attack, defense, special_attack, special_defense, speed)
--     values (0, 0, 0, 0, 0, 0)
--     returning id
-- ), stats_stat as (
--     insert into pokemon_stats(hp, attack, defense, special_attack, special_defense, speed)
--     values (0, 0, 0, 0, 0, 0)
--     returning id
-- )
-- insert into pokemon(species_id, "level", iv_key, ev_key, stats_key)
-- values ($1, $2, (select id from iv_stat), (select id from ev_stat), (select id from stats_stat))
-- returning *;


-- INSERT INTO pokemon (
--     species_id, "level"
-- ) VALUES (
--     $1, $2
-- ) RETURNING *;

-- name: ListPokemon :many
SELECT * FROM pokemon;