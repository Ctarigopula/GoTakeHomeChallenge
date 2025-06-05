package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

func NewLogger(l *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := l.With(
				"httpMethod", r.Method,
				"userAgent", r.UserAgent(),
			)
			ctx = context.WithValue(ctx, KeyLogger, logger)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
