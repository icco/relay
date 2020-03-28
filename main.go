package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

const (
	permissions = 51264
	channelID   = "#ops"
)

var log = InitLogging()

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalf("discord token is empty")
	}

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting up on http://localhost:%s", port)

	if os.Getenv("ENABLE_STACKDRIVER") != "" {
		labels := &stackdriver.Labels{}
		labels.Set("app", "relay", "The name of the current app.")
		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID:               "icco-cloud",
			MonitoredResource:       monitoredresource.Autodetect(),
			DefaultMonitoringLabels: labels,
			DefaultTraceAttributes:  map[string]interface{}{"app": "cron"},
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
		log.Fatalf("error creating Discord session: %w", err)
	}
	defer dg.Close()

	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi."))
	})

	r.Post("/hook", func(w http.ResponseWriter, r *http.Request) {
		var message []byte
		if _, err := io.ReadFull(r.Body, message); err != nil {
			log.Printf("could not read body: %w", err)
			http.Error(w, err.Error(), 500)
			return
		}

		if err := messageCreate(dg, string(message)); err != nil {
			log.Printf("could not send message: %w", err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("."))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func messageCreate(s *discordgo.Session, m string) error {
	_, err := s.ChannelMessageSend(channelID, "Pong!")
	return err
}
