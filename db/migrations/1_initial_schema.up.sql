-- CREATE TABLE userinfo
-- (
-- id serial NOT NULL,
-- username character varying(100) NOT NULL,
-- department character varying(500) NOT NULL,
-- created timestamp NOT NULL DEFAULT NOW(),
-- CONSTRAINT userinfo_pkey PRIMARY KEY (id)
-- )
-- WITH (OIDS=FALSE);

CREATE TABLE IF NOT EXISTS account (
id serial PRIMARY KEY,
email character varying(255) NOT NULL DEFAULT '',
password character varying(255) NOT NULL DEFAULT '',
firstname character varying(255) NOT NULL DEFAULT '',
lastname character varying(255) NOT NULL DEFAULT '',
avatar character varying(255) NOT NULL DEFAULT '',
priv bigint DEFAULT 0,
created timestamp NOT NULL DEFAULT now(),
updated timestamp NOT NULL DEFAULT now());
