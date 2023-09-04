package main

import (
	"database/sql"

	"cirello.io/pglock"
	"github.com/bwmarrin/discordgo"
	"github.com/icco/relay/database/model"
	"go.uber.org/zap"
)

// messageRecieve will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageRecieve(db *sql.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		c, err := model.GetLockClient(db)
		if err != nil {
			log.Errorw("getting lock client", zap.Error(err))
			return
		}

		l, err := c.Acquire(m.Message.ID, pglock.FailIfLocked(), pglock.WithData([]byte("sent")))
		if err != nil {
			log.Errorw("getting lock", zap.Error(err))
			return
		}
		defer l.Close()

		// If the message is "ping" reply with "Pong!"
		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}

		// If the message is "pong" reply with "Ping!"
		if m.Content == "pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}
	}
}
