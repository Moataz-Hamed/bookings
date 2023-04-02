package handlers

import (
	"net/http"

	"github.com/Moataz-Hamed/bookings/pkg/config"
	"github.com/Moataz-Hamed/bookings/pkg/models"
	"github.com/Moataz-Hamed/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHanlder(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Moataz"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.html", &models.TemplateData{})
}

func (m *Repository) Room1(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "room1.page.html", &models.TemplateData{})
}

func (m *Repository) Room2(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "room2.page.html", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) NewRes(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new.page.html", &models.TemplateData{})
}
