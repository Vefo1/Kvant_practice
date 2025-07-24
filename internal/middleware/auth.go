package middleware

import (
	"net/http"

	"github.com/Vefo1/Kvant_practice/pkg/logger"
)

// AuthMiddleware checks for a specific authorization token in the request header
func AuthMiddleware(authToken, headerName string, log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from the specified header
			providedToken := r.Header.Get(headerName)

			// Log the provided token for debugging (remove in production if sensitive)
			log.Debug("AuthMiddleware: Received token in header '%s': %s", headerName, providedToken)

			// Check if the provided token matches the expected token
			if providedToken != authToken {
				log.Warn("AuthMiddleware: Unauthorized access attempt from %s. Invalid token provided.", r.RemoteAddr)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// If token is valid, proceed to the next handler
			log.Debug("AuthMiddleware: Token is valid. Proceeding to next handler.")
			next.ServeHTTP(w, r)
		})
	}
}
