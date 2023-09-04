-- Auto-generated at Mon, 04 Sep 2023 21:44:14 UTC
-- Please do not change the name attributes

-- name: up
ALTER TABLE messages ADD COLUMN content TEXT NOT NULL;

-- name: down
ALTER TABLE messages DROP COLUMN content;