package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Project struct {
	Name string
	Description string
	URL string
}

type ProjectsHandler struct {}

func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{}
}

// handles rendering of the /projects page
func (h *ProjectsHandler) Projects(c echo.Context) error {
	projects := []Project{
		{
			Name: "DOOS",
			Description: "A to-do list management software written in GnuCOBOL",
			URL: "https://github.com/theluqmn/doos",
		},
		{
			Name: "Whouse",
			Description: "An inventory management system built using Go.",
			URL: "https://github.com/theluqmn/whouse",
		},
	}

	return c.Render(http.StatusOK, "projects", map[string]any{
		"Title": "Projects",
		"Description": "Featured projects by @theluqmn",
		"Projects": projects,
	})
}
