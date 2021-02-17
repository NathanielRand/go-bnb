package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NathanielRand/go-bnb/internal/config"
	"github.com/NathanielRand/go-bnb/internal/models"
	"github.com/NathanielRand/go-bnb/internal/render"
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

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About function is the handler for the about page.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact function is the handler for the contact page.
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// MakeReservation function is the handler for the about page.
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

// Availability function is the handler for the search availability page.
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// PostAvailability function is the handler for the search availability page.
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s.", start, end)))
}

// jsonRespone type
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON response.
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Tidal function is the handler for the tidal page.
func (m *Repository) Tidal(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "tidal.page.html", &models.TemplateData{})
}

// Cliffside function is the handler for the Cliffside page.
func (m *Repository) Cliffside(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "cliffside.page.html", &models.TemplateData{})
}

// Oceanspray function is the handler for the Oceanspray page.
func (m *Repository) Oceanspray(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "oceanspray.page.html", &models.TemplateData{})
}

// Seabreeze function is the handler for the Seabreeze page.
func (m *Repository) Seabreeze(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "seabreeze.page.html", &models.TemplateData{})
}

// Dawn function is the handler for the Dawn page.
func (m *Repository) Dawn(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "dawn.page.html", &models.TemplateData{})
}

// Sunset function is the handler for the Sunset page.
func (m *Repository) Sunset(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "sunset.page.html", &models.TemplateData{})
}
