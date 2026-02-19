package middleware

import "net/http"

var ErrUnauthorized = "Authorization is missing or invalid"

func EnsureAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		if t != "" { // TODO: Implement actual token validation
			http.Error(w, ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
