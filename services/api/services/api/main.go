package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"take-home-challenge/configuration"
	"take-home-challenge/middleware"
	"take-home-challenge/services/api/controllers"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l := slog.Default()

	conf, err := configuration.Load()
	if err != nil {
		l.ErrorContext(ctx, "Failed to load service configuration.",
			"error", err,
		)
		panic(fmt.Sprint("Failed to load service configuration."))
	}

	r := router(conf)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       120 * time.Second,
		WriteTimeout:      240 * time.Second,
		ReadHeaderTimeout: 120 * time.Second,
		IdleTimeout:       240 * time.Second,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
	}
	l.InfoContext(ctx, "Starting server.",
		"addr", server.Addr,
	)
	err = server.ListenAndServe()
	if err != nil {
		l.ErrorContext(ctx, "Failed to listen and serve.",
			"error", err,
		)
		panic(fmt.Sprint("Failed to listen and serve."))
	}
}

func router(conf configuration.Config) chi.Router {
	r := chi.NewRouter()

	r.Route("/api/v2", func(r chi.Router) {
		r.Use(
			middleware.NewLogger(slog.Default()),
		)

		r.Route("/users", func(r chi.Router) {
			r.Mount("/", controllers.NewUsersController(conf).Routes())
		})
		r.Route("/messages", func(r chi.Router) {
			r.Mount("/", controllers.NewMessagesController(conf).Routes())
		})
	})

	return r
}
