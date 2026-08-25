package routes

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func LinkGitHub(c echo.Context) error {
	return c.Redirect(http.StatusFound, "https://github.com/theluqmn")
}

func LinkTwitter(c echo.Context) error {
	return c.Redirect(http.StatusFound, "https://x.com/theluqmn")
}

func LinkInstagram(c echo.Context) error {
	return c.Redirect(http.StatusFound, "https://instagram.com/theluqmn")
}

func LinkYouTube(c echo.Context) error {
	return c.Redirect(http.StatusFound, "https://youtube.com/@theluqmn")
}

func LinkDiscord(c echo.Context) error {
	return c.Redirect(http.StatusFound, "https://discord.gg/TxWpjjMed6")
}

func LinkHackatime(c echo.Context) error {
	return c.Redirect(http.StatusFound, "https://hackatime.hackclub.com/@theluqmn")
}