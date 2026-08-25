package main

import (
	"fmt";

	"io";
	"html/template"
	"github.com/labstack/echo/v4"
)

type Template struct { templates *template.Template }

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Println("Gello world!")

	renderer := &Template{ templates: template.Must(template.ParseGlob("views/*.html")) }

	// echo instance config
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Renderer = renderer

	// serving routes
	e.Static("/static", "static")
	
	fmt.Println("Now online")
	e.Logger.Fatal(e.Start(":8080"))
}
