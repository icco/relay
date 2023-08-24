package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"cirello.io/pglock"
	"github.com/alecthomas/units"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/icco/gutil/logging"
	"github.com/icco/gutil/otel"
	"github.com/icco/relay/lib"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	permissions = 51264
	channelID   = "ops"
	project     = "relay"
	gcpID       = "icco-cloud"
)

var (
	log = logging.Must(logging.NewLogger(project))
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN is empty")
	}

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Infow("Starting up", "host", fmt.Sprintf("http://localhost:%s", port))

	ctx := context.Background()
	if err := otel.Init(ctx, log, gcpID, project); err != nil {
		log.Errorw("could not init opentelemetry", zap.Error(err))
	}

	dbUser := ("DB_USER")                         // e.g. 'my-db-user'
	dbPwd := os.Getenv("DB_PASS")                 // e.g. 'my-db-password'
	unixSocketPath := os.Getenv("DB_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
	dbName := os.Getenv("DB_NAME")                // e.g. 'my-database'

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatalw("cannot connect to database server", zap.Error(err))
	}
	defer db.Close()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalw("error creating Discord session", zap.Error(err))
	}
	dg.AddHandler(messageRecieve(db))

	if err := dg.Open(); err != nil {
		log.Fatalw("error opening discord connection", zap.Error(err))
	}
	defer dg.Close()

	r := chi.NewRouter()
	r.Use(otel.Middleware)
	r.Use(middleware.RealIP)
	r.Use(logging.Middleware(log.Desugar(), gcpID))

	crs := cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:     []string{"Link"},
		MaxAge:             300, // Maximum value not ignored by any of major browsers
	})
	r.Use(crs.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi."))
	})

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi."))
	})

	r.Post("/hook", hookHandler(dg))

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func hookHandler(dg *discordgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("content-type")
		log.Debugw("got content-type", "content-type", ct)

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorw("could not read buffer", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr2
		defer r.Body.Close()

		var msg string
		var parseMethod string
		switch ct {
		case "application/json":
			parseMethod = "json"
			msg = lib.BufferToMessage(buf)
		case "plain/text":
			parseMethod = "plaintext"
			msg = string(buf)
		case "application/x-www-form-urlencoded":
			parseMethod = "urlencoded"
			if err := r.ParseForm(); err != nil {
				log.Errorw("could not parse form", "body", string(buf), zap.Error(err))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			values := r.Form
			str, err := json.Marshal(values)
			if err != nil {
				log.Errorw("could not marshal values", "body", string(buf), "values", values, zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			msg = lib.BufferToMessage(str)
		case "":
			http.Error(w, "empty content type recieved", http.StatusBadRequest)
			return
		default:
			parseMethod = "form"
			if err := r.ParseMultipartForm(int64(20 * units.Megabyte)); err != nil {
				log.Errorw("could not parse multipart form", "body", string(buf), zap.Error(err))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			parts := strings.Split(ct, ";")
			log.Debugw("parsing form", "parts", parts)
			if len(parts) >= 1 && parts[0] == "multipart/form-data" {
				val := r.FormValue("payload")
				log.Debugw("attempting form parse", "payload", val)
				msg = lib.BufferToMessage([]byte(val))
			}
		}

		// Generate a 200 but sends no message
		if msg == "" {
			log.Infow("empty message generated", "body", string(buf), "parseMethod", parseMethod, "contentType", ct)
			w.Write([]byte(""))
			return
		}

		if dg == nil {
			log.Errorw("discord session is nil")
			w.Write([]byte(""))
			return
		}

		if err := messageCreate(dg, msg); err != nil {
			log.Errorw("could not send message", zap.Error(err))
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte(""))
		return
	}
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

// messageRecieve will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageRecieve(db *sql.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		c, err := lib.GetLockClient(db)
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
