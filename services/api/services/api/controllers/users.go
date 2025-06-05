package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"take-home-challenge/configuration"
	"take-home-challenge/coordinators"
	"take-home-challenge/middleware"
	"take-home-challenge/models"
)

type UsersController struct {
	coordinator coordinators.UsersCoordinator
}

func NewUsersController(conf configuration.Config) *UsersController {
	return &UsersController{
		coordinator: conf.UsersCoordinator,
	}
}

func (u *UsersController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/", u.create)
	r.Get("/{userID}", u.read)

	return r
}

func (u *UsersController) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := ctx.Value(middleware.KeyLogger).(*slog.Logger)

	res := Response{}

	data := &models.User{}
	if err := render.Bind(r, data); err != nil {
		res.BadRequest(w, r, err.Error())
		return
	}

	if err := u.coordinator.Create(data); err != nil {
		if coordinators.IsValidationError(err) {
			res.BadRequest(w, r, err.Error())
			return
		}

		l.InfoContext(ctx, "Failed to create user.",
			"error", err,
		)
		res.InternalServerError(w, r, err.Error())
		return
	}

	res.JSON(w, r, nil, http.StatusCreated)
}

func (u *UsersController) read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := ctx.Value(middleware.KeyLogger).(*slog.Logger)

	res := Response{}

	userID := chi.URLParam(r, "userID")

	id, _ := strconv.Atoi(userID)
	user, err := u.coordinator.Read(id)
	if err != nil {
		if coordinators.IsRecordNotFoundError(err) {
			res.NotFound(w, r, err.Error())
			return
		} else if err != nil {
			l.InfoContext(ctx, "Failed to read user.",
				"error", err,
			)
			res.InternalServerError(w, r, err.Error())
			return
		}
	}

	res.JSON(w, r, user, http.StatusOK)
}
