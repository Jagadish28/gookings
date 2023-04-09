package handler

import (
	"fmt"
	"net/http"

	"github.com/Jagadish28/bookings/pkg/config"
	"github.com/Jagadish28/bookings/pkg/models"
	"github.com/Jagadish28/bookings/pkg/render"
)

// Repo va
var Repo *Repository

// Repository is repo type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	fmt.Println("OOOOOOOOOOOOOO")
	fmt.Println(remoteIp)

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is about handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	sm := make(map[string]string)
	sm["test"] = "jaga"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	sm["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: sm})
}
