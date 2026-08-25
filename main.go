package main

import (
	"fmt";

	"io"; "net/http"
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
	e.GET("/", func (c echo.Context) error {
		fmt.Println("GET /")
		return c.Render(http.StatusOK, "index", map[string]any{
			"Title": "Home",
		})
	})
	
	e.GET("/projects", func (c echo.Context) error {
		fmt.Println("GET /")
		return c.Render(http.StatusOK, "projects", map[string]any{
			"Title": "Projects",
		})
	})
	
	fmt.Println("Now online")
	e.Logger.Fatal(e.Start(":8080"))
}
