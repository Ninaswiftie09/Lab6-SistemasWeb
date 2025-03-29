--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

-- Started on 2025-03-29 00:45:37

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 226 (class 1259 OID 16513)
-- Name: extra_time; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.extra_time (
    id integer NOT NULL,
    match_id integer,
    extra_time integer
);


ALTER TABLE public.extra_time OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16512)
-- Name: extra_time_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.extra_time_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.extra_time_id_seq OWNER TO postgres;

--
-- TOC entry 4941 (class 0 OID 0)
-- Dependencies: 225
-- Name: extra_time_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.extra_time_id_seq OWNED BY public.extra_time.id;


--
-- TOC entry 220 (class 1259 OID 16471)
-- Name: goals; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goals (
    id integer NOT NULL,
    match_id integer,
    team character varying(255),
    player character varying(255),
    minute integer
);


ALTER TABLE public.goals OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16470)
-- Name: goals_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goals_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.goals_id_seq OWNER TO postgres;

--
-- TOC entry 4942 (class 0 OID 0)
-- Dependencies: 219
-- Name: goals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.goals_id_seq OWNED BY public.goals.id;


--
-- TOC entry 218 (class 1259 OID 16462)
-- Name: matches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.matches (
    id integer NOT NULL,
    home_team character varying(255),
    away_team character varying(255),
    match_date date
);


ALTER TABLE public.matches OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16461)
-- Name: matches_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.matches_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.matches_id_seq OWNER TO postgres;

--
-- TOC entry 4943 (class 0 OID 0)
-- Dependencies: 217
-- Name: matches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.matches_id_seq OWNED BY public.matches.id;


--
-- TOC entry 224 (class 1259 OID 16499)
-- Name: red_cards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.red_cards (
    id integer NOT NULL,
    match_id integer,
    team character varying(255),
    player character varying(255),
    minute integer
);


ALTER TABLE public.red_cards OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16498)
-- Name: red_cards_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.red_cards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.red_cards_id_seq OWNER TO postgres;

--
-- TOC entry 4944 (class 0 OID 0)
-- Dependencies: 223
-- Name: red_cards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.red_cards_id_seq OWNED BY public.red_cards.id;


--
-- TOC entry 222 (class 1259 OID 16485)
-- Name: yellow_cards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.yellow_cards (
    id integer NOT NULL,
    match_id integer,
    team character varying(255),
    player character varying(255),
    minute integer
);


ALTER TABLE public.yellow_cards OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16484)
-- Name: yellow_cards_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.yellow_cards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.yellow_cards_id_seq OWNER TO postgres;

--
-- TOC entry 4945 (class 0 OID 0)
-- Dependencies: 221
-- Name: yellow_cards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.yellow_cards_id_seq OWNED BY public.yellow_cards.id;


--
-- TOC entry 4766 (class 2604 OID 16516)
-- Name: extra_time id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.extra_time ALTER COLUMN id SET DEFAULT nextval('public.extra_time_id_seq'::regclass);


--
-- TOC entry 4763 (class 2604 OID 16474)
-- Name: goals id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goals ALTER COLUMN id SET DEFAULT nextval('public.goals_id_seq'::regclass);


--
-- TOC entry 4762 (class 2604 OID 16465)
-- Name: matches id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches ALTER COLUMN id SET DEFAULT nextval('public.matches_id_seq'::regclass);


--
-- TOC entry 4765 (class 2604 OID 16502)
-- Name: red_cards id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.red_cards ALTER COLUMN id SET DEFAULT nextval('public.red_cards_id_seq'::regclass);


--
-- TOC entry 4764 (class 2604 OID 16488)
-- Name: yellow_cards id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.yellow_cards ALTER COLUMN id SET DEFAULT nextval('public.yellow_cards_id_seq'::regclass);


--
-- TOC entry 4935 (class 0 OID 16513)
-- Dependencies: 226
-- Data for Name: extra_time; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.extra_time (id, match_id, extra_time) FROM stdin;
1	1	5
2	2	3
\.


--
-- TOC entry 4929 (class 0 OID 16471)
-- Dependencies: 220
-- Data for Name: goals; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goals (id, match_id, team, player, minute) FROM stdin;
1	1	Real Madrid	Karim Benzema	23
2	1	Barcelona	Lionel Messi	45
3	2	Atletico Madrid	Antoine Griezmann	50
4	2	Valencia	Gonçalo Guedes	70
\.


--
-- TOC entry 4927 (class 0 OID 16462)
-- Dependencies: 218
-- Data for Name: matches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.matches (id, home_team, away_team, match_date) FROM stdin;
1	Real Madrid	Barcelona	2025-04-01
2	Atletico Madrid	Valencia	2025-04-02
\.


--
-- TOC entry 4933 (class 0 OID 16499)
-- Dependencies: 224
-- Data for Name: red_cards; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.red_cards (id, match_id, team, player, minute) FROM stdin;
1	1	Real Madrid	Casemiro	80
2	2	Valencia	Ezequiel Garay	85
\.


--
-- TOC entry 4931 (class 0 OID 16485)
-- Dependencies: 222
-- Data for Name: yellow_cards; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.yellow_cards (id, match_id, team, player, minute) FROM stdin;
1	1	Real Madrid	Sergio Ramos	30
2	1	Barcelona	Gerard Piqué	60
3	2	Atletico Madrid	Koke	25
4	2	Valencia	Daniel Parejo	65
\.


--
-- TOC entry 4946 (class 0 OID 0)
-- Dependencies: 225
-- Name: extra_time_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.extra_time_id_seq', 2, true);


--
-- TOC entry 4947 (class 0 OID 0)
-- Dependencies: 219
-- Name: goals_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goals_id_seq', 4, true);


--
-- TOC entry 4948 (class 0 OID 0)
-- Dependencies: 217
-- Name: matches_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.matches_id_seq', 3, true);


--
-- TOC entry 4949 (class 0 OID 0)
-- Dependencies: 223
-- Name: red_cards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.red_cards_id_seq', 2, true);


--
-- TOC entry 4950 (class 0 OID 0)
-- Dependencies: 221
-- Name: yellow_cards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.yellow_cards_id_seq', 4, true);


--
-- TOC entry 4776 (class 2606 OID 16518)
-- Name: extra_time extra_time_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.extra_time
    ADD CONSTRAINT extra_time_pkey PRIMARY KEY (id);


--
-- TOC entry 4770 (class 2606 OID 16478)
-- Name: goals goals_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goals_pkey PRIMARY KEY (id);


--
-- TOC entry 4768 (class 2606 OID 16469)
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- TOC entry 4774 (class 2606 OID 16506)
-- Name: red_cards red_cards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.red_cards
    ADD CONSTRAINT red_cards_pkey PRIMARY KEY (id);


--
-- TOC entry 4772 (class 2606 OID 16492)
-- Name: yellow_cards yellow_cards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.yellow_cards
    ADD CONSTRAINT yellow_cards_pkey PRIMARY KEY (id);


--
-- TOC entry 4780 (class 2606 OID 16519)
-- Name: extra_time extra_time_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.extra_time
    ADD CONSTRAINT extra_time_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON DELETE CASCADE;


--
-- TOC entry 4777 (class 2606 OID 16479)
-- Name: goals goals_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goals_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON DELETE CASCADE;


--
-- TOC entry 4779 (class 2606 OID 16507)
-- Name: red_cards red_cards_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.red_cards
    ADD CONSTRAINT red_cards_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON DELETE CASCADE;


--
-- TOC entry 4778 (class 2606 OID 16493)
-- Name: yellow_cards yellow_cards_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.yellow_cards
    ADD CONSTRAINT yellow_cards_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON DELETE CASCADE;


-- Completed on 2025-03-29 00:45:37

--
-- PostgreSQL database dump complete
--

