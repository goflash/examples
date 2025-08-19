package main

import (
	"log"
	"net/http"

	"github.com/goflash/flash/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mux", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("mux handler"))
	})

	app := flash.New()
	app.Mount("/mux", mux)
	app.GET("/app", func(c flash.Ctx) error { return c.String(http.StatusOK, "app handler") })

	log.Fatal(http.ListenAndServe(":8080", app))
}
