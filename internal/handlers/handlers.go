package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NathanielRand/go-bnb/internal/config"
	"github.com/NathanielRand/go-bnb/internal/forms"
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

// Availability function is the handler for the search availability page.
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// PostAvailability function is the handler for the post availability page.
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

// MakeReservation function is the handler for the make reservation page.
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	var emptyResservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyResservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation function is the handler for the post reservation page.
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	// Check first name is at least 2 characters.
	form.MinLength("first_name", 2, r)
	// Check last name is at least 2 characters.
	form.MinLength("last_name", 2, r)
	// Check email is at least 4 characters.
	form.MinLength("phone", 7, r)
	// Check email is valid format.
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// Throw reservation object into the session to be pulled out in the
	// reservation summary page and send it to the template to be displayed.
	m.App.Session.Put(r.Context(), "reservation", reservation)

	// Redirec to reservation sumary page.
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary function is the handler for the reservation summary page.
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
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
