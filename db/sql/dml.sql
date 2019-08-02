--
-- PostgreSQL database dump
--

-- Dumped from database version 11.4
-- Dumped by pg_dump version 11.4 (Debian 11.4-1.pgdg90+1)

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
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password) FROM stdin;
1	sample	user@example.com	password
2	sample2	user2@example.com	password
3	sample3	user3@example.com	password
4	sample4	user4@example.com	password
5	sample5	user5@example.com	password
\.


--
-- Data for Name: line_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.line_users (id, user_id, line_id, linked_account) FROM stdin;
\.


--
-- Data for Name: nonces; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.nonces (id, user_id, nonce, linked_account) FROM stdin;
\.


--
-- Name: line_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.line_users_id_seq', 1, false);


--
-- Name: nonces_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.nonces_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- PostgreSQL database dump complete
--

