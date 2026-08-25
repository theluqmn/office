package main

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"main/routes"
	"main/handlers"
	"github.com/labstack/echo/v4"
)

type Template struct { templates *template.Template }

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Println("Gello world!")
	journalEntryHandler := handlers.NewJournalEntryHandler()
	renderer := &Template{ templates: template.Must(template.ParseFS(os.DirFS("views"), "*.html", "*/*.html")) }

	// index journals
	err := journalEntryHandler.IndexJournals("./journal", "./data/journal.json")
	if err != nil { fmt.Println(err) }

	// echo instance config
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Renderer = renderer
	routes.RegisterRoutes(e)
	
	fmt.Println("Now online")
	e.Logger.Fatal(e.Start(":8080"))
}
