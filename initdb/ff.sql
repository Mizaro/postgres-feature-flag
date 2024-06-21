--
-- PostgreSQL port of the MySQL "World" database.
--
-- The sample data used in the world database is Copyright Statistics
-- Finland, http://www.stat.fi/worldinfigures.
--

BEGIN;

SET client_encoding = 'LATIN1';


CREATE TABLE feature (
                      id integer NOT NULL,
                      name text NOT NULL,
                      value text NOT NULL
);

---

COPY feature (id, name, value) FROM stdin;
1	v2_enabled	false
\.

---

ALTER TABLE ONLY feature
    ADD CONSTRAINT feature_pkey PRIMARY KEY (id);

COMMIT;

ANALYZE feature;
