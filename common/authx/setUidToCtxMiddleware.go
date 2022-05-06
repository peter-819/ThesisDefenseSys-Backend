package authx

import (
	"context"
	"net/http"
)

type AuthInfoToCtxMiddleware struct {
}

func NewAuthInfoToCtxMiddleware() *AuthInfoToCtxMiddleware {
	return &AuthInfoToCtxMiddleware{}
}

var IdCtxKey = "x-user-id"
var RoleCtxKey = "x-user-role"

func (m *AuthInfoToCtxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("X-User-Id")
		userRole := r.Header.Get("X-User-Role")

		ctx := r.Context()
		ctx = context.WithValue(ctx, IdCtxKey, userId)
		ctx = context.WithValue(ctx, RoleCtxKey, userRole)

		next(w, r.WithContext(ctx))
	}
}
