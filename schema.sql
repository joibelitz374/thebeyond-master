--
-- PostgreSQL database dump
--

-- Dumped from database version 18.0
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


SET default_table_access_method = heap;

--
-- Name: account; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.account (
    id integer NOT NULL,
    key_id text NOT NULL,
    short_id text NOT NULL,
    subscription_expires_at timestamp with time zone DEFAULT (now() + '1 day'::interval),
    region text NOT NULL,
    language text DEFAULT ''::text NOT NULL,
    currency text DEFAULT ''::text NOT NULL,
    protocol text DEFAULT 'vxr'::text NOT NULL,
    server_location text DEFAULT 'de'::text NOT NULL,
    last_key_refresh_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    promo text,
    discount integer DEFAULT 0,
    cluster_id smallint DEFAULT 1,
    service_check_sent smallint DEFAULT 0
);


--
-- Name: account_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.account_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: account_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.account_id_seq OWNED BY public.account.id;


--
-- Name: payment; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.payment (
    account_id integer NOT NULL,
    amount integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    expires_at timestamp with time zone
);


--
-- Name: platform_account; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.platform_account (
    platform_id smallint NOT NULL,
    external_account_id bigint NOT NULL,
    fk_account_id integer NOT NULL
);


--
-- Name: promo; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.promo (
    name text NOT NULL,
    creator integer NOT NULL,
    level smallint DEFAULT 1 NOT NULL,
    discount_client smallint DEFAULT 0 NOT NULL,
    clients smallint DEFAULT 0 NOT NULL,
    buyers smallint DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: promo_buyer; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.promo_buyer (
    promo text NOT NULL,
    fk_buyer_id integer NOT NULL
);


--
-- Name: account id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account ALTER COLUMN id SET DEFAULT nextval('public.account_id_seq'::regclass);


--
-- Name: account account_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_pkey PRIMARY KEY (id);


--
-- Name: platform_account platform_account_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.platform_account
    ADD CONSTRAINT platform_account_pkey PRIMARY KEY (platform_id, external_account_id);


--
-- Name: promo_buyer promo_buyer_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promo_buyer
    ADD CONSTRAINT promo_buyer_pkey PRIMARY KEY (promo, fk_buyer_id);


--
-- Name: promo promo_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promo
    ADD CONSTRAINT promo_pkey PRIMARY KEY (name);


--
-- Name: promo_creator_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX promo_creator_idx ON public.promo USING btree (creator);


--
-- Name: account account_promo_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_promo_fkey FOREIGN KEY (promo) REFERENCES public.promo(name) ON DELETE CASCADE;


--
-- Name: payment payment_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account(id) ON DELETE CASCADE;


--
-- Name: platform_account platform_account_fk_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.platform_account
    ADD CONSTRAINT platform_account_fk_account_id_fkey FOREIGN KEY (fk_account_id) REFERENCES public.account(id) ON DELETE CASCADE;


--
-- Name: promo_buyer promo_buyer_fk_buyer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promo_buyer
    ADD CONSTRAINT promo_buyer_fk_buyer_id_fkey FOREIGN KEY (fk_buyer_id) REFERENCES public.account(id) ON DELETE CASCADE;


--
-- Name: promo_buyer promo_buyer_promo_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promo_buyer
    ADD CONSTRAINT promo_buyer_promo_fkey FOREIGN KEY (promo) REFERENCES public.promo(name) ON DELETE CASCADE;


--
-- Name: promo promo_creator_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promo
    ADD CONSTRAINT promo_creator_fkey FOREIGN KEY (creator) REFERENCES public.account(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--
