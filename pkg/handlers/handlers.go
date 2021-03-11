package handlers

import (
	"net/http"

	"github.com/petrostrak/booking-with-go/pkg/config"
	"github.com/petrostrak/booking-with-go/pkg/models"
	"github.com/petrostrak/booking-with-go/pkg/render"
)

var (
	// Repo used by the handlers
	Repo *Repository
)

// Repository struct
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repositoty
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handler will return a simple msg to Writer
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About handler will return the info regarding the site
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "hello, again"

	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
