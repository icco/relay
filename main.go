package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/alecthomas/units"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/icco/relay/lib"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

const (
	permissions = 51264
	channelID   = "ops"
)

var (
	log = lib.InitLogging()
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalf("DISCORD_TOKEN is empty")
	}

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Infof("Starting up on http://localhost:%s", port)

	if os.Getenv("ENABLE_STACKDRIVER") != "" {
		labels := &stackdriver.Labels{}
		labels.Set("app", "relay", "The name of the current app.")
		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID:               "icco-cloud",
			MonitoredResource:       monitoredresource.Autodetect(),
			DefaultMonitoringLabels: labels,
			DefaultTraceAttributes:  map[string]interface{}{"app": "relay"},
		})

		if err != nil {
			log.WithError(err).Fatalf("failed to create the stackdriver exporter")
		}
		defer sd.Flush()

		view.RegisterExporter(sd)
		trace.RegisterExporter(sd)
		trace.ApplyConfig(trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		})
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.WithError(err).Fatal("error creating Discord session")
	}

	if err := dg.Open(); err != nil {
		log.WithError(err).Fatal("error opening connection")
	}
	defer dg.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(lib.LoggingMiddleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi."))
	})

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi."))
	})

	r.Post("/hook", func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("content-type")
		log.WithField("content-type", ct).Debug("got content-type")

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Error("could not read buffer")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr2
		defer r.Body.Close()

		var msg string
		switch ct {
		case "application/json":
			msg, err = lib.BufferToMessage(buf)
			if err != nil {
				log.WithError(err).WithField("body", string(buf)).Error("could not decode body")
				http.Error(w, err.Error(), 500)
				return
			}
		case "plain/text":
			msg = string(buf)
		default:
			err := r.ParseMultipartForm(int64(20 * units.Megabyte))
			if err != nil {
				log.WithError(err).WithField("body", string(buf)).Error("could not parse form")
				http.Error(w, err.Error(), 500)
				return
			}
			parts := strings.Split(ct, ";")
			log.WithField("parts", parts).Debug("parsing form")
			if len(parts) >= 1 && parts[0] == "multipart/form-data" {
				val := r.FormValue("payload")
				log.WithField("payload", val).Debug("attempting form parse")
				msg, err = lib.BufferToMessage([]byte(val))
				if err != nil {
					log.WithError(err).WithField("body", string(buf)).Error("could not decode body")
					http.Error(w, err.Error(), 500)
					return
				}
			}
		}

		// Generates a 200, but log an error, because this shouldn't happen.
		if msg == "" {
			log.WithField("body", string(buf)).Error("empty message generated")
			w.Write([]byte(""))
			return
		}

		if err := messageCreate(dg, msg); err != nil {
			log.WithError(err).Error("could not send message")
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
		log.WithError(err).Fatal("Failed to register ochttp views")
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
