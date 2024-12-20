--
-- PostgreSQL database dump
--

-- Dumped from database version 14.15
-- Dumped by pg_dump version 14.15 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: pokemon; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pokemon (
    id uuid NOT NULL,
    species_id integer NOT NULL,
    level integer NOT NULL,
    move_one_id integer NOT NULL,
    move_two_id integer NOT NULL,
    move_three_id integer NOT NULL,
    move_four_id integer NOT NULL
);


ALTER TABLE public.pokemon OWNER TO postgres;

--
-- Name: pokemon_stats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pokemon_stats (
    pokemon_id uuid NOT NULL,
    stats_id uuid NOT NULL,
    stat_type character varying(4) NOT NULL
);


ALTER TABLE public.pokemon_stats OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: stats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stats (
    id uuid NOT NULL,
    hp integer NOT NULL,
    attack integer NOT NULL,
    defense integer NOT NULL,
    special_attack integer NOT NULL,
    special_defense integer NOT NULL,
    speed integer NOT NULL
);


ALTER TABLE public.stats OWNER TO postgres;

--
-- Name: pokemon pokemon_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pokemon
    ADD CONSTRAINT pokemon_pkey PRIMARY KEY (id);


--
-- Name: pokemon_stats pokemon_stats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pokemon_stats
    ADD CONSTRAINT pokemon_stats_pkey PRIMARY KEY (pokemon_id, stats_id, stat_type);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: stats stats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_pkey PRIMARY KEY (id);


--
-- Name: pokemon_stats pokemon_stats_pokemon_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pokemon_stats
    ADD CONSTRAINT pokemon_stats_pokemon_id_fkey FOREIGN KEY (pokemon_id) REFERENCES public.pokemon(id);


--
-- Name: pokemon_stats pokemon_stats_stats_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pokemon_stats
    ADD CONSTRAINT pokemon_stats_stats_id_fkey FOREIGN KEY (stats_id) REFERENCES public.stats(id);


--
-- PostgreSQL database dump complete
--

