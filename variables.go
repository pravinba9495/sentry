package sentry

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	secret      = ""
	logFilePath = ""
	port        = 8899
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}
)
