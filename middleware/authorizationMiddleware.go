package middleware

import (
	"net/http"
	"stratus-api/jwt"
	"strings"

	"github.com/gorilla/mux"
)

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) !=2 || tokenParts[0]!="Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]
		
		vars := mux.Vars(r)
		username := vars["username"]

		check, err := jwt.AuthToken(username, token)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if !check {
			http.Error(w, "Invalid username", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}