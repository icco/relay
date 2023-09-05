-- Auto-generated at Tue, 05 Sep 2023 03:23:33 UTC
-- Please do not change the name attributes

-- name: up
ALTER TABLE messages DROP CONSTRAINT messages_message_key;

-- name: down
ALTER TABLE messages ADD CONSTRAINT messages_message_key UNIQUE (message);
