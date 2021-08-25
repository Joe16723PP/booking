// Package handlers to run all file typing "go run ."
package handlers

import (
	"github.com/joe16723/booking/pkg/config"
	"github.com/joe16723/booking/pkg/models"
	"github.com/joe16723/booking/pkg/render"
	"net/http"
)

// TemplateData holds data sent from handler to template

// Repo the repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set of repository
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	// get ip from client and save to session
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	// create data
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	// retrieve data from session
	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")

	// add remote_ip to template data
	stringMap["remote_ip"] = remoteIp

	// passing data to template
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
