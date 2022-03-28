package auth

import (
	"context"
	"movie_graphql_be/pkg/jwt"
	"net/http"
	"os"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func JwtMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			secretKey := []byte(os.Getenv("SECRET_KEY"))

			if tokenString == "" {
				next.ServeHTTP(w, r)
				return
			}

			payload, err := jwt.ParseToken(tokenString, secretKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			userData := jwt.Payload{
				ID:        payload.ID,
				Username:  payload.Username,
				Email:     payload.Email,
				FirstName: payload.FirstName,
				LastName:  payload.LastName,
			}

			ctx := context.WithValue(r.Context(), userCtxKey, &userData)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
//? to check authorization in the resolver
func ForContext(ctx context.Context) *jwt.Payload {
	raw, _ := ctx.Value(userCtxKey).(*jwt.Payload)
	return raw
}