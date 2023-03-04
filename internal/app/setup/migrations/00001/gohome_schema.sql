--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.1

-- Started on 2023-03-04 13:55:58 UTC

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

DROP DATABASE gohome;
--
-- TOC entry 3489 (class 1262 OID 16384)
-- Name: gohome; Type: DATABASE; Schema: -; Owner: levi
--

CREATE DATABASE gohome WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE gohome OWNER TO levi;

\connect gohome

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
-- TOC entry 222 (class 1259 OID 16428)
-- Name: controlpoint; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.controlpoint (
    id bigint NOT NULL,
    name text,
    ipaddress text,
    mac text
);


ALTER TABLE public.controlpoint OWNER TO levi;

--
-- TOC entry 240 (class 1259 OID 16638)
-- Name: controlpoint_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.controlpoint_id_seq
    START WITH 4
    INCREMENT BY 1
    MINVALUE 4
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.controlpoint_id_seq OWNER TO levi;

--
-- TOC entry 3490 (class 0 OID 0)
-- Dependencies: 240
-- Name: controlpoint_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.controlpoint_id_seq OWNED BY public.controlpoint.id;


--
-- TOC entry 223 (class 1259 OID 16433)
-- Name: controlpointnodes; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.controlpointnodes (
    id bigint NOT NULL,
    controlpointid bigint,
    nodeid bigint
);


ALTER TABLE public.controlpointnodes OWNER TO levi;

--
-- TOC entry 241 (class 1259 OID 16640)
-- Name: controlpointnodes_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.controlpointnodes_id_seq
    START WITH 49
    INCREMENT BY 1
    MINVALUE 49
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.controlpointnodes_id_seq OWNER TO levi;

--
-- TOC entry 3491 (class 0 OID 0)
-- Dependencies: 241
-- Name: controlpointnodes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.controlpointnodes_id_seq OWNED BY public.controlpointnodes.id;


--
-- TOC entry 230 (class 1259 OID 16462)
-- Name: magneticlog; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.magneticlog (
    id bigint NOT NULL,
    nodesensorlogid bigint,
    isclosed bigint DEFAULT '0'::bigint
);


ALTER TABLE public.magneticlog OWNER TO levi;

--
-- TOC entry 234 (class 1259 OID 16626)
-- Name: magneticlog_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.magneticlog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.magneticlog_id_seq OWNER TO levi;

--
-- TOC entry 3492 (class 0 OID 0)
-- Dependencies: 234
-- Name: magneticlog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.magneticlog_id_seq OWNED BY public.magneticlog.id;


--
-- TOC entry 229 (class 1259 OID 16459)
-- Name: moisturelog; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.moisturelog (
    id bigint NOT NULL,
    nodesensorlogid bigint,
    moisture bigint
);


ALTER TABLE public.moisturelog OWNER TO levi;

--
-- TOC entry 235 (class 1259 OID 16628)
-- Name: moisturelog_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.moisturelog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.moisturelog_id_seq OWNER TO levi;

--
-- TOC entry 3493 (class 0 OID 0)
-- Dependencies: 235
-- Name: moisturelog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.moisturelog_id_seq OWNED BY public.moisturelog.id;


--
-- TOC entry 219 (class 1259 OID 16411)
-- Name: node; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.node (
    id bigint NOT NULL,
    mac text,
    name text
);


ALTER TABLE public.node OWNER TO levi;

--
-- TOC entry 242 (class 1259 OID 16642)
-- Name: node_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.node_id_seq
    START WITH 49
    INCREMENT BY 1
    MINVALUE 49
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.node_id_seq OWNER TO levi;

--
-- TOC entry 3494 (class 0 OID 0)
-- Dependencies: 242
-- Name: node_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.node_id_seq OWNED BY public.node.id;


--
-- TOC entry 220 (class 1259 OID 16416)
-- Name: nodesensor; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.nodesensor (
    id bigint NOT NULL,
    nodeid bigint,
    sensortypeid bigint,
    name text,
    pin bigint,
    dhttype bigint
);


ALTER TABLE public.nodesensor OWNER TO levi;

--
-- TOC entry 243 (class 1259 OID 16644)
-- Name: nodesensor_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.nodesensor_id_seq
    START WITH 55
    INCREMENT BY 1
    MINVALUE 55
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.nodesensor_id_seq OWNER TO levi;

--
-- TOC entry 3495 (class 0 OID 0)
-- Dependencies: 243
-- Name: nodesensor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.nodesensor_id_seq OWNED BY public.nodesensor.id;


--
-- TOC entry 227 (class 1259 OID 16451)
-- Name: nodesensorlog; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.nodesensorlog (
    id bigint NOT NULL,
    nodeid bigint,
    datelogged text
);


ALTER TABLE public.nodesensorlog OWNER TO levi;

--
-- TOC entry 232 (class 1259 OID 16622)
-- Name: nodesensorlog_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.nodesensorlog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.nodesensorlog_id_seq OWNER TO levi;

--
-- TOC entry 3496 (class 0 OID 0)
-- Dependencies: 232
-- Name: nodesensorlog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.nodesensorlog_id_seq OWNED BY public.nodesensorlog.id;


--
-- TOC entry 221 (class 1259 OID 16421)
-- Name: nodeswitch; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.nodeswitch (
    id bigint NOT NULL,
    nodeid bigint,
    switchtypeid bigint,
    name text,
    pin bigint,
    momentarypressduration bigint DEFAULT '100'::bigint,
    isclosedon bigint DEFAULT '1'::bigint
);


ALTER TABLE public.nodeswitch OWNER TO levi;

--
-- TOC entry 244 (class 1259 OID 16646)
-- Name: nodeswitch_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.nodeswitch_id_seq
    START WITH 29
    INCREMENT BY 1
    MINVALUE 29
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.nodeswitch_id_seq OWNER TO levi;

--
-- TOC entry 3497 (class 0 OID 0)
-- Dependencies: 244
-- Name: nodeswitch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.nodeswitch_id_seq OWNED BY public.nodeswitch.id;


--
-- TOC entry 231 (class 1259 OID 16466)
-- Name: resistorlog; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.resistorlog (
    id bigint NOT NULL,
    nodesensorlogid bigint,
    resistorvalue bigint
);


ALTER TABLE public.resistorlog OWNER TO levi;

--
-- TOC entry 236 (class 1259 OID 16630)
-- Name: resistorlog_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.resistorlog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.resistorlog_id_seq OWNER TO levi;

--
-- TOC entry 3498 (class 0 OID 0)
-- Dependencies: 236
-- Name: resistorlog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.resistorlog_id_seq OWNED BY public.resistorlog.id;


--
-- TOC entry 216 (class 1259 OID 16396)
-- Name: sensortype; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.sensortype (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.sensortype OWNER TO levi;

--
-- TOC entry 245 (class 1259 OID 16648)
-- Name: sensortype_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.sensortype_id_seq
    START WITH 5
    INCREMENT BY 1
    MINVALUE 5
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sensortype_id_seq OWNER TO levi;

--
-- TOC entry 3499 (class 0 OID 0)
-- Dependencies: 245
-- Name: sensortype_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.sensortype_id_seq OWNED BY public.sensortype.id;


--
-- TOC entry 217 (class 1259 OID 16401)
-- Name: sensortypedata; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.sensortypedata (
    id bigint NOT NULL,
    sensortypeid bigint,
    name text,
    valuetype text
);


ALTER TABLE public.sensortypedata OWNER TO levi;

--
-- TOC entry 218 (class 1259 OID 16406)
-- Name: switchtype; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.switchtype (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.switchtype OWNER TO levi;

--
-- TOC entry 228 (class 1259 OID 16456)
-- Name: templog; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.templog (
    id bigint NOT NULL,
    nodesensorlogid bigint,
    temperaturef real,
    temperaturec real,
    humidity real
);


ALTER TABLE public.templog OWNER TO levi;

--
-- TOC entry 233 (class 1259 OID 16624)
-- Name: templog_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.templog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.templog_id_seq OWNER TO levi;

--
-- TOC entry 3500 (class 0 OID 0)
-- Dependencies: 233
-- Name: templog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.templog_id_seq OWNED BY public.templog.id;


--
-- TOC entry 226 (class 1259 OID 16446)
-- Name: viewnodeswitchdata; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.viewnodeswitchdata (
    id bigint NOT NULL,
    nodeid bigint,
    viewid bigint,
    nodeswitchid bigint,
    name text
);


ALTER TABLE public.viewnodeswitchdata OWNER TO levi;

--
-- TOC entry 239 (class 1259 OID 16636)
-- Name: viewnodeswitchdata_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.viewnodeswitchdata_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.viewnodeswitchdata_id_seq OWNER TO levi;

--
-- TOC entry 3501 (class 0 OID 0)
-- Dependencies: 239
-- Name: viewnodeswitchdata_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.viewnodeswitchdata_id_seq OWNED BY public.viewnodeswitchdata.id;


--
-- TOC entry 224 (class 1259 OID 16436)
-- Name: view; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.view (
    id bigint DEFAULT nextval('public.viewnodeswitchdata_id_seq'::regclass) NOT NULL,
    name text
);


ALTER TABLE public.view OWNER TO levi;

--
-- TOC entry 238 (class 1259 OID 16634)
-- Name: view_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.view_id_seq
    START WITH 11
    INCREMENT BY 1
    MINVALUE 11
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.view_id_seq OWNER TO levi;

--
-- TOC entry 3502 (class 0 OID 0)
-- Dependencies: 238
-- Name: view_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.view_id_seq OWNED BY public.view.id;


--
-- TOC entry 225 (class 1259 OID 16441)
-- Name: viewnodesensordata; Type: TABLE; Schema: public; Owner: levi
--

CREATE TABLE public.viewnodesensordata (
    id bigint NOT NULL,
    nodeid bigint,
    viewid bigint,
    nodesensorid bigint,
    name text
);


ALTER TABLE public.viewnodesensordata OWNER TO levi;

--
-- TOC entry 237 (class 1259 OID 16632)
-- Name: viewnodesensordata_id_seq; Type: SEQUENCE; Schema: public; Owner: levi
--

CREATE SEQUENCE public.viewnodesensordata_id_seq
    START WITH 31
    INCREMENT BY 1
    MINVALUE 31
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.viewnodesensordata_id_seq OWNER TO levi;

--
-- TOC entry 3503 (class 0 OID 0)
-- Dependencies: 237
-- Name: viewnodesensordata_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: levi
--

ALTER SEQUENCE public.viewnodesensordata_id_seq OWNED BY public.viewnodesensordata.id;


--
-- TOC entry 3257 (class 2604 OID 16639)
-- Name: controlpoint id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.controlpoint ALTER COLUMN id SET DEFAULT nextval('public.controlpoint_id_seq'::regclass);


--
-- TOC entry 3258 (class 2604 OID 16641)
-- Name: controlpointnodes id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.controlpointnodes ALTER COLUMN id SET DEFAULT nextval('public.controlpointnodes_id_seq'::regclass);


--
-- TOC entry 3265 (class 2604 OID 16627)
-- Name: magneticlog id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.magneticlog ALTER COLUMN id SET DEFAULT nextval('public.magneticlog_id_seq'::regclass);


--
-- TOC entry 3264 (class 2604 OID 16629)
-- Name: moisturelog id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.moisturelog ALTER COLUMN id SET DEFAULT nextval('public.moisturelog_id_seq'::regclass);


--
-- TOC entry 3252 (class 2604 OID 16643)
-- Name: node id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.node ALTER COLUMN id SET DEFAULT nextval('public.node_id_seq'::regclass);


--
-- TOC entry 3253 (class 2604 OID 16645)
-- Name: nodesensor id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensor ALTER COLUMN id SET DEFAULT nextval('public.nodesensor_id_seq'::regclass);


--
-- TOC entry 3262 (class 2604 OID 16623)
-- Name: nodesensorlog id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensorlog ALTER COLUMN id SET DEFAULT nextval('public.nodesensorlog_id_seq'::regclass);


--
-- TOC entry 3254 (class 2604 OID 16647)
-- Name: nodeswitch id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodeswitch ALTER COLUMN id SET DEFAULT nextval('public.nodeswitch_id_seq'::regclass);


--
-- TOC entry 3267 (class 2604 OID 16631)
-- Name: resistorlog id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.resistorlog ALTER COLUMN id SET DEFAULT nextval('public.resistorlog_id_seq'::regclass);


--
-- TOC entry 3263 (class 2604 OID 16625)
-- Name: templog id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.templog ALTER COLUMN id SET DEFAULT nextval('public.templog_id_seq'::regclass);


--
-- TOC entry 3260 (class 2604 OID 16633)
-- Name: viewnodesensordata id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodesensordata ALTER COLUMN id SET DEFAULT nextval('public.viewnodesensordata_id_seq'::regclass);


--
-- TOC entry 3261 (class 2604 OID 16650)
-- Name: viewnodeswitchdata id; Type: DEFAULT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodeswitchdata ALTER COLUMN id SET DEFAULT nextval('public.viewnodeswitchdata_id_seq'::regclass);


--
-- TOC entry 3308 (class 2606 OID 16621)
-- Name: nodesensorlog id_is_unique; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensorlog
    ADD CONSTRAINT id_is_unique UNIQUE (id);


--
-- TOC entry 3269 (class 2606 OID 16513)
-- Name: sensortype idx_16396_sensortype_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.sensortype
    ADD CONSTRAINT idx_16396_sensortype_pkey PRIMARY KEY (id);


--
-- TOC entry 3273 (class 2606 OID 16512)
-- Name: sensortypedata idx_16401_sensortypedata_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.sensortypedata
    ADD CONSTRAINT idx_16401_sensortypedata_pkey PRIMARY KEY (id);


--
-- TOC entry 3278 (class 2606 OID 16517)
-- Name: switchtype idx_16406_switchtype_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.switchtype
    ADD CONSTRAINT idx_16406_switchtype_pkey PRIMARY KEY (id);


--
-- TOC entry 3280 (class 2606 OID 16514)
-- Name: node idx_16411_node_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT idx_16411_node_pkey PRIMARY KEY (id);


--
-- TOC entry 3285 (class 2606 OID 16518)
-- Name: nodesensor idx_16416_nodesensor_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensor
    ADD CONSTRAINT idx_16416_nodesensor_pkey PRIMARY KEY (id);


--
-- TOC entry 3288 (class 2606 OID 16515)
-- Name: nodeswitch idx_16421_nodeswitch_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodeswitch
    ADD CONSTRAINT idx_16421_nodeswitch_pkey PRIMARY KEY (id);


--
-- TOC entry 3291 (class 2606 OID 16523)
-- Name: controlpoint idx_16428_controlpoint_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.controlpoint
    ADD CONSTRAINT idx_16428_controlpoint_pkey PRIMARY KEY (id);


--
-- TOC entry 3294 (class 2606 OID 16519)
-- Name: controlpointnodes idx_16433_controlpointnodes_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.controlpointnodes
    ADD CONSTRAINT idx_16433_controlpointnodes_pkey PRIMARY KEY (id);


--
-- TOC entry 3298 (class 2606 OID 16522)
-- Name: view idx_16436_view_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.view
    ADD CONSTRAINT idx_16436_view_pkey PRIMARY KEY (id);


--
-- TOC entry 3302 (class 2606 OID 16516)
-- Name: viewnodesensordata idx_16441_viewnodesensordata_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodesensordata
    ADD CONSTRAINT idx_16441_viewnodesensordata_pkey PRIMARY KEY (id);


--
-- TOC entry 3306 (class 2606 OID 16526)
-- Name: viewnodeswitchdata idx_16446_viewnodeswitchdata_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodeswitchdata
    ADD CONSTRAINT idx_16446_viewnodeswitchdata_pkey PRIMARY KEY (id);


--
-- TOC entry 3310 (class 2606 OID 16521)
-- Name: nodesensorlog idx_16451_nodesensorlog_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensorlog
    ADD CONSTRAINT idx_16451_nodesensorlog_pkey PRIMARY KEY (id);


--
-- TOC entry 3314 (class 2606 OID 16524)
-- Name: templog idx_16456_templog_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.templog
    ADD CONSTRAINT idx_16456_templog_pkey PRIMARY KEY (id);


--
-- TOC entry 3316 (class 2606 OID 16520)
-- Name: moisturelog idx_16459_moisturelog_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.moisturelog
    ADD CONSTRAINT idx_16459_moisturelog_pkey PRIMARY KEY (id);


--
-- TOC entry 3319 (class 2606 OID 16527)
-- Name: magneticlog idx_16462_magneticlog_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.magneticlog
    ADD CONSTRAINT idx_16462_magneticlog_pkey PRIMARY KEY (id);


--
-- TOC entry 3322 (class 2606 OID 16525)
-- Name: resistorlog idx_16466_resistorlog_pkey; Type: CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.resistorlog
    ADD CONSTRAINT idx_16466_resistorlog_pkey PRIMARY KEY (id);


--
-- TOC entry 3270 (class 1259 OID 16473)
-- Name: idx_16396_sqlite_autoindex_sensortype_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16396_sqlite_autoindex_sensortype_1 ON public.sensortype USING btree (id);


--
-- TOC entry 3271 (class 1259 OID 16471)
-- Name: idx_16396_sqlite_autoindex_sensortype_2; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16396_sqlite_autoindex_sensortype_2 ON public.sensortype USING btree (name);


--
-- TOC entry 3274 (class 1259 OID 16472)
-- Name: idx_16401_sqlite_autoindex_sensortypedata_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16401_sqlite_autoindex_sensortypedata_1 ON public.sensortypedata USING btree (id);


--
-- TOC entry 3275 (class 1259 OID 16485)
-- Name: idx_16406_sqlite_autoindex_switchtype_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16406_sqlite_autoindex_switchtype_1 ON public.switchtype USING btree (id);


--
-- TOC entry 3276 (class 1259 OID 16482)
-- Name: idx_16406_sqlite_autoindex_switchtype_2; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16406_sqlite_autoindex_switchtype_2 ON public.switchtype USING btree (name);


--
-- TOC entry 3281 (class 1259 OID 16477)
-- Name: idx_16411_sqlite_autoindex_node_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16411_sqlite_autoindex_node_1 ON public.node USING btree (id);


--
-- TOC entry 3282 (class 1259 OID 16475)
-- Name: idx_16411_sqlite_autoindex_node_2; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16411_sqlite_autoindex_node_2 ON public.node USING btree (mac);


--
-- TOC entry 3283 (class 1259 OID 16474)
-- Name: idx_16411_sqlite_autoindex_node_3; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16411_sqlite_autoindex_node_3 ON public.node USING btree (name);


--
-- TOC entry 3286 (class 1259 OID 16487)
-- Name: idx_16416_sqlite_autoindex_nodesensor_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16416_sqlite_autoindex_nodesensor_1 ON public.nodesensor USING btree (id);


--
-- TOC entry 3289 (class 1259 OID 16479)
-- Name: idx_16421_sqlite_autoindex_nodeswitch_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16421_sqlite_autoindex_nodeswitch_1 ON public.nodeswitch USING btree (id);


--
-- TOC entry 3292 (class 1259 OID 16496)
-- Name: idx_16428_sqlite_autoindex_controlpoint_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16428_sqlite_autoindex_controlpoint_1 ON public.controlpoint USING btree (id);


--
-- TOC entry 3295 (class 1259 OID 16489)
-- Name: idx_16433_sqlite_autoindex_controlpointnodes_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16433_sqlite_autoindex_controlpointnodes_1 ON public.controlpointnodes USING btree (id);


--
-- TOC entry 3296 (class 1259 OID 16497)
-- Name: idx_16436_sqlite_autoindex_view_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16436_sqlite_autoindex_view_1 ON public.view USING btree (id);


--
-- TOC entry 3299 (class 1259 OID 16480)
-- Name: idx_16441_sqlite_autoindex_viewnodesensordata_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16441_sqlite_autoindex_viewnodesensordata_1 ON public.viewnodesensordata USING btree (id);


--
-- TOC entry 3300 (class 1259 OID 16481)
-- Name: idx_16441_sqlite_autoindex_viewnodesensordata_2; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16441_sqlite_autoindex_viewnodesensordata_2 ON public.viewnodesensordata USING btree (nodeid, viewid, nodesensorid);


--
-- TOC entry 3303 (class 1259 OID 16502)
-- Name: idx_16446_sqlite_autoindex_viewnodeswitchdata_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16446_sqlite_autoindex_viewnodeswitchdata_1 ON public.viewnodeswitchdata USING btree (id);


--
-- TOC entry 3304 (class 1259 OID 16503)
-- Name: idx_16446_sqlite_autoindex_viewnodeswitchdata_2; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16446_sqlite_autoindex_viewnodeswitchdata_2 ON public.viewnodeswitchdata USING btree (nodeid, viewid, nodeswitchid);


--
-- TOC entry 3311 (class 1259 OID 16493)
-- Name: idx_16451_sqlite_autoindex_nodesensorlog_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16451_sqlite_autoindex_nodesensorlog_1 ON public.nodesensorlog USING btree (id);


--
-- TOC entry 3312 (class 1259 OID 16500)
-- Name: idx_16456_sqlite_autoindex_templog_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16456_sqlite_autoindex_templog_1 ON public.templog USING btree (id);


--
-- TOC entry 3317 (class 1259 OID 16492)
-- Name: idx_16459_sqlite_autoindex_moisturelog_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16459_sqlite_autoindex_moisturelog_1 ON public.moisturelog USING btree (id);


--
-- TOC entry 3320 (class 1259 OID 16504)
-- Name: idx_16462_sqlite_autoindex_magneticlog_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16462_sqlite_autoindex_magneticlog_1 ON public.magneticlog USING btree (id);


--
-- TOC entry 3323 (class 1259 OID 16501)
-- Name: idx_16466_sqlite_autoindex_resistorlog_1; Type: INDEX; Schema: public; Owner: levi
--

CREATE UNIQUE INDEX idx_16466_sqlite_autoindex_resistorlog_1 ON public.resistorlog USING btree (id);


--
-- TOC entry 3329 (class 2606 OID 16553)
-- Name: controlpointnodes controlpointnodes_controlpointid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.controlpointnodes
    ADD CONSTRAINT controlpointnodes_controlpointid_fkey FOREIGN KEY (controlpointid) REFERENCES public.controlpoint(id);


--
-- TOC entry 3330 (class 2606 OID 16558)
-- Name: controlpointnodes controlpointnodes_nodeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.controlpointnodes
    ADD CONSTRAINT controlpointnodes_nodeid_fkey FOREIGN KEY (nodeid) REFERENCES public.node(id);


--
-- TOC entry 3340 (class 2606 OID 16608)
-- Name: magneticlog magneticlog_nodesensorlogid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.magneticlog
    ADD CONSTRAINT magneticlog_nodesensorlogid_fkey FOREIGN KEY (nodesensorlogid) REFERENCES public.nodesensorlog(id);


--
-- TOC entry 3339 (class 2606 OID 16603)
-- Name: moisturelog moisturelog_nodesensorlogid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.moisturelog
    ADD CONSTRAINT moisturelog_nodesensorlogid_fkey FOREIGN KEY (nodesensorlogid) REFERENCES public.nodesensorlog(id);


--
-- TOC entry 3325 (class 2606 OID 16538)
-- Name: nodesensor nodesensor_nodeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensor
    ADD CONSTRAINT nodesensor_nodeid_fkey FOREIGN KEY (nodeid) REFERENCES public.node(id);


--
-- TOC entry 3326 (class 2606 OID 16533)
-- Name: nodesensor nodesensor_sensortypeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensor
    ADD CONSTRAINT nodesensor_sensortypeid_fkey FOREIGN KEY (sensortypeid) REFERENCES public.sensortype(id);


--
-- TOC entry 3337 (class 2606 OID 16593)
-- Name: nodesensorlog nodesensorlog_nodeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodesensorlog
    ADD CONSTRAINT nodesensorlog_nodeid_fkey FOREIGN KEY (nodeid) REFERENCES public.node(id);


--
-- TOC entry 3327 (class 2606 OID 16548)
-- Name: nodeswitch nodeswitch_nodeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodeswitch
    ADD CONSTRAINT nodeswitch_nodeid_fkey FOREIGN KEY (nodeid) REFERENCES public.node(id);


--
-- TOC entry 3328 (class 2606 OID 16543)
-- Name: nodeswitch nodeswitch_switchtypeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.nodeswitch
    ADD CONSTRAINT nodeswitch_switchtypeid_fkey FOREIGN KEY (switchtypeid) REFERENCES public.switchtype(id);


--
-- TOC entry 3341 (class 2606 OID 16613)
-- Name: resistorlog resistorlog_nodesensorlogid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.resistorlog
    ADD CONSTRAINT resistorlog_nodesensorlogid_fkey FOREIGN KEY (nodesensorlogid) REFERENCES public.nodesensorlog(id);


--
-- TOC entry 3324 (class 2606 OID 16528)
-- Name: sensortypedata sensortypedata_sensortypeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.sensortypedata
    ADD CONSTRAINT sensortypedata_sensortypeid_fkey FOREIGN KEY (sensortypeid) REFERENCES public.sensortype(id);


--
-- TOC entry 3338 (class 2606 OID 16598)
-- Name: templog templog_nodesensorlogid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.templog
    ADD CONSTRAINT templog_nodesensorlogid_fkey FOREIGN KEY (nodesensorlogid) REFERENCES public.nodesensorlog(id);


--
-- TOC entry 3331 (class 2606 OID 16573)
-- Name: viewnodesensordata viewnodesensordata_nodeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodesensordata
    ADD CONSTRAINT viewnodesensordata_nodeid_fkey FOREIGN KEY (nodeid) REFERENCES public.node(id);


--
-- TOC entry 3332 (class 2606 OID 16563)
-- Name: viewnodesensordata viewnodesensordata_nodesensorid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodesensordata
    ADD CONSTRAINT viewnodesensordata_nodesensorid_fkey FOREIGN KEY (nodesensorid) REFERENCES public.nodesensor(id);


--
-- TOC entry 3333 (class 2606 OID 16568)
-- Name: viewnodesensordata viewnodesensordata_viewid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodesensordata
    ADD CONSTRAINT viewnodesensordata_viewid_fkey FOREIGN KEY (viewid) REFERENCES public.view(id);


--
-- TOC entry 3334 (class 2606 OID 16578)
-- Name: viewnodeswitchdata viewnodeswitchdata_nodeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodeswitchdata
    ADD CONSTRAINT viewnodeswitchdata_nodeid_fkey FOREIGN KEY (nodeid) REFERENCES public.node(id);


--
-- TOC entry 3335 (class 2606 OID 16583)
-- Name: viewnodeswitchdata viewnodeswitchdata_nodeswitchid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodeswitchdata
    ADD CONSTRAINT viewnodeswitchdata_nodeswitchid_fkey FOREIGN KEY (nodeswitchid) REFERENCES public.nodeswitch(id);


--
-- TOC entry 3336 (class 2606 OID 16588)
-- Name: viewnodeswitchdata viewnodeswitchdata_viewid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: levi
--

ALTER TABLE ONLY public.viewnodeswitchdata
    ADD CONSTRAINT viewnodeswitchdata_viewid_fkey FOREIGN KEY (viewid) REFERENCES public.view(id);


--
-- TOC entry 2110 (class 826 OID 16619)
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: public; Owner: levi
--

ALTER DEFAULT PRIVILEGES FOR ROLE levi IN SCHEMA public GRANT ALL ON TABLES  TO levi;


-- Completed on 2023-03-04 13:55:58 UTC

--
-- PostgreSQL database dump complete
--

