package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
)

type contextKey string

const userContextKey = contextKey("user")

func ContextSetUser(r *http.Request, user *sqlc.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func ContextGetUser(r *http.Request) *sqlc.User {
	user, ok := r.Context().Value(userContextKey).(*sqlc.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}

func ContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}
