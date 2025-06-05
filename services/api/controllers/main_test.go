package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"

	"take-home-challenge/configuration"
	"take-home-challenge/middleware"
)

var (
	conf   configuration.Config
	server *httptest.Server
)

func TestMain(m *testing.M) {
	os.Setenv("DB_DSN", "postgres://postgres:password@localhost:5432")

	conf, _ = configuration.Load()

	r := chi.NewRouter()
	r.Route("/api/v2", func(r chi.Router) {
		r.Use(
			middleware.NewLogger(slog.Default()),
		)

		r.Route("/users", func(r chi.Router) {
			r.Mount("/", NewUsersController(conf).Routes())
		})
		r.Route("/messages", func(r chi.Router) {
			r.Mount("/", NewMessagesController(conf).Routes())
		})
	})

	conf.DB.Exec("TRUNCATE TABLE clients;")
	conf.DB.Exec("TRUNCATE TABLE client_messages;")
	conf.DB.Exec("TRUNCATE TABLE client_users;")

	server = httptest.NewServer(r)

	os.Exit(m.Run())
}

func fromBody[T any](res *http.Response) T {
	var t T
	json.NewDecoder(res.Body).Decode(&t)
	return t
}

func toBody(t any) io.Reader {
	b, _ := json.Marshal(t)
	return bytes.NewReader(b)
}
