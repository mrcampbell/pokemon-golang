CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.stats (
    id UUID PRIMARY KEY,
    hp INTEGER NOT NULL,
    attack INTEGER NOT NULL,
    defense INTEGER NOT NULL,
    special_attack INTEGER NOT NULL,
    special_defense INTEGER NOT NULL,
    speed INTEGER NOT NULL
);

CREATE TABLE public.pokemon (
    id UUID PRIMARY KEY,
    species_id INTEGER NOT NULL,
    "level" INTEGER NOT NULL,
    move_one_id INTEGER NOT NULL,
    move_two_id INTEGER NOT NULL,
    move_three_id INTEGER NOT NULL,
    move_four_id INTEGER NOT NULL
);

CREATE TABLE public.pokemon_stats (
    pokemon_id UUID NOT NULL,
    stats_id UUID NOT NULL,
    stat_type VARCHAR(4) NOT NULL, -- iv, ev, live
    PRIMARY KEY (pokemon_id, stats_id, stat_type),
    FOREIGN KEY (pokemon_id) REFERENCES public.pokemon(id),
    FOREIGN KEY (stats_id) REFERENCES public.stats(id)
);

-- ALTER TABLE public.pokemon_stats 
-- ADD CONSTRAINT pokemon_stats_pokemon_id_fkey 
-- FOREIGN KEY (pokemon_id) 
-- REFERENCES public.pokemon(id);
