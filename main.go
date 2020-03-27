package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi"
)

const (
	permissions = 51264
	channelID   = "#ops"
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
