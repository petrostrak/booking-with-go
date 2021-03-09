package handlers

import (
	"net/http"

	"github.com/petrostrak/booking-with-go/pkg/render"
)

// Home handler will return a simple msg to Writer
func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl")
}

// About handler will return the info regarding the site
func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}
