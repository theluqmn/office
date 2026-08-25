package routes

import (
	"main/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// getting handlers
	homeHandler := handlers.NewHomeHandler()
	projectHandler := handlers.NewProjectsHandler()

	// serving routes
	e.Static("/static", "static")
	e.GET("/", homeHandler.Home)
	e.GET("/projects", projectHandler.Projects)
	e.GET("/journal/:slug", handlers.ConvertMarkdownHandler)

	// shortener routes
	e.GET("/gh", LinkGitHub)
	e.GET("/x", LinkTwitter)
	e.GET("/ig", LinkInstagram)
	e.GET("/yt", LinkYouTube)
	e.GET("/dc", LinkDiscord)
	e.GET("/ht", LinkHackatime)
}