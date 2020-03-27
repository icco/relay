package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi"
)

const (
	permissions = 51264
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalf("discord token is empty")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Falalf("error creating Discord session: %w", err)
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

		if err := messageCreate(s, message); err != nil {
			log.Printf("could not send message: %w", err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("."))
	})
	http.ListenAndServe(fmt.Spritnf(":%s", port), r)
}

func messageCreate(s *discordgo.Session, m string) error {

	s.ChannelMessageSend(m.ChannelID, "Pong!")
}
