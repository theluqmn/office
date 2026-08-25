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

func NewJournalHandler() *JournalHandler {
	return &JournalHandler{}
}

// handles rendering of the /journal page
func (h *JournalHandler) Journal(c echo.Context) error {
	fileData, err := os.ReadFile("./data/journal.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to read journal file")
	}

	var data JournalData
	if err := json.Unmarshal(fileData, &data); err != nil {
		return c.String(http.StatusInternalServerError, "failed to parse journal data")
	}

	fmt.Println(data.Journals)

	return c.Render(http.StatusOK, "journal", map[string]any{
		"Title":       "Journal",
		"Description": "Luqman's online journal",
		"Entries":    data.Journals,
	})
}
