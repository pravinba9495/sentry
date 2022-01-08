package sentry

import "net/http"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		if params.Has("token") && params.Get("token") == secret {
			w.Header().Add("Access-Control-Allow-Origin", r.Host)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
