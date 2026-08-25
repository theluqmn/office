package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Parser struct {
	gm goldmark.Markdown
}

func NewGMParser() *Parser {
	gm := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta, // Enable front matter parsing
		),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	return &Parser{gm: gm}
}

func (p *Parser) ConvertMarkdown(filePath string) (template.HTML, map[string]any, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", nil, err
	}

	context := parser.NewContext()
	var buffer bytes.Buffer

	if err := p.gm.Convert(content, &buffer, parser.WithContext(context)); err != nil {
		return "", nil, err
	}

	metaData := meta.Get(context)

	return template.HTML(buffer.String()), metaData, nil
}

func ConvertMarkdownHandler(c echo.Context) error {
	slug := c.Param("slug")
	filePath := "./journal/" + slug + ".md"
	mdParser := NewGMParser()

	contentHTML, metaData, err := mdParser.ConvertMarkdown(filePath)
	if err != nil { return c.String(http.StatusNotFound, "Journal entry not found") }

	title, _ := metaData["title"].(string)
	if title == "" { title = "Unnamed Journal Entry" }

	description, _ := metaData["description"].(string)
	if description == "" { description = "No description was provided" }

	return c.Render(http.StatusOK, "journal", map[string]any{
		"Title":       title,
		"Description": description,
		"Content":     contentHTML,
		"Meta":        metaData,
	})
}