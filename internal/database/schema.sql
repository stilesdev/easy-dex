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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: pokemon; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.pokemon (
    id character varying NOT NULL,
    nid character varying NOT NULL,
    dex_num integer NOT NULL,
    form_id character varying,
    name character varying NOT NULL,
    form_name character varying,
    region character varying NOT NULL,
    generation integer NOT NULL,
    type1 character varying NOT NULL,
    type2 character varying,
    tera_type character varying,
    color character varying NOT NULL,
    primary_ability character varying NOT NULL,
    secondary_ability character varying,
    hidden_ability character varying,
    is_default boolean NOT NULL,
    is_form boolean NOT NULL,
    is_legendary boolean NOT NULL,
    is_mythical boolean NOT NULL,
    is_ultra_beast boolean NOT NULL,
    ultra_beast_code character varying,
    is_paradox boolean DEFAULT false NOT NULL,
    paradox_species character varying[],
    is_convergent boolean DEFAULT false NOT NULL,
    convergent_species character varying[],
    is_special_ability_form boolean NOT NULL,
    is_cosmetic_form boolean NOT NULL,
    is_female_form boolean NOT NULL,
    has_gender_differences boolean NOT NULL,
    is_battle_only_form boolean NOT NULL,
    is_switchable_form boolean NOT NULL,
    is_fusion boolean NOT NULL,
    fused_with character varying[],
    is_mega boolean NOT NULL,
    is_primal boolean NOT NULL,
    is_gmax boolean NOT NULL,
    is_regional boolean NOT NULL,
    can_gmax boolean NOT NULL,
    can_dynamax boolean NOT NULL,
    can_be_alpha boolean NOT NULL,
    debut_in character varying NOT NULL,
    obtainable_in character varying[] NOT NULL,
    version_exclusive_in character varying[] NOT NULL,
    event_only_in character varying[] NOT NULL,
    storable_in character varying[] NOT NULL,
    registrable_in character varying[] NOT NULL,
    shiny_released boolean NOT NULL,
    shiny_base character varying,
    base_stats json NOT NULL,
    weight integer NOT NULL,
    height integer NOT NULL,
    male_rate real NOT NULL,
    female_rate real NOT NULL,
    base_species character varying,
    base_forms character varying[],
    forms character varying[],
    family character varying,
    evolves_from json,
    refs json NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20240218192745');
