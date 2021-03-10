package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/petrostrak/booking-with-go/pkg/config"
	"github.com/petrostrak/booking-with-go/pkg/handlers"
)

// Routes func
func routes(appC *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
