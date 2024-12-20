CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.pokemon_stats (
    id UUID PRIMARY KEY,
    hp INTEGER NOT NULL,
    attack INTEGER NOT NULL,
    defense INTEGER NOT NULL,
    special_attack INTEGER NOT NULL,
    special_defense INTEGER NOT NULL,
    speed INTEGER NOT NULL
);

-- todo, joining table

CREATE TABLE public.pokemon (
    id UUID PRIMARY KEY,
    species_id INTEGER NOT NULL,
    "level" INTEGER NOT NULL,
    iv_key UUID NOT NULL,
    ev_key UUID NOT NULL,
    stats_key UUID NOT NULL,
    FOREIGN KEY (iv_key) REFERENCES public.pokemon_stats(id) ON DELETE CASCADE,
    FOREIGN KEY (ev_key) REFERENCES public.pokemon_stats(id) ON DELETE CASCADE,
    FOREIGN KEY (stats_key) REFERENCES public.pokemon_stats(id) ON DELETE CASCADE
);

-- ALTER TABLE public.pokemon_stats 
-- ADD CONSTRAINT pokemon_stats_pokemon_id_fkey 
-- FOREIGN KEY (pokemon_id) 
-- REFERENCES public.pokemon(id);
