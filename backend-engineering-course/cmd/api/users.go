package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mario-aj/social/internal/store"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByID(ctx, userID)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundErrorResponse(w, r, err)
			return
		default:
			app.internalServerErrorResponse(w, r, err)
			return
		}
	}

	if err := writeJSON(w, http.StatusOK, user); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}
