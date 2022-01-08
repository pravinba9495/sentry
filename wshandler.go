package sentry

import (
	"bufio"
	"io"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

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
