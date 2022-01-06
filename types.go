package sentry

import (
	"net/http"
)

// SentryOptions defines the parameters for running a sentry instance
type SentryOptions struct {
	Port        int
	Secret      string
	LogFilePath string
}

// SentryInstance defines the model for a sentry instance
type SentryInstance struct {
	opts   *SentryOptions
	server *http.Server
}
