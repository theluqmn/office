package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	URL         string `json:"URL"`
}

type ProjectsData struct {
	Projects []Project `json:"projects"`
}

type ProjectsHandler struct{}
func NewProjectsHandler() *ProjectsHandler { return &ProjectsHandler{} }

// handles rendering of the /projects page
func (h *ProjectsHandler) Projects(c echo.Context) error {
	fileData, err := os.ReadFile("data/projects.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to read projects file")
	}

	var data ProjectsData
	if err := json.Unmarshal(fileData, &data); err != nil {
		return c.String(http.StatusInternalServerError, "failed to parse projects data")
	}

	return c.Render(http.StatusOK, "projects", map[string]any{
		"Title":       "Projects",
		"Description": "Featured projects by @theluqmn",
		"Projects":    data.Projects,
	})
}