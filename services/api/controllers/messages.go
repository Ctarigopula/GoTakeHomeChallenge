package controllers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"take-home-challenge/configuration"
	"take-home-challenge/coordinators"
	"take-home-challenge/middleware"
	"take-home-challenge/models"
	"take-home-challenge/services/api/controllers/payloads"
)

const (
	ErrInvalidPayload = "Invalid request payload. Please check the JSON format and required fields."
	ErrInvalidMessage = "Invalid message data. Please ensure all fields are correct."
)

type MessagesController struct {
	coordinator coordinators.MessagesCoordinator
}

// NewMessagesController creates and returns a new MessagesController with the given configuration.
func NewMessagesController(conf configuration.Config) *MessagesController {
	return &MessagesController{
		coordinator: conf.MessagesCoordinator,
	}
}

// getLogger retrieves the logger from the context or returns the default logger if none is found.
func getLogger(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(middleware.KeyLogger).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}

// Routes defines message-related HTTP endpoints with JSON responses.
func (m *MessagesController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/", m.create)
	r.Post("/delete", m.delete)
	r.Get("/{messageID}", m.read)

	return r
}

// create handles HTTP requests to create a new message.Validates the request payload,
// and responds with appropriate HTTP status codes and error messages.
func (m *MessagesController) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := getLogger(ctx)
	res := Response{}

	message := &models.Message{}
	if err := render.Bind(r, message); err != nil {
		logger.WarnContext(ctx, "Invalid request body", "error", err)
		res.BadRequest(w, r, ErrInvalidPayload)
		return
	}

	// Create the message
	if err := m.coordinator.Create(message); err != nil {
		switch {
		case coordinators.IsValidationError(err):
			logger.WarnContext(ctx, "Validation failed", "error", err)
			res.BadRequest(w, r, ErrInvalidMessage)
		case errors.Is(err, context.DeadlineExceeded):
			res.InternalServerError(w, r, "Request timed out")
		default:
			logger.ErrorContext(ctx, "Failed to create message", "error", err)
			res.InternalServerError(w, r, "Could not create message")
		}
		return
	}

	res.JSON(w, r, nil, http.StatusCreated)
}

// delete handles message deletion requests, validates the input and marks the given messages as deleted.
func (m *MessagesController) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := getLogger(ctx)
	res := Response{}

	data := &payloads.MessagesMarkDeleted{}
	if err := render.Bind(r, data); err != nil {
		logger.WarnContext(ctx, "Invalid delete request", "error", err)
		res.BadRequest(w, r, ErrInvalidPayload)
		return
	}

	if err := m.coordinator.MarkDeleted(data.DeletedWhen, data.IDs); err != nil {
		if coordinators.IsValidationError(err) {
			logger.WarnContext(ctx, "Validation failed", "error", err)
			res.BadRequest(w, r, ErrInvalidMessage)
			return
		}
		logger.InfoContext(ctx, "Failed to mark messages as deleted.", "error", err)
		res.InternalServerError(w, r, "Could not mark messages as deleted.")
		return
	}

	res.JSON(w, r, nil, http.StatusOK)
}

// read handles GET /messages/{messageID} requests. It validates the message ID, fetches the message from the database,
// and returns the result or appropriate error response.
func (m *MessagesController) read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := getLogger(ctx)
	res := Response{}

	// Extract and validate messageID from URL
	messageID := chi.URLParam(r, "messageID")
	id, err := strconv.Atoi(messageID)
	if err != nil {
		res.BadRequest(w, r, "Invalid message ID")
		return
	}

	// Retrieve the message
	msg, err := m.coordinator.Read(id)
	if err != nil {
		switch {
		case coordinators.IsRecordNotFoundError(err):
			res.NotFound(w, r, "Message not found")
		case errors.Is(err, context.DeadlineExceeded):
			res.InternalServerError(w, r, "Request timed out")
		default:
			logger.ErrorContext(r.Context(), "Failed to read message", "error", err)
			res.InternalServerError(w, r, "Could not read message")
		}
		return
	}

	res.JSON(w, r, msg, http.StatusOK)
}
