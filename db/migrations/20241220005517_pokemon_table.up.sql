CREATE TABLE public.pokemon (
    id SERIAL PRIMARY KEY,
    species_id INTEGER NOT NULL,
    "level" INTEGER NOT NULL
);