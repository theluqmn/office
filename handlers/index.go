package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {}
func NewHomeHandler() *HomeHandler { return &HomeHandler{} }

// handles rendering of the root page (home)
func (h *HomeHandler) Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", map[string]any{
		"Title": "Home",
		"Description": "The office in the internet, with a look inspired by amber CRT displays.",
	})
}
