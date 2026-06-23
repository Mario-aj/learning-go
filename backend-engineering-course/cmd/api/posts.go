package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mario-aj/social/internal/store"
)

const postCtxKey = "post"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

type CreatePostCommentPayload struct {
	Content string `json:"content" validate:"required,max=300"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,

		// TODO: change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	post.Comments = comments

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postIdParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(postIdParam, 10, 64)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	err = app.store.Posts.Delete(ctx, id)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundErrorResponse(w, r, err)
		default:
			app.internalServerErrorResponse(w, r, err)
		}

		return
	}

	err = app.store.Comments.DeleteByPostID(ctx, id)

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload UpdatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) createPostCommentHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload CreatePostCommentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
	}

	comment := &store.Comment{
		Content: payload.Content,
		PostID:  post.ID,

		// TODO: Change it once the auth is completed.
		UserID: 1,
	}

	if err := app.store.Comments.Create(r.Context(), comment); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}

	if err := writeJSON(w, http.StatusCreated, comment); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) postsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			app.internalServerErrorResponse(w, r, err)
			return
		}

		ctx := r.Context()

		post, err := app.store.Posts.GetById(ctx, id)

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

		ctx = context.WithValue(ctx, postCtxKey, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromCtx(r *http.Request) *store.Post {
	post, _ := r.Context().Value(postCtxKey).(*store.Post)

	return post
}
