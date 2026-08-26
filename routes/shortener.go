package routes

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

func LinkGitHub(c echo.Context) error {
	log.Printf("[%s] GET /gh (redirect)", c.RealIP())
	return c.Redirect(http.StatusFound, "https://github.com/theluqmn")
}

func LinkTwitter(c echo.Context) error {
	log.Printf("[%s] GET /x (redirect)", c.RealIP())
	return c.Redirect(http.StatusFound, "https://x.com/theluqmn")
}

func LinkInstagram(c echo.Context) error {
	log.Printf("[%s] GET /ig (redirect)", c.RealIP())
	return c.Redirect(http.StatusFound, "https://instagram.com/theluqmn")
}

func LinkYouTube(c echo.Context) error {
	log.Printf("[%s] GET /yt (redirect)", c.RealIP())
	return c.Redirect(http.StatusFound, "https://youtube.com/@theluqmn")
}

func LinkDiscord(c echo.Context) error {
	log.Printf("[%s] GET /dc (redirect)", c.RealIP())
	return c.Redirect(http.StatusFound, "https://discord.gg/TxWpjjMed6")
}

func LinkHackatime(c echo.Context) error {
	log.Printf("[%s] GET /ht (redirect)", c.RealIP())
	return c.Redirect(http.StatusFound, "https://hackatime.hackclub.com/@theluqmn")
}