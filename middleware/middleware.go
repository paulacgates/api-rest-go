package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
	})
}
