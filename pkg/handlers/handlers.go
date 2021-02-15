package handlers

import (
	"net/http"

	"github.com/NathanielRand/go-bnb/pkg/config"
	"github.com/NathanielRand/go-bnb/pkg/models"
	"github.com/NathanielRand/go-bnb/pkg/render"
)

// Repo is the repository used by the handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home function is the handler for the home page.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About function is the handler for the about page.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact function is the handler for the contact page.
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

// Reservation function is the handler for the about page.
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "reservation.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Tidal function is the handler for the tidal page.
func (m *Repository) Tidal(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "tidal.page.html", &models.TemplateData{})
}

// Cliffside function is the handler for the Cliffside page.
func (m *Repository) Cliffside(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "cliffside.page.html", &models.TemplateData{})
}

// Oceanspray function is the handler for the Oceanspray page.
func (m *Repository) Oceanspray(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "oceanspray.page.html", &models.TemplateData{})
}

// Seabreeze function is the handler for the Seabreeze page.
func (m *Repository) Seabreeze(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "seabreeze.page.html", &models.TemplateData{})
}

// Dawn function is the handler for the Dawn page.
func (m *Repository) Dawn(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "dawn.page.html", &models.TemplateData{})
}

// Sunset function is the handler for the Sunset page.
func (m *Repository) Sunset(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "sunset.page.html", &models.TemplateData{})
}
