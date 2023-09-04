package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/alecthomas/units"
	"github.com/bwmarrin/discordgo"
	"github.com/icco/relay/lib"
	"go.uber.org/zap"
)

const ()

func hookHandler(dg *discordgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
