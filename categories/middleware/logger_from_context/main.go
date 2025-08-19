package main

import (
	"log"
	"net/http"
	"os"

	"log/slog"

	logctx "github.com/goflash/flash/logctx"
	"github.com/goflash/flash/v2"
	mw "github.com/goflash/flash/v2/middleware"
)

func main() {
	app := flash.New()

	// Configure slog and have the app inject it into each request context
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.SetLogger(logger)

	// Optional: emit structured access logs per request
	app.Use(mw.Logger())

	app.GET("/", func(c flash.Ctx) error {
		// Retrieve the logger from the request context
		l := logctx.LoggerFromContext(c.Context())
		l.Info("handling request", "path", c.Path())
		return c.String(http.StatusOK, "logged via context!")
	})

	log.Fatal(http.ListenAndServe(":8080", app))
}
