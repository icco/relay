-- Auto-generated at Mon, 04 Sep 2023 19:42:31 UTC
-- Please do not change the name attributes

-- name: up
CREATE TABLE IF NOT EXISTS messages (
  ID SERIAL PRIMARY KEY,
  message JSONB NOT NULL,
  published BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  CONSTRAINT messages_message_key UNIQUE (message)
);

-- name: down
DROP TABLE IF EXISTS messages;
