package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/alecthomas/units"
	"github.com/bwmarrin/discordgo"
	"github.com/icco/relay/database/model"
	"github.com/icco/relay/parse"
	"go.uber.org/zap"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	out := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(out); err != nil {
		log.Errorw("could not encode json", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func hookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ct := r.Header.Get("content-type")
		log.Debugw("got content-type", "content-type", ct)

		buf, err := io.ReadAll(r.Body)
		if err != nil {
			log.Errorw("could not read buffer", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr2
		defer r.Body.Close()

		var msg string
		var parseMethod string
		switch ct {
		case "application/json":
			parseMethod = "json"
			msg = parse.BufferToMessage(buf)
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
			msg = parse.BufferToMessage(str)
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
				msg = parse.BufferToMessage([]byte(val))
			}
		}

		m := &model.Message{
			Content:   msg,
			Message:   []byte("{}"),
			Published: false,
		}

		if err := m.Create(ctx, db); err != nil {
			log.Errorw("could not save message", "parse_method", parseMethod, zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(""))
	}
}

func cronHandler(dg *discordgo.Session, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if dg == nil {
			log.Errorw("discord session is nil")
			w.Write([]byte(""))
			return
		}

		dbo, err := model.OpenDatabase(db)
		if err != nil {
			log.Errorw("could not open database", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var messages []model.Message
		dbo.WithContext(ctx).Where("published = false").Find(&messages)
		for _, msg := range messages {
			if err := msg.Send(ctx, dg, db); err != nil {
				log.Errorw("could not send message", zap.Error(err))
				http.Error(w, err.Error(), 500)
				return
			}
		}
	}
}
