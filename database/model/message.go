package model

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/icco/gutil/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	log = logging.Must(logging.NewLogger("relay"))
)

// OpenDatabase opens a connection to the database using the DATABASE_URL
// environment variable.  It returns a pointer to a gorm.DB instance and an
// error if the DATABASE_URL is empty or if there was an error opening the
// connection.
func OpenDatabase(db *sql.DB) (*gorm.DB, error) {
	opts := postgres.Config{}
	if db != nil {
		opts.Conn = db
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is empty")
	} else if opts.Conn == nil {
		opts.DSN = dsn
	}
	return gorm.Open(postgres.New(opts), &gorm.Config{})
}

// Create inserts a new message into the database.  It returns an error if the
// database connection fails or if the operation fails.
func (m *Message) Create(ctx context.Context, db *sql.DB) error {
	dbo, err := OpenDatabase(db)
	if err != nil {
		return err
	}

	result := dbo.WithContext(ctx).Where(Message{ID: m.ID}).FirstOrCreate(m)

	return result.Error
}

func (m *Message) Send(ctx context.Context, dg *discordgo.Session, db *sql.DB) error {
	if err := messageCreate(dg, m.Content); err != nil {
		return err
	}
	log.Debugw("message sent", "message", m)

	dbo, err := OpenDatabase(db)
	if err != nil {
		return err
	}

	m.Published = true
	return dbo.WithContext(ctx).Save(m).Error
}

func messageCreate(s *discordgo.Session, m string) error {
	chnl, err := fetchPrimaryTextChannelID(s)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(chnl, m)
	return err
}

func fetchPrimaryTextChannelID(sess *discordgo.Session) (string, error) {
	var channelid string
	guilds, err := sess.UserGuilds(1, "", "")
	if err != nil {
		return "", err
	}

	guild, err := sess.Guild(guilds[0].ID)
	if err != nil {
		return "", err
	}

	channels, err := sess.GuildChannels(guild.ID)
	if err != nil {
		return "", err
	}

	for _, channel := range channels {
		channel, err = sess.Channel(channel.ID)
		if err != nil {
			return "", err
		}
		if channel.Type == discordgo.ChannelTypeGuildText {
			channelid = channel.ID
			break
		}
	}

	if channelid == "" {
		return "", fmt.Errorf("no primary channel found")
	}

	return channelid, nil
}
