package main

import (
	"log"
	"net/http"

	"github.com/goflash/flash/v2"
)

func main() {
	app := flash.New()

	app.SetErrorHandler(func(c flash.Ctx, err error) {
		_ = c.String(http.StatusTeapot, "custom onerror: "+err.Error())
	})

	app.SetNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found here"))
	}))

	app.GET("/err", func(c flash.Ctx) error { return http.ErrBodyNotAllowed })

	log.Fatal(http.ListenAndServe(":8080", app))
}
