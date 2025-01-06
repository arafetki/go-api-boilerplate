package utils

import (
	"context"
	"net/http"
)

type DummyUser struct {
	ID int32
}

var AnonymousUser = &DummyUser{}

func (u *DummyUser) IsAnonymous() bool {
	return u == AnonymousUser
}

type contextKey string

const userContextKey = contextKey("user")

func ContextSetUser(r *http.Request, user *DummyUser) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func ContextGetUser(r *http.Request) *DummyUser {
	user, ok := r.Context().Value(userContextKey).(*DummyUser)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
