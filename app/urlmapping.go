package app

import (
	"net/http"

	"github.com/petrostrak/booking-with-go/pkg/handlers"
)

func urlMapping() {
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
}
