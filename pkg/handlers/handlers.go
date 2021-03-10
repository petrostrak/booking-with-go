package handlers

import (
	"net/http"

	"github.com/petrostrak/booking-with-go/pkg/config"
	"github.com/petrostrak/booking-with-go/pkg/render"
)

var (
	// Repo used by the handlers
	Repo *Repository
)

// Repository struct
type Repository struct {
	AppC *config.AppConfig
}

// NewRepo creates a new repositoty
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		AppC: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handler will return a simple msg to Writer
func (Rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl")
}

// About handler will return the info regarding the site
func (Rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}
