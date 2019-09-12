package middleware

import (
	"log"
	"net/http"
)

type (
	BasicMiddlewareInterface interface {
		Middleware(next http.Handler) http.Handler
	}

	AuthenticationMiddleware struct {
		tokenUsers map[string]string
	}
)

func New() BasicMiddlewareInterface {
	authMiddleware := &AuthenticationMiddleware{
		tokenUsers: map[string]string{},

	}
	authMiddleware.tokenUsers["00000000"] = "user0"
	authMiddleware.tokenUsers["aaaaaaaa"] = "userA"
	authMiddleware.tokenUsers["05f717e5"] = "randomUser"
	authMiddleware.tokenUsers["deadbeef"] = "user0"

	return authMiddleware
}


// Middleware function, which will be called for each request
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
