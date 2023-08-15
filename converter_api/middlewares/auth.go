package middlewares

import (
	"fmt"
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := os.Getenv("SECRET_TOKEN"); err == "" {
			_ = fmt.Errorf("the SECRET_TOKEN not found on environment variables")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		secretToken := os.Getenv("SECRET_TOKEN")
		token := r.Header.Get("Authorization")

		if token != secretToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
