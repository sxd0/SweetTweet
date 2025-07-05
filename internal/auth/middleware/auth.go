package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/sxd0/SweetTweet/pkg/jwt"
)

type contextKey string

const UserIDKey contextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := jwt.ParseToken(token)
		if err != nil {
			log.Println("JWT parsing error:", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		log.Printf("Parsed JWT claims: %+v\n", claims)

		ctx := context.WithValue(r.Context(), UserIDKey, claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
