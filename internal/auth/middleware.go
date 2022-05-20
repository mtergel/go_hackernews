package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/tergelm/go_hackernews/internal/pkg/jwt"
	"github.com/tergelm/go_hackernews/internal/users"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if authToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate jwt
			tokenStr := authToken
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// Create user object
			user := users.User{Username: username}

			// Check password
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Add user to context
			user.Id = strconv.Itoa(id)
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Used in middleware
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)

	return raw
}
