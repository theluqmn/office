package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type JournalEntry struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Date        string `json:"Date"`
	URL         string `json:"URL"`
}

type JournalDataStore struct {
	Journals []JournalEntry `json:"journals"`
}

type Parser struct {
	gm goldmark.Markdown
}

func NewGMParser() *Parser {
	gm := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
		),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	return &Parser{gm: gm}
}

// extract front matter data
func (p *Parser) ExtractMeta(filePath string) (map[string]any, error) {
	content, err := os.ReadFile(filePath)
	if err != nil { return nil, err }

	context := parser.NewContext()
	if err := p.gm.Convert(content, &bytes.Buffer{}, parser.WithContext(context)); err != nil { return nil, err }

	return meta.Get(context), nil
}

// converts file content to HTML and returns metadata
func (p *Parser) ConvertMarkdown(filePath string) (template.HTML, map[string]any, error) {
	content, err := os.ReadFile(filePath)
	if err != nil { return "", nil, err }

	context := parser.NewContext()
	var buffer bytes.Buffer
	if err := p.gm.Convert(content, &buffer, parser.WithContext(context)); err != nil { return "", nil, err }

	metaData := meta.Get(context)

	return template.HTML(buffer.String()), metaData, nil
}

type JournalEntryHandler struct { parser *Parser }
func NewJournalEntryHandler() *JournalEntryHandler { return &JournalEntryHandler{ parser: NewGMParser() } }

// reads markdown files and indexes it inside data.json
func (h *JournalEntryHandler) IndexJournals(journalDir string, jsonPath string) error {
	files, err := os.ReadDir(journalDir)
	if err != nil {
		return fmt.Errorf("failed to read journal directory: %w", err)
	}

	var journalEntries []JournalEntry

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }

		filePath := filepath.Join(journalDir, file.Name())
		metaData, err := h.parser.ExtractMeta(filePath)
		if err != nil { continue }

		slug := strings.TrimSuffix(file.Name(), ".md")

		title, _ := metaData["title"].(string)
		if title == "" { title = "" }

		description, _ := metaData["description"].(string)
		if description == "" { description = "" }

		date, _ := metaData["date"].(string)

		journalEntries = append(journalEntries, JournalEntry{
			Title:       title,
			Description: description,
			Date:        date,
			URL:         "/journal/" + slug,
		})
	}

	// reverse the slice so newest/last items come first
	for i, j := 0, len(journalEntries)-1; i < j; i, j = i+1, j-1 {
		journalEntries[i], journalEntries[j] = journalEntries[j], journalEntries[i]
	}

	store := JournalDataStore{Journals: journalEntries}

	output, err := json.MarshalIndent(store, "", "    ")
	if err != nil { return fmt.Errorf("failed to marshal json: %w", err) }

	log.Println("journal entries indexed successfully")
	return os.WriteFile(jsonPath, output, 0644)
}

// serving the journal entries
func (h *JournalEntryHandler) JournalEntries(c echo.Context) error {
	slug := c.Param("slug")
	cleanSlug := filepath.Clean(slug)

	if strings.Contains(cleanSlug, "..") { return c.String(http.StatusBadRequest, "invalid path") }

	filePath := filepath.Join("../journal/entries", cleanSlug+".md")

	contentHTML, metaData, err := h.parser.ConvertMarkdown(filePath)
	if err != nil { return c.String(http.StatusNotFound, "journal entry not found") }

	title, _ := metaData["title"].(string)
	if title == "" { title = "Unnamed Journal Entry" }

	description, _ := metaData["description"].(string)
	if description == "" { description = "No description was provided" }

	date, _ := metaData["date"].(string)

	log.Printf("[%s] GET /journal/%s", c.RealIP(), cleanSlug)
	return c.Render(http.StatusOK, "journal-entry", map[string]any{
		"Title":       title,
		"Description": description,
		"Date":		   date,
		"Content":     contentHTML,
		"Meta":        metaData,
	})
}