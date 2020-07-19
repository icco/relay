-- Auto-generated at Sun, 19 Jul 2020 21:29:17 UTC
-- Please do not change the name attributes

-- name: up
CREATE TABLE IF NOT EXISTS migrations (
 id          VARCHAR(15) NOT NULL PRIMARY KEY,
 description TEXT        NOT NULL,
 created_at  TIMESTAMP   NOT NULL
);

-- name: down
DROP TABLE IF EXISTS migrations;
