# sentry
[![Go Reference](https://pkg.go.dev/badge/github.com/pravinba9495/sentry.svg)](https://pkg.go.dev/github.com/pravinba9495/sentry) ![Go Report Card](https://goreportcard.com/badge/github.com/pravinba9495/sentry)

Stream logs through websockets, written in Go.

## Usage
### Server
```go
package main

import (
    "github.com/pravinba9495/sentry"
)

func main() {

    // Initialize sentry options 
    opts := &sentry.SentryOptions{
	Port:        8080,
	Secret:      "SOME_RANDOM_SECRET",
	LogFilePath: "logfile.txt",
    }

    // Create and start the sentry instance
    instance := sentry.NewInstance(opts)
    err := <-instance.Run()
    
    log.Println(err)
}
```

## License
MIT
