package sentry

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// NewInstance returns a new sentry instance
func NewInstance(opts *SentryOptions) *SentryInstance {

	secret = opts.Secret
	logFilePath = opts.LogFilePath
	port = opts.Port

	if secret == "" {
		panic(errors.New("sentry secret cannot be empty"))
	}

	if len(secret) < 8 {
		panic(errors.New("sentry secret should be at least 8 characters"))
	}

	_, err := os.Stat(logFilePath)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(authMiddleware)
	r.HandleFunc("/sentry", wsHandler)
	return &SentryInstance{
		opts: opts,
		server: &http.Server{
			Addr:    "127.0.0.1:" + fmt.Sprint(port),
			Handler: r,
		},
	}
}
