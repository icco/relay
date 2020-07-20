-- Auto-generated at Mon, 20 Jul 2020 03:55:50 UTC
-- Please do not change the name attributes

-- name: up
CREATE SEQUENCE locks_rvn OWNED BY locks.record_version_number;

-- name: down
DELETE SEQUENCE locks_rvn;
