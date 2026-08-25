package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type JournalData struct {
	Journals []JournalEntry `json:"journals"`
}

type JournalHandler struct{}
func NewJournalHandler() *JournalHandler { return &JournalHandler{} }

// handles rendering of the /journal page
func (h *JournalHandler) JournalIndex(c echo.Context) error {
	fileData, err := os.ReadFile("./data/journal.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to read journal file")
	}

	var data JournalData
	if err := json.Unmarshal(fileData, &data); err != nil {
		return c.String(http.StatusInternalServerError, "failed to parse journal data")
	}

	return c.Render(http.StatusOK, "journal-index", map[string]any{
		"Title":       "Journal",
		"Description": "Luqman's online journal",
		"Entries":    data.Journals,
	})
}