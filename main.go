package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/icco/gutil/logging"
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

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalw("cannot connect to database server", zap.Error(err))
	}
	db.SetMaxOpenConns(5)
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

	r.Get("/", healthcheckHandler)
	r.Get("/healthz", healthcheckHandler)
	r.Get("/healthcheck", healthcheckHandler)

	r.Post("/hook", hookHandler(db))

	r.Post("/cron", cronHandler(dg, db))

	log.Fatal(http.ListenAndServe(":"+port, r))
}
