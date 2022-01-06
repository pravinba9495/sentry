package sentry

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	secret      = ""
	logFilePath = ""
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}
)

// Run starts a sentry instance
func (instance *SentryInstance) Run() chan error {
	err := make(chan error, 1)
	go func() {
		err <- instance.server.ListenAndServe()
	}()
	return err
}

// NewInstance returns a new sentry instance
func NewInstance(opts *SentryOptions) *SentryInstance {

	secret = opts.Secret
	logFilePath = opts.LogFilePath

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(authMiddleware)
	r.HandleFunc("/sentry", wsHandler)
	return &SentryInstance{
		opts: opts,
		server: &http.Server{
			Addr:    "127.0.0.1:" + fmt.Sprint(opts.Port),
			Handler: r,
		},
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	cmd := exec.Command("tail", "-f", logFilePath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	} else {
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return
		} else {
			mw := io.MultiReader(stdout, stderr)
			if err := cmd.Start(); err != nil {
				return
			} else {
				scanner := bufio.NewScanner(mw)
				for scanner.Scan() {
					txt := scanner.Text()
					if err := c.WriteMessage(websocket.TextMessage, []byte(txt)); err != nil {
						return
					}
				}
				return
			}
		}
	}
}

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
