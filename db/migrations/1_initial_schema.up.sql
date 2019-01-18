CREATE TABLE userinfo
(
id serial NOT NULL,
username character varying(100) NOT NULL,
department character varying(500) NOT NULL,
created timestamp NOT NULL DEFAULT NOW(),
CONSTRAINT userinfo_pkey PRIMARY KEY (id)
)
WITH (OIDS=FALSE);
