package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type ProjectsHandler struct {}

func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{}
}

// handles rendering of the /projects page
func (h *ProjectsHandler) Projects(c echo.Context) error {
	return c.Render(http.StatusOK, "projects", map[string]any{
		"Title": "Projects",
	})
}
