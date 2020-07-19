-- Auto-generated at Sun, 19 Jul 2020 21:54:23 UTC
-- Please do not change the name attributes

-- name: up
CREATE TABLE IF NOT EXISTS locks (
			name CHARACTER VARYING(255) PRIMARY KEY,
			record_version_number BIGINT,
			data BYTEA,
			owner CHARACTER VARYING(255)

-- name: down
DROP TABLE IF EXISTS locks;

