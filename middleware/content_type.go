package middleware

import (
	"net/http"
	"slices"
	"strings"
)

var (
	AllowedContentTypes = []string{
		"application/json",
	}

	ErrWrongContentType = "Content-Type is not supported"
)

func EnsureContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Content-Type")
		if slices.Contains(AllowedContentTypes, strings.ToLower(t)) {
			http.Error(w, ErrWrongContentType, http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}
