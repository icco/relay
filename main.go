package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"cirello.io/pglock"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/alecthomas/units"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/icco/gutil/logging"
	"github.com/icco/relay/lib"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
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

	if os.Getenv("ENABLE_STACKDRIVER") != "" {
		labels := &stackdriver.Labels{}
		labels.Set("app", project, "The name of the current app.")
		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID:               gcpID,
			MonitoredResource:       monitoredresource.Autodetect(),
			DefaultMonitoringLabels: labels,
			DefaultTraceAttributes:  map[string]interface{}{"app": project},
		})

		if err != nil {
			log.Fatalw("failed to create the stackdriver exporter", zap.Error(err))
		}
		defer sd.Flush()

		view.RegisterExporter(sd)
		trace.RegisterExporter(sd)
		trace.ApplyConfig(trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		})
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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
		log.Fatalw("error opening connection", zap.Error(err))
	}
	defer dg.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(logging.Middleware(log, gcpID))

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

	r.Post("/hook", func(w http.ResponseWriter, r *http.Request) {
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
		switch ct {
		case "application/json":
			msg = lib.BufferToMessage(buf)
		case "plain/text":
			msg = string(buf)
		default:
			err := r.ParseMultipartForm(int64(20 * units.Megabyte))
			if err != nil {
				log.Errorw("could not parse form", "body", string(buf), zap.Error(err))
				http.Error(w, err.Error(), 500)
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

		// Generates a 200, but log an error, because this shouldn't happen.
		if msg == "" {
			log.Errorw("empty message generated", "body", string(buf))
			w.Write([]byte(""))
			return
		}

		if err := messageCreate(dg, msg); err != nil {
			log.Errorw("could not send message", zap.Error(err))
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte(""))
	})

	h := &ochttp.Handler{
		Handler:     r,
		Propagation: &propagation.HTTPFormat{},
	}
	if err := view.Register([]*view.View{
		ochttp.ServerRequestCountView,
		ochttp.ServerResponseCountByStatusCode,
	}...); err != nil {
		log.Fatalw("Failed to register ochttp views", zap.Error(err))
	}

	log.Fatal(http.ListenAndServe(":"+port, h))
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
		log.Infow("recieved message", "session", s, "message", m)

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
