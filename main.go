package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"log"

	"main/routes"
	"main/handlers"
	"github.com/labstack/echo/v4"
)

type Template struct { templates *template.Template }

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Print("\033[H\033[2J")
	log.Println("gello world! the office is initialising...")
	journalEntryHandler := handlers.NewJournalEntryHandler()
	renderer := &Template{ templates: template.Must(template.ParseFS(os.DirFS("views"), "*.html", "*/*.html")) }

	// index journals
	err := journalEntryHandler.IndexJournals("./journal", "./data/journal.json")
	if err != nil { log.Printf("failed to index journals: %v", err) }

	// echo instance config
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Renderer = renderer
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()
	
	routes.RegisterRoutes(e)
	
	log.Println("the office is now online")
	e.Logger.Fatal(e.Start(":8080"))
}
