-- Auto-generated at Mon, 04 Sep 2023 21:44:14 UTC
-- Please do not change the name attributes

-- name: up
ALTER TABLE `messages` ADD `content` TEXT NOT NULL DEFAULT '' ;

-- name: down
ALTER TABLE `messages` DROP `content` ;

